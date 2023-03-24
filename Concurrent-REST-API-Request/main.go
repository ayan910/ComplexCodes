// You can edit this code!
// Click here and start typing.https://api.chucknorris.io/jokes/random
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Data struct {
	// Categories string `json:"categories"`
	// Created_at string `json:"created_at"`
	// Icon_url   string `json:"icon_url"`
	// Id         string `json:"id"`
	// Updated_at string `json:"updated_at"`
	// Url        string `json:"url"`
	Value string `json:"value"`
}

func main() {
	start := time.Now()
	var n int = 10000
	var url string = "https://api.chucknorris.io/jokes/random"
	result, err := getValues(n, url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
	timeElapsed := time.Since(start)
	fmt.Printf("Time Taken: %s", timeElapsed)
}

func callEndPoint(url string) (Data, error) {

	var data Data
	myClient := &http.Client{}
	resp, err := myClient.Get(url)
	if err != nil {
		// fmt.Println(err)
		return Data{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Data{}, err
	}
	// fmt.Println(string(po))
	if err := json.Unmarshal(body, &data); err != nil {
		// fmt.Println(err)
		return Data{}, nil
	}
	// fmt.Println(data.Value)
	return data, nil
}

func getValues(n int, url string) ([]string, error) {

	var responseArray []string
	dataChannel := make(chan string)
	for i := 0; i < n; i++ {

		go func() {
			data, err := callEndPoint(url)
			if err != nil {
				fmt.Print(err)
			}
			dataChannel <- data.Value
		}()
	}

	for i := 0; i < n; i++ {
		responseArray = append(responseArray, <-dataChannel+"\n")
	}

	return responseArray, nil
}
