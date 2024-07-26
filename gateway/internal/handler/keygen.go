package handler

import (
	"log"
	"net/http"
	"strconv"

	"gateway/internal/k8s"

	"github.com/gin-gonic/gin"
)

type KeygenRequest struct {
	N int `json:"n" binding:"required"`
	M int `json:"m" binding:"required"`
}

type KeygenResponse struct {
	PublicKey string `json:"publickey"`
}

func Keygen(c *gin.Context) {
	var req KeygenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a public key based on n and m (for example purposes, we'll just concatenate them)
	publicKey := "publickey_" + strconv.Itoa(req.N) + "_" + strconv.Itoa(req.M)

	// 대기 풀에서 리소스가 여유 있는 Pod 가져오기
	pods, err := k8s.GetPodsFromPool(req.M)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, pod := range pods {
		log.Printf("Using pod: %s", pod.Name)
	}

	c.JSON(http.StatusOK, KeygenResponse{PublicKey: publicKey})
}
