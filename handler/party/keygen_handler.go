package party

import (
	"context"
	"fmt"
	"log"

	"tss_project/internal/tss/keygen"
	"tss_project/proto"
	pb "tss_project/proto"

	"github.com/bnb-chain/tss-lib/tss"
)

type KeygenHandler struct {
	proto.UnimplementedTSSServiceServer
}

func (h *KeygenHandler) InitiateKeygen(ctx context.Context, req *proto.KeygenRequest) (*proto.KeygenResponse, error) {
	log.Printf("Received keygen request: threshold=%d, total_parties=%d", req.Threshold, req.TotalParties)

	// Create PartyID for this party
	partyID := tss.NewPartyID(req.PartyId, fmt.Sprintf("Party-%s", req.PartyId), nil)

	// Create SortedPartyIDs
	var parties tss.SortedPartyIDs
	for _, p := range req.Parties {
		party := tss.NewPartyID(p.Id, p.Moniker, nil)
		parties = append(parties, party)
	}

	// Create KeygenManager
	km := keygen.NewKeygenManager(int(req.Threshold), int(req.TotalParties), partyID, parties)

	// Start keygen process
	publicKey, err := km.StartKeygen(ctx)
	if err != nil {
		log.Printf("Keygen failed: %v", err)
		return &pb.KeygenResponse{
			Success:      false,
			ErrorMessage: err.Error(),
		}, nil
	}

	log.Printf("Keygen successful, public key: %s", publicKey.String())

	return &pb.KeygenResponse{
		Success:   true,
		PublicKey: publicKey.String(),
	}, nil
}
