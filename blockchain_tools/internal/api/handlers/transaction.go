package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/layla-lili/blockchain_tools/pkg/client/rpc"
    "github.com/layla-lili/blockchain_tools/pkg/types"
)

func GetTransaction(client *rpc.Client) gin.HandlerFunc {
    return func(c *gin.Context) {
        hash := c.Param("hash")
        tx, err := client.GetTransaction(c.Request.Context(), hash)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, tx)
    }
}

func SendTransaction(client *rpc.Client) gin.HandlerFunc {
    return func(c *gin.Context) {
        var tx types.Transaction
        if err := c.BindJSON(&tx); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        hash, err := client.SendTransaction(c.Request.Context(), &tx)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"hash": hash})
    }
}