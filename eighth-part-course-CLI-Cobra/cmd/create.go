/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		db := GetDb()
		category := GetCategoryDB(db)

		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")

		_, err := category.Create(name, description)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	categoryCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("name", "n", "", "Name of the Category")
	createCmd.Flags().StringP("description", "d", "", "Description of the Category")
	//createCmd.MarkFlagRequired("name") <- one way to receive a value
	//createCmd.MarkFlagRequired("description") <- one way to receive a value
	createCmd.MarkFlagsRequiredTogether("name", "description")

}
