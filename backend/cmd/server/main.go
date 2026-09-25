package main

import (
	"github.com/Bibintanggg/forge/internal/api"
)

func main() {
	router := api.NewRouter()
	router.Run(":8080")

}
