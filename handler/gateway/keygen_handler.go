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

// KeygenRequest는 키 생성 요청의 구조를 정의합니다.
type KeygenRequest struct {
	Threshold    int `json:"threshold"`     // 필요한 최소 서명자 수
	TotalParties int `json:"total_parties"` // 총 참여자 수
}

// KeygenResponse는 키 생성 응답의 구조를 정의합니다.
type KeygenResponse struct {
	PublicKey string `json:"public_key"`      // 생성된 공개키
	Error     string `json:"error,omitempty"` // 오류 메시지 (있는 경우)
}

// KeygenHandler는 키 생성 요청을 처리하는 HTTP 핸들러 함수입니다.
func KeygenHandler(w http.ResponseWriter, r *http.Request) {
	// 요청 본문을 KeygenRequest 구조체로 디코딩
	var req KeygenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "잘못된 요청 페이로드", http.StatusBadRequest)
		return
	}

	// Kubernetes에 파티 파드 생성
	partyPods, err := kubernetes.CreatePartyPods(req.TotalParties)
	if err != nil {
		http.Error(w, "파티 파드 생성 실패: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// PartyID 초기화
	parties := initializeParties(partyPods)
	fmt.Println("parties:", parties)

	// 이 게이트웨이를 위한 로컬 파티로 첫 번째 파티 사용
	localPartyID := parties[0]

	// KeygenManager 생성
	km := keygen.NewKeygenManager(req.Threshold, req.TotalParties, localPartyID, parties)
	fmt.Println("km:", parties)

	// 키 생성 프로세스 시작
	ctx := context.Background()
	publicKey, err := km.StartKeygen(ctx)

	// Kubernetes 리소스 정리
	cleanupResources(partyPods)

	// 응답 생성 및 전송
	sendResponse(w, publicKey, err)
}

// initializeParties는 파티 파드 정보를 기반으로 PartyID 슬라이스를 초기화합니다.
func initializeParties(partyPods []kubernetes.PartyPod) tss.SortedPartyIDs {
	var unsortedParties tss.UnSortedPartyIDs
	for i, pod := range partyPods {
		// 인덱스를 1부터 시작하도록 수정
		index := big.NewInt(int64(i + 1))
		partyID := tss.NewPartyID(pod.ID, pod.Name, index)
		unsortedParties = append(unsortedParties, partyID)

		// 디버그 프린트 추가
		fmt.Printf("Party created: ID=%s, Index=%d\n", pod.ID, index)
	}
	return tss.SortPartyIDs(unsortedParties)
}

// cleanupResources는 Kubernetes 리소스를 정리합니다.
func cleanupResources(partyPods []kubernetes.PartyPod) {
	if err := kubernetes.DeletePartyPods(partyPods); err != nil {
		fmt.Printf("파티 파드 삭제 중 오류 발생: %v\n", err)
	}
}

// sendResponse는 키 생성 결과를 HTTP 응답으로 전송합니다.
func sendResponse(w http.ResponseWriter, publicKey *big.Int, err error) {
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
