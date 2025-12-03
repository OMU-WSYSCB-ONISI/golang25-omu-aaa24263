package main

import (
	"fmt"
	"net/http"
	"time"
	"math/rand"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
    fmt.Println("Week 03 課題")

	http.HandleFunc("/webfortune", webfortunehandler)

	http.ListenAndServe(":8080", nil)
}

func webfortunehandler(w http.ResponseWriter, r *http.Request) {
	fortunes := []string{"大吉", "中吉", "吉", "小吉", "末吉", "凶"}

	index := rand.Intn(len(fortunes))
	result := fortunes[index]

	fmt.Fprintf(w, "今日の運勢は**%s**です！", result)
}
