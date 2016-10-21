package dealrepo

import (
	"github.com/fahric/dataService/config"
	"github.com/fahric/dataService/models"
	"log"
	"fmt"
)

func GetDeals(resultChn chan []*models.Deal){
	var result []*models.Deal
	rows, err := config.Connect().Query("SELECT TOP 1 sysid,description FROM dbo.tblDeal ")
	defer config.Disconnect()
	if err != nil {
		fmt.Println("[dealRepo]Error on connect")
		config.Disconnect()

		log.Fatal(err)
	}
	for rows.Next() {
		r := new(models.Deal)
		if err := rows.Scan(&r.Id,&r.Description); err != nil {
			log.Fatal(err)
			resultChn<-result
		}
		result = append(result,r)

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
		resultChn<-result
	}

	resultChn<-result
}

func GetDeal(id int,resultChn chan *models.Deal){

	rows := config.Connect().QueryRow("SELECT sysid,description FROM dbo.tblDeal WHERE sysid = $1 ",id)
	defer config.Disconnect()

	r := new(models.Deal)
	err := rows.Scan(&r.Id,&r.Description);

	if err != nil {
		config.Disconnect()
		log.Fatal(err)
	}

	resultChn<-r
}
