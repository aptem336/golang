package main

import (
	"fmt"
	"golang/service/catapulto"
	"golang/service/catapulto/api"
	"os"
)

func main() {
	client := catapulto.NewClient(
		"https://api.catapulto.ru/api/v1",
		"",
	)
	localityList, err := client.GetLocalityList(&api.LocalityRequest{
		Term: "Казань",
		ISO:  "RU",
	})
	if err != nil {
		fmt.Printf("can't get localityList: %s", err)
		os.Exit(1)
	}
	fmt.Print(localityList)
}
