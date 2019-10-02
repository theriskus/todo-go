package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	getConfig()

	db = connect()
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(0)
	db.SetMaxOpenConns(50)

	flag.Parse()
	hub := newHub()
	go hub.run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	http.HandleFunc("/", indexQuest)
	http.HandleFunc("/index", indexQuest)
	http.HandleFunc("/delete", deleteQuest)
	http.HandleFunc("/update", updateQuest)
	http.HandleFunc("/create", createQuest)
	err := http.ListenAndServe(":9900", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
