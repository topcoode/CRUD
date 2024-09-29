package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type UpdateData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// Server URL
	url := "http://localhost:8080/update"

	// Request data
	updateData := UpdateData{
		Name:  "Mahindra",
		Email: "mahindra.new@example.com",
	}

	// Convert struct to JSON
	requestBody, err := json.Marshal(updateData)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// Create a PUT request
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// Send PUT request
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
