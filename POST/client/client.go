package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RequestData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// Server URL
	url := "http://localhost:8080/submit"

	// Request data
	requestData := RequestData{
		Name:  "Mahindra",
		Email: "mahindra@example.com",
	}

	// Convert struct to JSON
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// Send POST request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Print the response
	fmt.Println("Response from server:", string(body))
}
