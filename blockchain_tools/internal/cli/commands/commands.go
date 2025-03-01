// internal/cli/commands/commands.go
package commands



import (
	"github.com/spf13/cobra"
	"github.com/layla-lili/blockchain_tools/internal/cli/config"
	"github.com/layla-lili/blockchain_tools/internal/common/logging"
)

var (
	cfgFile string
	logger  = logging.NewLogger()
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "blockchain-cli",
	Short: "A CLI for interacting with the blockchain",
	Long: `A command line interface for interacting with the blockchain MVP.
This tool provides commands for querying the blockchain state, sending transactions,
managing accounts, and monitoring the network.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	// Global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.blockchain-cli.yaml)")
	rootCmd.PersistentFlags().String("rpc-url", "http://localhost:8545", "URL of the blockchain RPC endpoint")
	rootCmd.PersistentFlags().String("format", "table", "Output format (table, json, yaml)")
	rootCmd.PersistentFlags().Bool("debug", false, "Enable debug logging")

	// Add subcommands
	rootCmd.AddCommand(newBlockCmd())
	rootCmd.AddCommand(newTransactionCmd())
	rootCmd.AddCommand(newAccountCmd())
	rootCmd.AddCommand(newNodeCmd())
	rootCmd.AddCommand(newVersionCmd())
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	config.InitConfig(cfgFile)
	
	// Set up logging based on debug flag
	if debug, _ := rootCmd.PersistentFlags().GetBool("debug"); debug {
		logging.SetLevel("debug")
	}
}