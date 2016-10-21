package main

import (

	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
	"github.com/fahric/dataService/config"
	"database/sql"
)
type deal struct{
	id int
	description string
}
func main() {
	resultChn := make(chan []*deal)
	go getADeal(resultChn)
	result := <-resultChn
	for _,d := range result{
		fmt.Printf("Deal id is %d and description is %s\n", d.id,d.description)

	}

}

func getADeal(resultChn chan []*deal){
	var result []*deal
	db, err := sql.Open("mssql", config.GetConnection())
	if err != nil {
		fmt.Println(err)
		resultChn<-result
	}
	rows, err := db.Query("SELECT TOP 1 sysid,description FROM dbo.tblDeal ")

	if err != nil {
		log.Fatal(err)
		resultChn<-result
	}
	for rows.Next() {
		r := new(deal)
		if err := rows.Scan(&r.id,&r.description); err != nil {
			log.Fatal(err)
			resultChn<-result
		}
		result = append(result,r)

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
		resultChn<-result
	}
	defer db.Close()
	resultChn<-result
}
