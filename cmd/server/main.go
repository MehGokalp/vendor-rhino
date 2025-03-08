package server

import (
	"github.com/mehgokalp/vendor-rhino/internal/card/factory"
	"github.com/mehgokalp/vendor-rhino/internal/repository"
	"github.com/mehgokalp/vendor-rhino/internal/server"
	"github.com/spf13/cobra"
)

func Server(
	currencyRepository *repository.CurrencyRepository,
	cardRepository *repository.CardRepository,
	cardFactory *factory.CardFactory,
) *cobra.Command {
	cmdName := "server"

	return &cobra.Command{
		Use:   cmdName,
		Short: "Run backend server",
		RunE: func(cmd *cobra.Command, _ []string) error {
			r := server.GetRouter(
				currencyRepository,
				cardRepository,
				cardFactory,
			)

			if err := r.Run(); err != nil {
				return err
			}

			return nil
		},
	}
}
