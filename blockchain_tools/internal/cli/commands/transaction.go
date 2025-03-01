package commands

import (
	"context"
	"fmt"

	"github.com/layla-lili/blockchain_tools/internal/cli/formatter"
	"github.com/layla-lili/blockchain_tools/pkg/client/rpc"
	"github.com/layla-lili/blockchain_tools/pkg/types"
	"github.com/spf13/cobra"
)

func newTransactionCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:   "tx",
		Short: "Manage blockchain transactions",
		Long:  `Commands to create, send, and query blockchain transactions.`,
	}

	// Add transaction subcommands
	txCmd.AddCommand(newGetTransactionCmd())
	txCmd.AddCommand(newSendTransactionCmd())
	txCmd.AddCommand(newListTransactionsCmd())

	return txCmd
}

func newGetTransactionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get [hash]",
		Short: "Get transaction details by hash",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()

			rpcURL, _ := cmd.Flags().GetString("rpc-url")
			client, err := rpc.NewClient(rpcURL)
			if err != nil {
				return fmt.Errorf("failed to create client: %w", err)
			}

			tx, err := client.GetTransaction(ctx, args[0])
			if err != nil {
				return fmt.Errorf("failed to get transaction: %w", err)
			}

			format, _ := cmd.Flags().GetString("format")
			fmt := formatter.GetFormatter(format)
			return fmt.Format(cmd.OutOrStdout(), tx)
		},
	}
}

func newSendTransactionCmd() *cobra.Command {
	var (
		to    string
		value uint64
		data  string
	)

	cmd := &cobra.Command{
		Use:   "send",
		Short: "Send a new transaction",
		Long:  `Create and send a new transaction to the blockchain.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()

			rpcURL, _ := cmd.Flags().GetString("rpc-url")
			client, err := rpc.NewClient(rpcURL)
			if err != nil {
				return fmt.Errorf("failed to create client: %w", err)
			}

			tx := &types.Transaction{
				To:    to,
				Value: value,
				Data:  []byte(data),
			}

			hash, err := client.SendTransaction(ctx, tx)
			if err != nil {
				return fmt.Errorf("failed to send transaction: %w", err)
			}

			fmt.Printf("Transaction sent successfully! Hash: %s\n", hash)
			return nil
		},
	}

	// Transaction-specific flags
	cmd.Flags().StringVar(&to, "to", "", "Recipient address")
	cmd.Flags().Uint64Var(&value, "value", 0, "Transaction value")
	cmd.Flags().StringVar(&data, "data", "", "Transaction data (optional)")
	cmd.MarkFlagRequired("to")

	return cmd
}

func newListTransactionsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List recent transactions",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()

			rpcURL, _ := cmd.Flags().GetString("rpc-url")
			client, err := rpc.NewClient(rpcURL)
			if err != nil {
				return fmt.Errorf("failed to create client: %w", err)
			}

			txs, err := client.ListTransactions(ctx)
			if err != nil {
				return fmt.Errorf("failed to list transactions: %w", err)
			}

			format, _ := cmd.Flags().GetString("format")
			fmt := formatter.GetFormatter(format)
			return fmt.Format(cmd.OutOrStdout(), txs)
		},
	}
}
