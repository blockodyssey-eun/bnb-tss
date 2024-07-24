package gateway

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"

	"tss_project/internal/kubernetes"
	"tss_project/internal/tss/keygen"

	"github.com/bnb-chain/tss-lib/tss"
)

type KeygenRequest struct {
	Threshold    int `json:"threshold"`
	TotalParties int `json:"total_parties"`
}

type KeygenResponse struct {
	PublicKey string `json:"public_key"`
	Error     string `json:"error,omitempty"`
}

func KeygenHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("asdf")
	var req KeygenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Create Kubernetes party pods
	partyPods, err := kubernetes.CreatePartyPods(req.TotalParties)

	if err != nil {
		http.Error(w, "Failed to create party pods: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Initialize PartyIDs
	var parties tss.SortedPartyIDs
	for i, pod := range partyPods {
		partyID := tss.NewPartyID(pod.ID, pod.Name, big.NewInt(int64(i)))
		parties = append(parties, partyID)
	}

	// Use the first party as the local party for this gateway
	localPartyID := parties[0]

	km := keygen.NewKeygenManager(req.Threshold, req.TotalParties, localPartyID, parties)

	ctx := context.Background()
	publicKey, err := km.StartKeygen(ctx)

	// Clean up Kubernetes resources
	kubernetes.DeletePartyPods(partyPods)

	response := KeygenResponse{}
	if err != nil {
		response.Error = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		response.PublicKey = publicKey.String()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
