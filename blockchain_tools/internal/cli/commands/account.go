package commands

import (
	"context"
	"fmt"
	"math/big"

	"github.com/layla-lili/blockchain_tools/internal/cli/formatter"
	"github.com/layla-lili/blockchain_tools/pkg/client/rpc"
	"github.com/spf13/cobra"
)

func newAccountCmd() *cobra.Command {
	accountCmd := &cobra.Command{
		Use:   "account",
		Short: "Manage blockchain accounts",
		Long:  `Commands to create and manage blockchain accounts/wallets.`,
	}

	// Add account subcommands
	accountCmd.AddCommand(newAccountCreateCmd())
	accountCmd.AddCommand(newAccountListCmd())
	accountCmd.AddCommand(newAccountBalanceCmd())

	return accountCmd
}

func newAccountCreateCmd() *cobra.Command {
	var password string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new account",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()

			rpcURL, _ := cmd.Flags().GetString("rpc-url")
			client, err := rpc.NewClient(rpcURL)
			if err != nil {
				return fmt.Errorf("failed to create client: %w", err)
			}

			account, err := client.CreateAccount(ctx, password)
			if err != nil {
				return fmt.Errorf("failed to create account: %w", err)
			}

			format, _ := cmd.Flags().GetString("format")
			fmt := formatter.GetFormatter(format)
			return fmt.Format(cmd.OutOrStdout(), account)
		},
	}

	cmd.Flags().StringVarP(&password, "password", "p", "", "Account password")
	cmd.MarkFlagRequired("password")

	return cmd
}

func newAccountListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all accounts",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()

			rpcURL, _ := cmd.Flags().GetString("rpc-url")
			client, err := rpc.NewClient(rpcURL)
			if err != nil {
				return fmt.Errorf("failed to create client: %w", err)
			}

			accounts, err := client.ListAccounts(ctx)
			if err != nil {
				return fmt.Errorf("failed to list accounts: %w", err)
			}

			format, _ := cmd.Flags().GetString("format")
			fmt := formatter.GetFormatter(format)
			return fmt.Format(cmd.OutOrStdout(), accounts)
		},
	}
}

func newAccountBalanceCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "balance [address]",
		Short: "Get account balance",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()

			rpcURL, _ := cmd.Flags().GetString("rpc-url")
			client, err := rpc.NewClient(rpcURL)
			if err != nil {
				return fmt.Errorf("failed to create client: %w", err)
			}

			address := args[0]
			balance, err := client.GetAccountBalance(ctx, address) // Use new method
			if err != nil {
				return fmt.Errorf("failed to get balance: %w", err)
			}

			response := struct {
				Address string   `json:"address"`
				Balance *big.Int `json:"balance"`
			}{
				Address: address,
				Balance: balance,
			}

			format, _ := cmd.Flags().GetString("format")
			fmt := formatter.GetFormatter(format)
			return fmt.Format(cmd.OutOrStdout(), response)
		},
	}
}
