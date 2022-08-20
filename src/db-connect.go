package main

import (
	//"fmt"
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func db_exec (dbPath string, sqlStatement string) {
	db, _ := sql.Open("sqlite3", dbPath)
	defer db.Close()
  
	_, err := db.Exec(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
}

func db_select(dbPath string) (dbHosts []Host) {
	db, _ := sql.Open("sqlite3", dbPath)
	defer db.Close()

	sqlStatement := `SELECT * FROM "now"`

	res, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}

	dbHosts = []Host{}
  	for res.Next() {
    	var oneHost Host
    	err = res.Scan(&oneHost.Id, &oneHost.Name, &oneHost.Ip, &oneHost.Mac, &oneHost.Hw, &oneHost.Date, &oneHost.Known, &oneHost.Now)
    	if err != nil {
			log.Fatal(err)
    	}

    	dbHosts = append(dbHosts, oneHost)
  	}

	//fmt.Println("Select all:", dbHosts)
	return dbHosts
}