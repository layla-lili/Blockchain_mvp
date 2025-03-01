package commands

import (
	"context"
	"fmt"

	"github.com/layla-lili/blockchain_tools/internal/cli/formatter"
	"github.com/layla-lili/blockchain_tools/pkg/client/rpc"
	"github.com/spf13/cobra"
)

func newNodeCmd() *cobra.Command {
	nodeCmd := &cobra.Command{
		Use:   "node",
		Short: "Manage blockchain node",
		Long:  `Commands to interact with and monitor the blockchain node.`,
	}

	// Add node subcommands
	nodeCmd.AddCommand(newNodeStatusCmd())
	nodeCmd.AddCommand(newNodePeersCmd())
	nodeCmd.AddCommand(newNodeSyncCmd())

	return nodeCmd
}

func newNodeStatusCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Get node status",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()

			rpcURL, _ := cmd.Flags().GetString("rpc-url")
			client, err := rpc.NewClient(rpcURL)
			if err != nil {
				return fmt.Errorf("failed to create client: %w", err)
			}

			status, err := client.GetNodeStatus(ctx)
			if err != nil {
				return fmt.Errorf("failed to get node status: %w", err)
			}

			format, _ := cmd.Flags().GetString("format")
			fmt := formatter.GetFormatter(format)
			return fmt.Format(cmd.OutOrStdout(), status)
		},
	}
}

func newNodePeersCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "peers",
		Short: "List connected peers",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()

			rpcURL, _ := cmd.Flags().GetString("rpc-url")
			client, err := rpc.NewClient(rpcURL)
			if err != nil {
				return fmt.Errorf("failed to create client: %w", err)
			}

			peers, err := client.GetPeers(ctx)
			if err != nil {
				return fmt.Errorf("failed to get peers: %w", err)
			}

			format, _ := cmd.Flags().GetString("format")
			fmt := formatter.GetFormatter(format)
			return fmt.Format(cmd.OutOrStdout(), peers)
		},
	}
}

func newNodeSyncCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "sync",
		Short: "Get node synchronization status",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()

			rpcURL, _ := cmd.Flags().GetString("rpc-url")
			client, err := rpc.NewClient(rpcURL)
			if err != nil {
				return fmt.Errorf("failed to create client: %w", err)
			}

			syncStatus, err := client.GetSyncStatus(ctx)
			if err != nil {
				return fmt.Errorf("failed to get sync status: %w", err)
			}

			format, _ := cmd.Flags().GetString("format")
			fmt := formatter.GetFormatter(format)
			return fmt.Format(cmd.OutOrStdout(), syncStatus)
		},
	}
}
