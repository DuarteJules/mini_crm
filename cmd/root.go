package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/DuarteJules/mini_crm/internal/storage"
)

var (
	cfgFile string
	Store   storage.Storer
)

type Config struct {
	Storage struct {
		Type     string `mapstructure:"type"`
		File     string `mapstructure:"file"`
		Database string `mapstructure:"database"`
	} `mapstructure:"storage"`
}

var rootCmd = &cobra.Command{
	Use:   "mini-crm",
	Short: "Mini-CRM est une application CLI de gestion de contacts",
	Long:  `Mini-CRM est un outil en ligne de commande qui permet d'ajouter, lister, mettre à jour et supprimer des contacts.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		viper.SetDefault("type", "gorm")
		viper.SetDefault("file", "contacts.json")
		viper.SetDefault("database", "contacts.db")
		viper.SetConfigFile(cfgFile)
		if err := viper.ReadInConfig(); err != nil {
			return fmt.Errorf("erreur lecture config: %w", err)
		}

		var cfg Config
		if err := viper.Unmarshal(&cfg); err != nil {
			return fmt.Errorf("erreur parsing config: %w", err)
		}

		switch cfg.Storage.Type {
		case "json":
			Store = storage.NewJsonStore(cfg.Storage.File)
		case "gorm":
			db, err := gorm.Open(sqlite.Open(cfg.Storage.Database), &gorm.Config{})
			if err != nil {
				log.Fatal(err)
			}
			Store = storage.NewGormStore(db)
		default:
			return fmt.Errorf("type de stockage inconnu: %s", cfg.Storage.Type)
		}
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "config.yaml", "fichier de configuration")
	_ = viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
}
