/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// categoryCmd represents the category command
var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//name, _ := cmd.Flags().GetString("name")
		fmt.Println("Category  called called with name: " + category)

		exists, _ := cmd.Flags().GetBool("exists")
		fmt.Println("Category called with exists: " + fmt.Sprint(exists))

		id, _ := cmd.Flags().GetInt16("ID")
		fmt.Println("Category called with ID: " + fmt.Sprint(id))

	},
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Calling pre before Run")
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Calling after Run")
	},
	//RunE: func(cmd *cobra.Command, args []string) error { <-- run with some error
	//	return fmt.Errorf("Run with Error")
	//},
}

var category string

func init() {
	rootCmd.AddCommand(categoryCmd)

	// StringVarP will pass the value of the flag to variable category
	categoryCmd.PersistentFlags().StringVarP(&category, "name", "n", "", "CategoryName") //< this will work for sub commands like list and create

	// StringP allows adding shortcut
	//categoryCmd.PersistentFlags().StringP("name", "n", "Y", "Name of the Category") //< this will work for sub commands like list and create

	categoryCmd.PersistentFlags().BoolP("exists", "e", false, "Check if Category boolean input exists")

	categoryCmd.PersistentFlags().Int16P("ID", "i", 0, "ID of the Category")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// categoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// categoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
