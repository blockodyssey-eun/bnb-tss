package party

import (
	"context"
	"fmt"
	"log"
	"math/big"

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
	partyIndex, ok := new(big.Int).SetString(req.PartyId, 10)
	if !ok {
		return nil, fmt.Errorf("invalid PartyId: %s", req.PartyId)
	}
	partyID := tss.NewPartyID(req.PartyId, fmt.Sprintf("Party-%s", req.PartyId), partyIndex)

	// Create UnSortedPartyIDs
	var parties tss.UnSortedPartyIDs
	for i, p := range req.Parties {
		index := big.NewInt(int64(i))
		party := tss.NewPartyID(p.Id, p.Moniker, index)
		parties = append(parties, party)
	}

	// Sort the parties
	sortedParties := tss.SortPartyIDs(parties)

	fmt.Println("sortedParties: ", sortedParties)
	// Create KeygenManager
	km := keygen.NewKeygenManager(int(req.Threshold), int(req.TotalParties), partyID, sortedParties)

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
