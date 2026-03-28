package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yebyen/useless/cli/pkg/backend"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "useless",
		Short: "Interact with the Useless Machine",
		Long:  `A CLI tool to push the button and check status of the Useless Machine across K8s and Spin API.`,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mecris.config)")
	rootCmd.PersistentFlags().String("backend", "k8s", "backend to use (k8s or api)")
	viper.BindPFlag("backend", rootCmd.PersistentFlags().Lookup("backend"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".mecris.config")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		// fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func getBackend() (backend.Backend, error) {
	b := viper.GetString("backend")
	switch b {
	case "k8s":
		return backend.NewK8sBackend()
	case "api":
		return &backend.ApiBackend{BaseURL: viper.GetString("api_url")}, nil
	default:
		return nil, fmt.Errorf("unknown backend: %s", b)
	}
}
