package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var boxCommand = &cobra.Command{
	Use: "box",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			cake  = 20
			apple = 25
		)
		fmt.Println("Cake : ", cake)
		fmt.Println("Apple : ", apple)
	},
}
