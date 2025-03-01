package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/layla-lili/blockchain_tools/pkg/client/rpc"
)

func GetAccount(client *rpc.Client) gin.HandlerFunc {
    return func(c *gin.Context) {
        address := c.Param("address")
        accounts, err := client.ListAccounts(c.Request.Context())
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        for _, acc := range accounts {
            if acc.Address == address {
                c.JSON(http.StatusOK, acc)
                return
            }
        }

        c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
    }
}

func GetBalance(client *rpc.Client) gin.HandlerFunc {
    return func(c *gin.Context) {
        address := c.Param("address")
        balance, err := client.GetAccountBalance(c.Request.Context(), address)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"address": address, "balance": balance.String()})
    }
}