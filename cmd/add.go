package cmd

import (
	"fmt"

	"github.com/DuarteJules/mini_crm/internal/storage"
	"github.com/spf13/cobra"
)

var (
	nom   string
	email string
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Ajoute un nouveau contact",
	Run: func(cmd *cobra.Command, args []string) {
		c := storage.Contact{Nom: nom, Email: email}
		created := Store.Ajouter(c)
		fmt.Printf("Contact ajouté: ID=%d, Nom=%s, Email=%s\n", created.ID, created.Nom, created.Email)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&nom, "nom", "n", "", "Nom du contact")
	addCmd.Flags().StringVarP(&email, "email", "e", "", "Email du contact")
	addCmd.MarkFlagRequired("nom")
	addCmd.MarkFlagRequired("email")
}
