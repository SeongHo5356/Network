package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// create

func postRequest(url, data string) {
	response, err := http.Post(url, "application/x-www-form-urlencoded", bytes.NewBufferString(data))
	if err != nil {
		fmt.Println("POST", "request error:", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Printf("#%s Code: %d >> JSON: %s >> JSON Result: %s\n", "POST", response.StatusCode, string(body), string(body))

}

// read

func getRequest(url string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("GET", "request error:", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Printf("#%s Code: %d >> JSON: %s >> JSON Result: %s\n", "GET", response.StatusCode, string(body), string(body))

}

// update

func putRequest(url, data string) {
	request, err := http.NewRequest(http.MethodPut, url, bytes.NewBufferString(data))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		fmt.Println("PUT", "request error:", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Printf("#%s Code: %d >> JSON: %s >> JSON Result: %s\n", "PUT", response.StatusCode, string(body), string(body))

}

// delete

func deleteRequest(url string) {
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		fmt.Println("DELETE", "request error:", err)
		return
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Printf("#%s Code: %d >> JSON: %s >> JSON Result: %s\n", "Delete", response.StatusCode, string(body), string(body))
}

func main() {
	// Reads a non-registered member: error-case
	getRequest("http://127.0.0.1:5000/membership_api/0001")

	// Creates a new registered member: non-error case
	postRequest("http://127.0.0.1:5000/membership_api/0001", "0001=apple")

	// Reads a registered member: non-error case
	getRequest("http://127.0.0.1:5000/membership_api/0001")

	// Creates an already registered member: error case
	postRequest("http://127.0.0.1:5000/membership_api/0001", "0001=xpple")

	// Updates a non-registered member: error case
	putRequest("http://127.0.0.1:5000/membership_api/0002", "0002=xrange")

	// Updates a registered member: non-error case
	postRequest("http://127.0.0.1:5000/membership_api/0002", "0002=xrange")
	putRequest("http://127.0.0.1:5000/membership_api/0002", "0002=orange")

	// Deletes a registered member: non-error case
	deleteRequest("http://127.0.0.1:5000/membership_api/0001")

	// Deletes a non-registered member: non-error case
	deleteRequest("http://127.0.0.1:5000/membership_api/0001")
}
