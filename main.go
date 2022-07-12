package main

import "github.com/itsapep/golang-with-mongodb/delivery"

func main() {
	// MONGO_HOST=127.0.0.1 MONGO_PORT=27017 MONGO_DB=contohDb MONGO_USER=yurhamafif MONGO_PASSWORD=12345678 API_HOST=localhost API_PORT=8888 go run .
	delivery.NewServer().Run()
}
