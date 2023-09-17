package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

var AllCommand = &cobra.Command{
	Use:   "all",
	Short: "get all orders ID",
	Long:  "get all orders ID",
	Run: func(cmd *cobra.Command, args []string) {
		url := fmt.Sprintf("%s/", URL_API) // Remplacez par l'URL de votre API
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Erreur lors de la requête à l'API:", err)
			return
		}
		defer resp.Body.Close()
		//body, err := io.ReadAll(resp.Body)
		//fmt.Println("Body de la reponse:")
		//fmt.Println(string(body))
		if resp.StatusCode != http.StatusOK {
			fmt.Println("La requête a échoué avec le code de statut:", resp.Status)
			return
		}

		var items Items
		body, err := io.ReadAll(resp.Body)
		err = json.Unmarshal(body, &items)
		if err != nil {
			fmt.Println("Erreur lors de la lecture de la réponse JSON:", err)
			return
		}

		fmt.Println("JSON Brute", items)
		if len(items.Orders) == 0 {
			fmt.Println("il n'y a aucune commande !")
		}
		for _, order := range items.Orders {
			fmt.Printf("ID de commande : %d\n", order.OrderID)
		}

	},
}
