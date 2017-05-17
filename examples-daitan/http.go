package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type fixer struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res, err := http.Get("http://api.fixer.io/latest")
		if err != nil {
			fmt.Println(err)
			return
		}

		var f fixer
		if err := json.NewDecoder(res.Body).Decode(&f); err != nil {
			fmt.Println(err)
			return
		}

		if err := json.NewEncoder(w).Encode(f); err != nil {
			fmt.Println(err)
		}
	})
	http.ListenAndServe(":8082", nil)
}
