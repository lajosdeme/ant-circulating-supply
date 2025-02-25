package cmd

import (
	"github.com/lajosdeme/autonomi-circulating-supply/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "supply",
	Short: "Entrypoint to circulating supply service",
	Args:  cobra.NoArgs,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.Load()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
