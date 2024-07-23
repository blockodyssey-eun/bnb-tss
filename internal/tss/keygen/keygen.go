package keygen

import (
	"context"
	"fmt"
	"math/big"

	"github.com/bnb-chain/tss-lib/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/tss"
)

// KeygenManager handles the key generation process
type KeygenManager struct {
	Threshold    int
	TotalParties int
	PartyID      *tss.PartyID
	Parties      tss.SortedPartyIDs
}

// NewKeygenManager creates a new KeygenManager
func NewKeygenManager(threshold, totalParties int, partyID *tss.PartyID, parties tss.SortedPartyIDs) *KeygenManager {
	return &KeygenManager{
		Threshold:    threshold,
		TotalParties: totalParties,
		PartyID:      partyID,
		Parties:      parties,
	}
}

// StartKeygen starts the key generation process
func (km *KeygenManager) StartKeygen(ctx context.Context) (*big.Int, error) {
	peerCtx := tss.NewPeerContext(km.Parties)
	params := tss.NewParameters(tss.S256(), peerCtx, km.PartyID, km.TotalParties, km.Threshold)

	outCh := make(chan tss.Message, km.TotalParties)
	endCh := make(chan keygen.LocalPartySaveData, 1)

	party := keygen.NewLocalParty(params, outCh, endCh)

	// Start the key generation process
	err := party.Start()
	if err != nil {
		return nil, fmt.Errorf("failed to start key generation: %v", err)
	}

	// Wait for the key generation process to complete
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case data := <-endCh:
		return data.ECDSAPub.X(), nil
	}
}
