/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/daniel/CLI/internal/database"
	"github.com/spf13/cobra"
)

func newCreateCmd(categoryDb database.Category) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create a new category",
		Long:  `Create a new Category`,
		RunE:  runCreate(categoryDb),
	}
}

// createCmd represents the create command
// Doesn't use more
//var createCmd = &cobra.Command{
//	Use:   "create",
//	Short: "A brief description of your command",
//	Long: `A longer description that spans multiple lines and likely contains examples
//and usage of using your command. For example:
//
//Cobra is a CLI library for Go that empowers applications.
//This application is a tool to generate the needed files
//to quickly create a Cobra application.`,
//	RunE: runCreate(GetCategoryDB(GetDb())),
//}

func runCreate(categoryDb database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")

		_, err := categoryDb.Create(name, description)
		if err != nil {
			return err
		}

		return nil
	}
}

func init() {
	createCmd := newCreateCmd(GetCategoryDB(GetDb()))
	categoryCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("name", "n", "", "Name of the Category")
	createCmd.Flags().StringP("description", "d", "", "Description of the Category")
	//createCmd.MarkFlagRequired("name") <- one way to receive a value
	//createCmd.MarkFlagRequired("description") <- one way to receive a value
	createCmd.MarkFlagsRequiredTogether("name", "description")

}
