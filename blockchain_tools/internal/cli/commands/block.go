// internal/cli/commands/block.go
package commands

import (
	"context"
	"fmt"
	"strconv"

	"github.com/layla-lili/blockchain_tools/internal/cli/formatter"
	"github.com/layla-lili/blockchain_tools/pkg/client/rpc"
	"github.com/layla-lili/blockchain_tools/pkg/types"
	"github.com/spf13/cobra"
)

func newBlockCmd() *cobra.Command {
	blockCmd := &cobra.Command{
		Use:   "block",
		Short: "Manage blockchain blocks",
		Long:  `Commands to query and interact with blockchain blocks.`,
	}

	// Add subcommands
	blockCmd.AddCommand(newGetBlockCmd())
	blockCmd.AddCommand(newGetBlocksCmd())
	blockCmd.AddCommand(newGetBlockCountCmd())

	return blockCmd
}

func newGetBlockCmd() *cobra.Command {
	getBlockCmd := &cobra.Command{
		Use:   "get [hash_or_height]",
		Short: "Get a single block by hash or height",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()

			// Get RPC URL from flag or config
			rpcURL, _ := cmd.Flags().GetString("rpc-url")

			// Create client
			client, err := rpc.NewClient(rpcURL)
			if err != nil {
				return fmt.Errorf("failed to create client: %w", err)
			}

			// Parse argument (could be hash or height)
			arg := args[0]

			var block interface{}

			// Try to parse as block height first
			if height, parseErr := strconv.ParseUint(arg, 10, 64); parseErr == nil {
				block, err = client.GetBlockByHeight(ctx, height)
			} else {
				// Assume it's a hash
				block, err = client.GetBlockByHash(ctx, arg)
			}

			if err != nil {
				return fmt.Errorf("failed to get block: %w", err)
			}

			// Get output format
			format, _ := cmd.Flags().GetString("format")

			// Format output
			fmt := formatter.GetFormatter(format)
			return fmt.Format(cmd.OutOrStdout(), block)
		},
	}

	return getBlockCmd
}

func newGetBlocksCmd() *cobra.Command {
	getBlocksCmd := &cobra.Command{
		Use:   "list [start_height] [end_height]",
		Short: "List a range of blocks",
		Long: `List multiple blocks from the blockchain by specifying a start and end height.
Example: blockchain-cli block list 1000 1010`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()

			// Get RPC URL from flag
			rpcURL, _ := cmd.Flags().GetString("rpc-url")

			// Create client
			client, err := rpc.NewClient(rpcURL)
			if err != nil {
				return fmt.Errorf("failed to create client: %w", err)
			}

			// Parse start and end heights
			startHeight, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid start height: %w", err)
			}

			endHeight, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid end height: %w", err)
			}

			// Validate range
			if endHeight < startHeight {
				return fmt.Errorf("end height must be greater than or equal to start height")
			}

			// Safety limit
			const maxBlocks = 100
			if endHeight-startHeight > maxBlocks {
				return fmt.Errorf("maximum range of %d blocks exceeded", maxBlocks)
			}

			// Fetch blocks
			var blocks []*types.Block
			for height := startHeight; height <= endHeight; height++ {
				block, err := client.GetBlockByHeight(ctx, height)
				if err != nil {
					return fmt.Errorf("failed to get block at height %d: %w", height, err)
				}
				blocks = append(blocks, block)
			}

			// Format output
			format, _ := cmd.Flags().GetString("format")
			fmt := formatter.GetFormatter(format)
			return fmt.Format(cmd.OutOrStdout(), blocks)
		},
	}

	return getBlocksCmd
}

func newGetBlockCountCmd() *cobra.Command {
	getBlockCountCmd := &cobra.Command{
		Use:   "count",
		Short: "Get the current block height",
		Long:  `Retrieve the current block height (total number of blocks) in the blockchain.`,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()

			// Get RPC URL from flag
			rpcURL, _ := cmd.Flags().GetString("rpc-url")

			// Create client
			client, err := rpc.NewClient(rpcURL)
			if err != nil {
				return fmt.Errorf("failed to create client: %w", err)
			}

			// Get latest block to determine height
			block, err := client.GetLatestBlock(ctx)
			if err != nil {
				return fmt.Errorf("failed to get latest block: %w", err)
			}

			// Format output
			format, _ := cmd.Flags().GetString("format")
			fmt := formatter.GetFormatter(format)

			// Create a simple response structure
			response := struct {
				BlockHeight uint64 `json:"blockHeight"`
				Timestamp   int64  `json:"timestamp"`
			}{
				BlockHeight: block.Height,
				Timestamp:   block.Timestamp,
			}

			return fmt.Format(cmd.OutOrStdout(), response)
		},
	}

	return getBlockCountCmd
}
