// internal/cli/config/config.go
package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

// Config holds the application configuration
type Config struct {
	RpcURL    string `mapstructure:"rpc_url"`
	Format    string `mapstructure:"format"`
	Debug     bool   `mapstructure:"debug"`
	KeyFile   string `mapstructure:"key_file"`
	APIKey    string `mapstructure:"api_key"`
}

// GetConfig returns the current configuration
func GetConfig() *Config {
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Error unmarshalling config: %s\n", err)
	}
	return &config
}

// InitConfig initializes the configuration
func InitConfig(cfgFile string) {
	if cfgFile != "" {
		// Use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".blockchain-cli"
		viper.AddConfigPath(home)
		viper.SetConfigName(".blockchain-cli")
	}

	// Read environment variables
	viper.SetEnvPrefix("BLOCKCHAIN")
	viper.AutomaticEnv()

	// Set defaults
	viper.SetDefault("rpc_url", "http://localhost:8545")
	viper.SetDefault("format", "table")
	viper.SetDefault("debug", false)
	viper.SetDefault("key_file", filepath.Join(homeDir(), ".blockchain-cli", "keys.json"))

	// If a config file is found, read it in
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// homeDir returns the user's home directory
func homeDir() string {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return home
}
