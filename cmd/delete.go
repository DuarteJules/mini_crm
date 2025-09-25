package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteID int

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Supprime un contact",
	Run: func(cmd *cobra.Command, args []string) {
		ok := Store.Supprimer(deleteID)
		if !ok {
			fmt.Printf("Contact ID=%d introuvable.\n", deleteID)
			return
		}
		fmt.Printf("Contact ID=%d supprimé.\n", deleteID)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().IntVarP(&deleteID, "id", "i", 0, "ID du contact")
	deleteCmd.MarkFlagRequired("id")
}
