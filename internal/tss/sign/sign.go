package sign

import (
	"context"
	"fmt"
	"math/big"

	"github.com/bnb-chain/tss-lib/common"
	"github.com/bnb-chain/tss-lib/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/ecdsa/signing"
	"github.com/bnb-chain/tss-lib/tss"
)

// SignManager handles the signing process
type SignManager struct {
	PartyID *tss.PartyID
	Parties tss.SortedPartyIDs
	Key     keygen.LocalPartySaveData
}

// NewSignManager creates a new SignManager
func NewSignManager(partyID *tss.PartyID, parties tss.SortedPartyIDs, key keygen.LocalPartySaveData) *SignManager {
	return &SignManager{
		PartyID: partyID,
		Parties: parties,
		Key:     key,
	}
}

// StartSign starts the signing process
func (sm *SignManager) StartSign(ctx context.Context, msg *big.Int) (*big.Int, *big.Int, error) {
	peerCtx := tss.NewPeerContext(sm.Parties)
	params := tss.NewParameters(tss.S256(), peerCtx, sm.PartyID, len(sm.Parties), len(sm.Parties)-1)

	outCh := make(chan tss.Message, len(sm.Parties))
	endCh := make(chan common.SignatureData, 1)

	party := signing.NewLocalParty(msg, params, sm.Key, outCh, endCh)

	// Start the signing process
	err := party.Start()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to start signing: %v", err)
	}

	// Wait for the signing process to complete
	select {
	case <-ctx.Done():
		return nil, nil, ctx.Err()
	case data := <-endCh:
		R := new(big.Int).SetBytes(data.R)
		S := new(big.Int).SetBytes(data.S)
		return R, S, nil
	}
}
