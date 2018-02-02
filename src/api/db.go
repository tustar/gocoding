package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	MaxOpenConns = 2000
	MaxIdleConns = 1000

	CreateFolderSql = `CREATE TABLE folder (
		_id INTEGER PRIMARY KEY,
			app TEXT NOT NULL,
			_data TEXT NOT NULL,
			description TEXT,
			depth INTEGER NOT NULL,
		type INTEGER NOT NULL);`
	QueryFolderSql  = `SELECT * FROM folder WHERE _id = ?`
	QueryFoldersSql = `SELECT * FROM folder`
)

var db *sql.DB
var folders []Folder

func init() {
	db, _ = sql.Open("mysql", "root:tustar@/filemanager")
	db.SetMaxOpenConns(MaxOpenConns)
	db.SetMaxIdleConns(MaxIdleConns)
	err := db.Ping()
	panicErr(err)
}

func queryFolder(id int) (Folder) {
	// Prepare statement for reading data
	stmtOut, err := db.Prepare(QueryFolderSql)
	panicErr(err)
	defer stmtOut.Close()

	rows, err := stmtOut.Query(id)
	panicErr(err)

	defer rows.Close()
	var index int64
	var app string
	var data string
	var description string
	var depth int8
	var ttype int8

	for rows.Next() {
		err := rows.Scan(&index, &app, &data, &description, &depth, &ttype)
		panicErr(err)
	}

	err = rows.Err()
	panicErr(err)

	folder := Folder{ID: index, App: app, Data: data,
		Description: description, Depth: depth, Ttype: ttype}
	return folder
}

func queryFolders() ([]Folder) {

	rows, err := db.Query(QueryFoldersSql)
	panicErr(err)

	defer rows.Close()
	var index int64
	var app string
	var data string
	var description string
	var depth int8
	var ttype int8

	for rows.Next() {
		err := rows.Scan(&index, &app, &data, &description, &depth, &ttype)
		panicErr(err)

		folders = append(folders, Folder{ID: index, App: app, Data: data,
			Description: description, Depth: depth, Ttype: ttype})
	}

	err = rows.Err()
	panicErr(err)

	return folders
}
