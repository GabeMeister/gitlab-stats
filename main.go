package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	apiToken := os.Getenv("GITLAB_API_TOKEN")
	fmt.Print("\n\n", "*** yessir ***", "\n", apiToken, "\n\n\n")
}
