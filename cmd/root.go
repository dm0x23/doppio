/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dop",
	Short: "Doppio – a double shot of speed for your shell",
	Long:  `Doppio helps you create and manage shell shortcuts with ease.`,
}

func Execute() error {
	return rootCmd.Execute()
}
