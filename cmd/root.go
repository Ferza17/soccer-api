package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCommand.AddCommand(graphqlCommand, boxCommand)
}

var rootCommand = &cobra.Command{
	Use: "root",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("running root command...")
	},
}

func Run() {
	if err := rootCommand.Execute(); err != nil {
		log.Panic(err)
	}
}
