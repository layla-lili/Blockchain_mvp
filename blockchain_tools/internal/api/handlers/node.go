package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/layla-lili/blockchain_tools/pkg/client/rpc"
)

func GetNodeStatus(client *rpc.Client) gin.HandlerFunc {
    return func(c *gin.Context) {
        status, err := client.GetNodeStatus(c.Request.Context())
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, status)
    }
}

func GetPeers(client *rpc.Client) gin.HandlerFunc {
    return func(c *gin.Context) {
        peers, err := client.GetPeers(c.Request.Context())
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, peers)
    }
}