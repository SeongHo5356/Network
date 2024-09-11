package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("## HTTP client started.")

	fmt.Println("## GET request for http://localhost:8080/temp/")
	getRequest("http://localhost:8080/temp/")

	fmt.Println("## GET request for http://localhost:8080/?var1=9&var2=9")
	getRequest("http://localhost:8080/?var1=9&var2=9")

	fmt.Println("## POST request for http://localhost:8080/ with var1 is 9 and var2 is 9")
	postRequest("http://localhost:8080", map[string]string{"var1": "9", "var2": "9"})

	fmt.Println("## HTTP client completed.")
}

func getRequest(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("## GET response [start]")
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(data))
	fmt.Println("## GET response [end]\n")
}

func postRequest(URL string, data map[string]string) {

	values := url.Values{}
	for key, value := range data {
		values.Add(key, value)
	}

	response, err := http.PostForm(URL, values)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer response.Body.Close()

	fmt.Println("## POST response [start]")
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(body))
	fmt.Println("## POST response [end]\n")
}
