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
	authResponse, err := client.GetAuthToken(&api.AuthRequest{
		Username: "",
		Password: "",
	})
	if err != nil {
		fmt.Printf("can't get localityList: %s", err)
		os.Exit(1)
	}
	fmt.Print(authResponse)
	localityList, err := client.GetLocalityList(&api.LocalityRequest{
		Term:  "Казань",
		ISO:   "RU",
		Limit: 5,
	})
	if err != nil {
		fmt.Printf("can't get localityList: %s", err)
		os.Exit(1)
	}
	fmt.Print(localityList)
}
