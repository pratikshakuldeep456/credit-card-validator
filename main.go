package main

import (
	"credit-card-validator/validator"
	"encoding/json"
	"fmt"
	"net/http"
)

type Request struct {
	CreditCard string `json:"card_number"`
}

type Response struct {
	Valid bool `json:"valid"`
}

func ValidateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

	var req Request

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, " error decoding", http.StatusBadRequest)
		return
	}

	cardNo := req.CreditCard
	validd := validator.ValidateNumber(cardNo)

	response := Response{
		Valid: validd,
	}

	json.NewEncoder(w).Encode(response)

}
func main() {

	//input card no , cvv, expiry date
	//validate expiry
	//detect card type ->algo for that

	cardNo := "799273987132113079e09443"

	ans := validator.ValidateNumber(cardNo)

	fmt.Println(ans)

	http.HandleFunc("/validate", ValidateHandler)

	fmt.Println("server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
