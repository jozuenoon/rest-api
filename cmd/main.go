package main

import (
	"flag"
	"github.com/jozuenoon/rest-api/cmd/payment"
	"golang.org/x/net/context"
)

var (
	databasePath = flag.String("database-path", "db/payment.store", "specify database path")
)

func main() {
	flag.Parse()
	srv := &payment.Server{
		DatabasePath: *databasePath,
	}
	srv.Run(context.Background())
}
