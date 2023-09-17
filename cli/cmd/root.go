package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
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
		fmt.Println("Command:")
		fmt.Println("get [id]")
	},
}

func Execute() {
	rootCmd.Root().AddCommand(getCommand)
	rootCmd.Root().AddCommand(AllCommand)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

type Order struct {
	OrderID     uint64     `json:"order_id"`
	CustomerID  uuid.UUID  `json:"customer_id"`
	LineItems   []LineItem `json:"line_items"`
	CreatedAt   *time.Time `json:"created_at"`
	ShippedAt   *time.Time `json:"shipped_at"`
	CompletedAt *time.Time `json:"completed_at"`
}

type LineItem struct {
	ItemID   uuid.UUID `json:"item_id"`
	Quantity uint      `json:"quantity"`
	Price    uint      `json:"price"`
}

type Items struct {
	Orders []Order `json:"items"`
}
