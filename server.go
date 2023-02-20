package main

import (
	"github.com/bruceneco/dicio-api/bootstrap"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	bootstrap.RootApp.Execute()
}
