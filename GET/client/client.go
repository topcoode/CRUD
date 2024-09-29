package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// Server URL
	baseURL := "http://localhost:8080/data"

	// Data to send as query parameters
	params := url.Values{}
	params.Add("name", "Mahindra")
	params.Add("email", "mahindra@example.com")

	// Build the full URL with query parameters
	fullURL := baseURL + "?" + params.Encode()

	// Send GET request
	resp, err := http.Get(fullURL)
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
