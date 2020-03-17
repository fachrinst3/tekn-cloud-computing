package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Tcc struct {
	Kegiatan string
}

func main() {

	db, err := sql.Open("mysql", "admin:root@tcp(127.0.0.1:3306)/tcc6")

	// handle error
	if err != nil {
		panic(err.Error())
	}

	// defer kalo main function uda kelar
	defer db.Close()

	fmt.Println("Koneksi Mysql database sukses")

	results, err := db.Query("SELECT * FROM tcc")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var tcc Tcc //inisiasi objek

		err = results.Scan(&tcc.Kegiatan)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(tcc.Kegiatan) //print atribut tcc
	}

}
