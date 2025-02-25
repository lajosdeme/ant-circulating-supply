package internal

import (
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lajosdeme/autonomi-circulating-supply/config"
)

type Response struct {
	CirculatingSupply float64 `json:"circulating_supply"`
}

const AntCirculatingSupplyAddress = "0xA8C9d2521c0201E60Aa7E581876FbB1397eC0E58"

func Start() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Only allow GET requests
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		circulatingSupply, err := getCirculatingSupply()
		if err != nil {
			log.Printf("Error encoding response: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}

		supplyFloat, err := strconv.ParseFloat(circulatingSupply, 64)
		if err != nil {
			log.Printf("Error encoding response: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}

		// Create response with the fixed data
		response := Response{
			CirculatingSupply: supplyFloat,
		}

		// Set content type header
		w.Header().Set("Content-Type", "application/json")

		// Encode the response to JSON
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	})

	// Start the server on port 3000
	port := ":3000"
	log.Printf("Starting server on port %s...", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func getCirculatingSupply() (string, error) {
	client, err := ethclient.Dial(config.C.ArbRpcUrl)
	if err != nil {
		return "", err
	}

	antCirculatingSupply, err := NewAntCirculatingSupply(common.HexToAddress(AntCirculatingSupplyAddress), client)
	if err != nil {
		return "", err
	}

	totalCirculatingSupply, err := antCirculatingSupply.TotalCirculatingSupply(nil)
	if err != nil {
		return "", err
	}

	circulatingSupplyFloat := weiToEther(totalCirculatingSupply)

	return circulatingSupplyFloat.Text('f', 18), nil
}

func weiToEther(wei *big.Int) *big.Float {
	// Define 10^18 as a big.Int
	divisor := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)

	// Convert wei to Ether by dividing by 10^18
	ether := new(big.Float).SetInt(wei)
	ether.Quo(ether, new(big.Float).SetInt(divisor))

	return ether
}
