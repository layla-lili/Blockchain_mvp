package commands

import (
	"fmt"
	"strings"
	"time"

	"github.com/layla-lili/blockchain_tools/pkg/client/rpc"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.Flags().String("rpc-url", "http://localhost:8545", "RPC URL of the blockchain node")
	testCmd.Flags().Bool("verbose", false, "Show detailed information")
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test connection to local blockchain",
	RunE: func(cmd *cobra.Command, args []string) error {
		rpcURL, _ := cmd.Flags().GetString("rpc-url")
		verbose, _ := cmd.Flags().GetBool("verbose")

		// Test connection to Anvil
		client, err := rpc.NewClient(rpcURL)
		if err != nil {
			return fmt.Errorf("failed to connect to node: %v", err)
		}

		// Get blockchain info
		chainID, err := client.ChainID(cmd.Context())
		if err != nil {
			return fmt.Errorf("failed to get chain ID: %v", err)
		}

		blockNumber, err := client.BlockNumber(cmd.Context())
		if err != nil {
			return fmt.Errorf("failed to get block number: %v", err)
		}

		cmd.Printf("Connected to local node (Chain ID: %d, Block: %d)\n", chainID, blockNumber)

		// Get test accounts
		accounts, err := client.GetAccounts(cmd.Context())
		if err != nil {
			return fmt.Errorf("failed to get accounts: %v", err)
		}

		cmd.Printf("\nAvailable test accounts:\n")
		for i, acc := range accounts {
			balance, err := client.GetBalance(cmd.Context(), acc)
			if err != nil {
				cmd.Printf("(%d) %s: error getting balance: %v\n", i, acc, err)
				continue
			}
			cmd.Printf("(%d) %s: %s ETH\n", i, acc, balance)
		}

		if verbose {
			// Test block production
			cmd.Printf("\nWaiting for new block...")
			time.Sleep(2 * time.Second)

			newBlockNumber, err := client.BlockNumber(cmd.Context())
			if err != nil {
				return fmt.Errorf("failed to get new block number: %v", err)
			}

			if newBlockNumber > blockNumber {
				cmd.Printf("New block produced: %d\n", newBlockNumber)
			} else {
				cmd.Printf("No new block yet (still at %d)\n", blockNumber)
			}

			// Only try to get peer count if we're not using Anvil
			if !strings.Contains(rpcURL, "localhost:8545") {
				peers, err := client.PeerCount(cmd.Context())
				if err != nil {
					cmd.Printf("Note: Peer count not supported on Anvil\n")
				} else {
					cmd.Printf("Connected peers: %d\n", peers)
				}
			}
		}

		return nil
	},
}
