package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type fixer struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}

func readFixer(ch chan []byte) {
	res, err := http.Get("http://api.fixer.io/latest")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ch <- b
}

func main() {
	start := time.Now()

	ch := make(chan []byte)
	for i := 0; i < 5; i++ {
		go readFixer(ch)
	}

	for i := 0; i < 5; i++ {
		b := <-ch
		var f fixer
		if err := json.Unmarshal(b, &f); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(f)
	}
	fmt.Println("took", time.Since(start))
}
