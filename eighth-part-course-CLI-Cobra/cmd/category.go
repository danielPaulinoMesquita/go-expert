/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// categoryCmd represents the category command
var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var category string

func init() {
	rootCmd.AddCommand(categoryCmd)
}
