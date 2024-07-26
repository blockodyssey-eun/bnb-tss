package handler

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"

	grpcClient "gateway/internal/grpc"
	"gateway/internal/k8s"
)

type KeygenRequest struct {
	N int `json:"n" binding:"required"`
	M int `json:"m" binding:"required"`
}

type KeygenResponse struct {
	PublicKey string `json:"publickey"`
}

// Keygen은 HTTP 요청을 처리하는 핸들러 함수입니다.
func Keygen(keygenServer *grpcClient.KeygenServiceServer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req KeygenRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 대기 풀에서 리소스가 여유 있는 Pod 가져오기
		pods, err := k8s.GetPodsFromPool(req.M)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var wg sync.WaitGroup
		var mu sync.Mutex
		var publicKeys []string

		for _, pod := range pods {
			wg.Add(1)
			go func(podIP string) {
				defer wg.Done()
				// Pod의 키 생성 서비스를 호출합니다.
				_, err := grpcClient.CallKeygenService(podIP, int32(req.N), int32(req.M), pods)
				if err != nil {
					log.Printf("Failed to call keygen service on pod %s: %v", podIP, err)
				}
			}(pod.Status.PodIP)
		}

		// 키 생성 완료 메시지를 수신
		go func() {
			for key := range keygenServer.FinishedKeys {
				mu.Lock()
				publicKeys = append(publicKeys, key)
				mu.Unlock()
				if len(publicKeys) == req.M {
					close(keygenServer.FinishedKeys)
				}
			}
		}()

		// 모든 고루틴이 완료될 때까지 대기
		wg.Wait()

		// 키 생성 완료 메시지를 수신할 때까지 스트리밍을 통해 대기
		if len(publicKeys) == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate keys"})
			return
		}

		// 생성된 첫 번째 공개키를 응답으로 반환합니다.
		// 주의: 실제 구현에서는 모든 키를 결합하거나 처리하는 로직이 필요할 수 있습니다.
		c.JSON(http.StatusOK, KeygenResponse{PublicKey: publicKeys[0]})
	}
}
