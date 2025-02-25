package cmd

import (
	"github.com/lajosdeme/autonomi-circulating-supply/internal"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts ANT circulating supply service",
	Run: func(cmd *cobra.Command, args []string) {
		internal.Start()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
