package main

import (
	"github.com/jinzhu/gorm"
	"github.com/mehgokalp/vendor-rhino/cmd/server"
	"github.com/mehgokalp/vendor-rhino/internal/card/factory"
	"github.com/mehgokalp/vendor-rhino/internal/repository"
	"github.com/mehgokalp/vendor-rhino/migrations"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "vendor-rhino",
		Short: "Main entry-point command for the application",
	}

	db, err := gorm.Open("mysql", os.Getenv("VENDOR_RHINO_DATABASE_DNS"))
	if err != nil {
		panic(err)
	}
	migrations.AutoMigrate(db)

	currencyRepository := repository.NewCurrencyRepository(db)
	cardRepository := repository.NewCardRepository(db)
	cardFactory := factory.NewCardFactory()

	rootCmd.AddCommand(
		server.Server(
			currencyRepository,
			cardRepository,
			cardFactory,
		),
	)
}
