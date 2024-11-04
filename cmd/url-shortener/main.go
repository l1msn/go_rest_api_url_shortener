package main

import (
	"fmt"
	"url_shortner/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO: init logger: slog in go
	// TODO: init storage: sqllite
	// TODO: init router: chi + net/http, render
	// TODO: run server
}
