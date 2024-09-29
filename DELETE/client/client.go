package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DeleteData struct {
	ID string `json:"id"`
}

func main() {
	// Server URL
	url := "http://localhost:8080/delete"

	// Request data
	deleteData := DeleteData{
		ID: "12345",
	}

	// Convert struct to JSON
	requestBody, err := json.Marshal(deleteData)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// Create a DELETE request
	req, err := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// Send DELETE request
	client := &http.Client{}
	resp, err := client.Do(req)
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
