package main

import (
	"context"
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var conn = getDbConnection()

	var id string
	var size int64
	var err = conn.QueryRow(context.Background(), "select id, size from test where size>3").Scan(&id, &size)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(id, size)
}
