package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var getCommand = &cobra.Command{
	Use:   "get",
	Short: "get oder info by id",
	Long:  "get order info and status by id",
	Run: func(cmd *cobra.Command, args []string) {
		orderID := args[0]
		url := fmt.Sprintf("%s/%s", URL_API, orderID) // Remplacez par l'URL de votre API
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Erreur lors de la requête à l'API:", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Println("La requête a échoué avec le code de statut:", resp.Status)
			return
		}

		var order Order
		err = json.NewDecoder(resp.Body).Decode(&order)
		if err != nil {
			fmt.Println("Erreur lors de la lecture de la réponse JSON:", err)
			return
		}

		fmt.Printf("ID de commande : %d\n", order.OrderID)
		fmt.Printf("ID du client : %s\n", order.CustomerID)
		fmt.Printf("Créé le : %s\n", order.CreatedAt)

		if order.ShippedAt != nil {
			fmt.Printf("Expédié le : %s\n", *order.ShippedAt)
		} else {
			fmt.Println("Pas encore expédié")
		}

		if order.CompletedAt != nil {
			fmt.Printf("Complété le : %s\n", *order.CompletedAt)
		} else {
			fmt.Println("Pas encore complété")
		}

		fmt.Println("Articles :")
		for _, item := range order.LineItems {
			fmt.Printf("- ID de l'article : %s\n", item.ItemID)
			fmt.Printf("  Quantité : %d\n", item.Quantity)
			fmt.Printf("  Prix : %d\n", item.Price)
		}
	},
}
