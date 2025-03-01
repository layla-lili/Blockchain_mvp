package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/layla-lili/blockchain_tools/pkg/client/rpc"
)

func GetBlock(client *rpc.Client) gin.HandlerFunc {
    return func(c *gin.Context) {
        number, err := strconv.ParseUint(c.Param("number"), 10, 64)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid block number"})
            return
        }

        block, err := client.GetBlockByHeight(c.Request.Context(), number)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, block)
    }
}

func GetLatestBlock(client *rpc.Client) gin.HandlerFunc {
    return func(c *gin.Context) {
        block, err := client.GetLatestBlock(c.Request.Context())
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, block)
    }
}