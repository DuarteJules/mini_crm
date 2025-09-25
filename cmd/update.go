package cmd

import (
	"fmt"

	"github.com/DuarteJules/mini_crm/internal/storage"
	"github.com/spf13/cobra"
)

var (
	updateID int
	newNom   string
	newEmail string
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Met à jour un contact existant",
	Run: func(cmd *cobra.Command, args []string) {
		c := storage.Contact{ID: uint(updateID), Nom: newNom, Email: newEmail}
		updated, ok := Store.MettreAJour(c)
		if !ok {
			fmt.Printf("Contact ID=%d introuvable.\n", updateID)
			return
		}
		fmt.Printf("Contact mis à jour: %+v\n", updated)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().IntVarP(&updateID, "id", "i", 0, "ID du contact")
	updateCmd.Flags().StringVar(&newNom, "nom", "", "Nouveau nom")
	updateCmd.Flags().StringVar(&newEmail, "email", "", "Nouvel email")
	updateCmd.MarkFlagRequired("id")
}
