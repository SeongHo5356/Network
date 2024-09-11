package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type MyHttpHandler struct{}

func (h MyHttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	

	switch r.Method {
	case "GET":
		fmt.Println("\n## do_GET() activated.")
		fmt.Println("::Client address   : ", strings.Split(r.RemoteAddr, ":")[0])
		fmt.Println("::Client port      : ", strings.Split(r.RemoteAddr, ":")[1])
		fmt.Println("::Request command  : ", r.Method)
		fmt.Println("::Request line     : ", r.RequestURI)
		fmt.Println("::Request path     : ", r.URL.Path)
		fmt.Println("::Request version  : ", r.Proto)
		h.handleGET(w, r)

	case "POST":
		fmt.Println("\n## do_POST() activated.")
		fmt.Println("::Client address   : ", strings.Split(r.RemoteAddr, ":")[0])
		fmt.Println("::Client port      : ", strings.Split(r.RemoteAddr, ":")[1])
		fmt.Println("::Request command  : ", r.Method)
		fmt.Println("::Request line     : ", r.RequestURI)
		fmt.Println("::Request path     : ", r.URL.Path)
		fmt.Println("::Request version  : ", r.Proto)
		h.handlePOST(w, r)
	}
}

func (h MyHttpHandler) sendHTTPResponseHeader(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")
}

func (h MyHttpHandler) handleGET(w http.ResponseWriter, r *http.Request) {
	h.sendHTTPResponseHeader(w)
	if strings.Contains(r.RequestURI, "?") {
		routine := strings.Split(r.RequestURI, "?")[1]
		parameter := h.parameterRetrieval(routine)
		result := h.simpleCalc(parameter[0], parameter[1])
		response := fmt.Sprintf("<html>GET request for calculation => %d x %d = %d</html>", parameter[0], parameter[1], result)
		w.Write([]byte(response))
		fmt.Printf("## GET request for calculation => %d x %d = %d\n", parameter[0], parameter[1], result)
	} else {
		w.Write([]byte(fmt.Sprintf("<html><p>HTTP Request GET for Path: %s</p></html>", r.URL.Path)))
		fmt.Printf("## GET request for directory => %s\n", r.URL.Path)
	}
}

func (h MyHttpHandler) handlePOST(w http.ResponseWriter, r *http.Request) {
	h.sendHTTPResponseHeader(w)
	postData, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	parameter := h.parameterRetrieval(string(postData))
	result := h.simpleCalc(parameter[0], parameter[1])
	response := fmt.Sprintf("POST request for calculation => %d x %d = %d", parameter[0], parameter[1], result)
	w.Write([]byte(response))
	fmt.Printf("## POST request data => %s\n", string(postData))
	fmt.Printf("## POST request for calculation => %d x %d = %d\n", parameter[0], parameter[1], result)
}

func (h MyHttpHandler) send_http_response_header() {

}

func (h MyHttpHandler) simpleCalc(para1, para2 int) int {
	return para1 * para2
}

func (h MyHttpHandler) parameterRetrieval(msg string) []int {
	//var result []int
	result := make([]int, 2)
	fields := strings.Split(msg, "&")
	for i, field := range fields {
		parts := strings.Split(field, "=")
		if len(parts) == 2 {
			value, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Error:", err)
				return result
			}
			result[i] = value
		}
	}
	return result
}

func main() {
	server_name := "localhost"
	server_port := 8080
	address := fmt.Sprintf("%s%s%d", server_name, ":", server_port)

	http.Handle("/", MyHttpHandler{})

	fmt.Printf("HTTP server started at http://%s:%d\n", server_name, server_port)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}

}
