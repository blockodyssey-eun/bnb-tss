package gateway

import (
	"context"
	"encoding/json"
	"math/big"
	"net/http"

	"tss_project/internal/kubernetes"
	"tss_project/internal/tss/sign"

	"github.com/bnb-chain/tss-lib/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/tss"
)

type SignRequest struct {
	Message      string `json:"message"`
	TotalParties int    `json:"total_parties"`
}

type SignResponse struct {
	R     string `json:"r"`
	S     string `json:"s"`
	Error string `json:"error,omitempty"`
}

func SignHandler(w http.ResponseWriter, r *http.Request) {
	var req SignRequest
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

	// TODO: Retrieve the actual LocalPartySaveData
	// This is a placeholder. In a real scenario, you'd retrieve this data from secure storage
	key := keygen.NewLocalPartySaveData(len(parties))

	sm := sign.NewSignManager(localPartyID, parties, key)

	ctx := context.Background()
	message := new(big.Int).SetBytes([]byte(req.Message))
	R, S, err := sm.StartSign(ctx, message)

	// Clean up Kubernetes resources
	kubernetes.DeletePartyPods(partyPods)

	response := SignResponse{}
	if err != nil {
		response.Error = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		response.R = R.String()
		response.S = S.String()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
