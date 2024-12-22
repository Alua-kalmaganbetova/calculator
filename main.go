package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"your_project/internal/calculator"
	"errors"
)

type request struct {
	Expression string `json:"expression"`
}

type response struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req request
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	
	if !calculator.IsValidExpression(req.Expression) {
		http.Error(w, "Expression is not valid", http.StatusUnprocessableEntity)
		return
	}

	// Вычисление результата
	result, err := calculator.Calc(req.Expression)
	if err != nil {
		if errors.Is(err, calculator.ErrInvalidExpression) {
			http.Error(w, "Expression is not valid", http.StatusUnprocessableEntity)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	resp := response{Result: fmt.Sprintf("%f", result)}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/api/v1/calculate", calculateHandler)
	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
