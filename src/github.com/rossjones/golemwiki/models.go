package main

import (
	"database/sql"
	"fmt"
	_ "github.com/bmizerany/pq"
)

var db *sql.DB

func init_database() {
	var err error

	cstring := fmt.Sprintf("user=%s dbname=%s sslmode=disable",
		config.Database["user"],
		config.Database["dbname"])
	/*
	   "password": "",
	   "host": "",
	   "port": ""
	*/

	db, err = sql.Open("postgres", cstring)
	if err != nil {
		panic("Failed to connect to the database")
	}
}

func get_record_count(query string) int {
	count := 0
	row := db.QueryRow(query)
	err := row.Scan(&count)
	if err != nil {
		fmt.Println("Failed to get count - " + query)
		return -1
	}
	return count
}

func get_all_groups() []Group {
	count := get_record_count("select count(id) from golem_group;")
	it := 0

	rows, err := db.Query("select id, name, slug from golem_group;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	vals := make([]Group, count)
	for rows.Next() {
		g := Group{}
		err = rows.Scan(&g.Id, &g.Name, &g.Slug)
		if err != nil {
			panic("Failed to run query to get groups")
		}
		vals[it] = g
		it += 1
	}

	return vals
}
