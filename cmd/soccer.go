package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/Ferza17/soccer-api/graphql"
)

var graphqlCommand = &cobra.Command{
	Use: "soccer",
	Run: func(cmd *cobra.Command, args []string) {
		chClose := make(chan os.Signal, 1)
		signal.Notify(chClose, syscall.SIGINT, syscall.SIGTERM)
		go func() {
			log.Println("Starting GraphQL API server ...")
			graphql.NewServer(
				os.Getenv("CODENAME"),
				os.Getenv("HTTP_HOST"),
				os.Getenv("HTTP_PORT"),
				graphql.NewLogger(logger),
				graphql.NewTeamsDB(teamsDB),
				graphql.NewPlayersDB(playersDB),
			).Serve()
		}()
		<-chClose
		if err := Shutdown(context.Background()); err != nil {
			log.Fatalln(err)
			return
		}
		log.Println("Exit...")
	},
}
