package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/maxhenkes/blobber-storage/api"
)

func main() {

	api.EnableUploadRoute()

	// var conn = getDbConnection()
	// var id int64
	// var err = conn.QueryRow(context.Background(), "select id from buckets").Scan(&id)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	// 	os.Exit(1)
	// }

	// fmt.Println(id)
	// conn.Close(context.Background())
}
