package bitpanda

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// TODO Save in .env environment
const BP_API_KEY_NAME = "BITPANDA_API_KEY"

type Time struct {
	DateISO8601 string `json:"date_iso8601"`
	Unix        string `json:"unix"`
}

type Attributes struct {
	Amount                     string `json:"amount"`
	Recipient                  string `json:"recipient"`
	Time                       Time   `json:"time"`
	Confirmations              int    `json:"confirmations"`
	InOrOut                    string `json:"in_or_out"`
	Type                       string `json:"type"`
	Status                     string `json:"status"`
	AmountEUR                  string `json:"amount_eur"`
	PurposeText                string `json:"purpose_text"`
	RelatedWalletTransactionID string `json:"related_wallet_transaction_id"`
	RelatedWalletID            string `json:"related_wallet_id"`
	WalletID                   string `json:"wallet_id"`
	Confirmed                  bool   `json:"confirmed"`
	CryptocoinID               string `json:"cryptocoin_id"`
	LastChanged                Time   `json:"last_changed"`
	Fee                        string `json:"fee"`
	CurrentFiatID              string `json:"current_fiat_id"`
	CurrentFiatAmount          string `json:"current_fiat_amount"`
	TxID                       string `json:"tx_id"`
}

type Transaction struct {
	Type       string     `json:"type"`
	Attributes Attributes `json:"attributes"`
	ID         string     `json:"id"`
}

type Meta struct {
	TotalCount int    `json:"total_count"`
	Cursor     string `json:"cursor"`
	NextCursor string `json:"next_cursor"`
	PageSize   int    `json:"page_size"`
}

type Links struct {
	Next string `json:"next"`
	Self string `json:"self"`
}

type Response struct {
	Data  []Transaction `json:"data"`
	Meta  Meta          `json:"meta"`
	Links Links         `json:"links"`
}

// TODO Get all transactions

func GetTransactions(cursor string) (*Response, error) {
	byteResponse, err := getTransactions(cursor)
	if err != nil {
		return nil, err
	}
	return parseResponse(byteResponse)
}

func getTransactions(cursor string) ([]byte, error) {
	url := "https://api.bitpanda.com/v1/wallets/transactions"

	if cursor != "" {
		url = fmt.Sprintf("%s?cursor=%s", url, cursor)
	}

	apiKey, ok := os.LookupEnv(BP_API_KEY_NAME)
	if !ok {
		return nil, fmt.Errorf("Could not find API Key %s in environment variables", BP_API_KEY_NAME)
	}

	fmt.Println(url)

	// Create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Add("X-Api-Key", apiKey)

	// Create client and send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	return body, nil
}

// parseResponse parses the JSON response into the Response struct
func parseResponse(jsonData []byte) (*Response, error) {
	var response Response
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}
	return &response, nil
}
