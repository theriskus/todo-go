package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

//Quest is a element from todo list
type Quest struct {
	ID     int    `json:"id"`
	Qu     string `json:"quest"`
	Status int    `json:"status"`
}

func indexQuest(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	quests, err := db.Query("select * from quests")
	if err != nil {
		panic(err)
	}
	defer quests.Close()
	qsts := make([]*Quest, 0)
	for quests.Next() {
		qt := new(Quest)
		err := quests.Scan(&qt.ID, &qt.Qu, &qt.Status)
		if err != nil {
			log.Fatal(err)
		}
		qsts = append(qsts, qt)
	}
	b, err := json.Marshal(qsts)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(b))
}

func deleteQuest(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	decoder := json.NewDecoder(r.Body)
	var res Quest
	err := decoder.Decode(&res)
	if err != nil {
		panic(err)
	}
	db.Query("delete from quests where id=?", res.ID)
	fmt.Fprintf(w, "[{\"status\": true}]")
}
func updateQuest(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	decoder := json.NewDecoder(r.Body)
	var res Quest
	err := decoder.Decode(&res)
	if err != nil {
		panic(err)
	}
	db.Query("update quests set status = ? where id=?", res.Status, strconv.FormatInt(int64(res.ID), 10))
	fmt.Fprintf(w, "[{\"status\": true}]")
}

func createQuest(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	decoder := json.NewDecoder(r.Body)
	var res Quest
	err := decoder.Decode(&res)
	if err != nil {
		panic(err)
	}
	db.Query("insert into quests set quest = ?, status = ?", res.Qu, res.Status)
	fmt.Fprintf(w, "[{\"status\": true}]")
}
