package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Liste tous les contacts",
	Run: func(cmd *cobra.Command, args []string) {
		contacts := Store.Lister()
		for _, c := range contacts {
			fmt.Printf("ID=%d | Nom=%s | Email=%s\n", c.ID, c.Nom, c.Email)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
