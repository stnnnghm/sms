package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/vonage/vonage-go-sdk"
)

// Can also import ../godotenv/autoloader
// then _ = gotodenv.Load(".envFile1", ".envFile2")
// https://github.com/joho/godotenv

var (
	API_KEY,
	API_SECRET,
	FROM_NUMBER,
	TO_NUMBER string
)

func init() {
	var myEnv map[string]string
	myEnv, err := godotenv.Read("VARS.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	API_KEY = myEnv["API_KEY"]
	API_SECRET = myEnv["API_SECRET"]
	FROM_NUMBER = myEnv["FROM_NUMBER"]
	TO_NUMBER = myEnv["TO_NUMBER"]
}
func main() {

	auth := vonage.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	smsClient := vonage.NewSMSClient(auth)

	response, err, _ := smsClient.Send(FROM_NUMBER, TO_NUMBER, "Hi there", vonage.SMSOpts{})
	if response.Messages[0].Status == "0" {
		fmt.Println("Messsage sent")
	} else {
		fmt.Printf("Error:%v\n", err.Messages[0])
	}
}
