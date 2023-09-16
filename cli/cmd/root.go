package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "orders-cli",
	Short: "Very useful CLI to manipulate your orders",
	Long: `A CLI to manipulate the orders api
				this CLI is build with Cobra and http client
				you can list, add, search, modify and delete 
				orders`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("welcome to the Orders cli")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
