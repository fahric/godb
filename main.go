package main

import (

	"fmt"
	"github.com/fahric/dataService/models"
	"github.com/fahric/dataService/dealrepo"
)

func main() {
	singleResultChn := make(chan *models.Deal)
	multipleResultChn := make(chan []*models.Deal)
	go dealrepo.GetDeal(607867,singleResultChn)
	go dealrepo.GetDeals(multipleResultChn)

	singleDeal := <-singleResultChn

	fmt.Printf("Deal id is %d and description is %s\n", singleDeal.Id,singleDeal.Description)
	multipleDeal := <-multipleResultChn

	for _,d := range multipleDeal{
		fmt.Printf("Deal id is %d and description is %s\n", d.Id,d.Description)
	}

}

