package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/simplebank/api"
	db "github.com/simplebank/db/sqlc"
	"github.com/spf13/viper"
)

const (
	driver         = "postgres"
	databaseSource = "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable"
	serverAddress  = "0.0.0.0:8090"
)

func main() {
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("can't load config", err)
	}
	conn, err := sql.Open(viper.GetString("DRIVER"), viper.GetString("DATABASESOURCE"))
	if err != nil {
		log.Fatal("can't connect to db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(viper.GetString("SERVERADDRESS"))
	if err != nil {
		log.Fatal("can't start server", err)
	}
}
