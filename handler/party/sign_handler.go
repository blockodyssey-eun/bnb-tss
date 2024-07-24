package party

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"tss_project/internal/tss/sign"
	"tss_project/proto"
	pb "tss_project/proto"

	"github.com/bnb-chain/tss-lib/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/tss"
)

type SignHandler struct {
	proto.UnimplementedTSSServiceServer
}

func (h *SignHandler) InitiateSign(ctx context.Context, req *proto.SignRequest) (*proto.SignResponse, error) {
	log.Printf("Received sign request: session_id=%s", req.SessionId)

	// Create PartyID for this party
	partyID := tss.NewPartyID(req.PartyId, fmt.Sprintf("Party-%s", req.PartyId), nil)

	// Create SortedPartyIDs
	var parties tss.SortedPartyIDs
	for _, p := range req.Parties {
		party := tss.NewPartyID(p.Id, p.Moniker, nil)
		parties = append(parties, party)
	}

	// Retrieve the actual LocalPartySaveData
	// This is a placeholder. In a real scenario, you'd retrieve this data from secure storage
	key := keygen.NewLocalPartySaveData(len(parties))

	// Create SignManager
	sm := sign.NewSignManager(partyID, parties, key)

	// Start sign process
	message := new(big.Int).SetBytes(req.MessageToSign)
	R, S, err := sm.StartSign(ctx, message)
	if err != nil {
		log.Printf("Sign failed: %v", err)
		return &pb.SignResponse{
			Success:      false,
			ErrorMessage: err.Error(),
		}, nil
	}

	log.Printf("Sign successful, R: %s, S: %s", R.String(), S.String())

	return &pb.SignResponse{
		Success: true,
		SignatureData: &pb.SignatureData{
			R: R.Bytes(),
			S: S.Bytes(),
		},
	}, nil
}
