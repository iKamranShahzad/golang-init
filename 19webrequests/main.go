package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://ikamranshahzad.vercel.app"

func main() {
	fmt.Println("Hello, World!")

	response, err := http.Get(url) // Get issues a GET to the specified URL. If the response is one of the following redirect codes, Get follows the redirect after calling the Client's CheckRedirect function. If the request is canceled or an error occurs, Get returns a non-nil error. A non-2xx response doesn't cause an error. Any returned error will be of type *url.Error. The returned response is always non-nil; it is up to the caller to close the response body when finished. Get is a wrapper around DefaultClient.Get.
	checkNil(err)
	fmt.Printf("The response is of type %T\n", response)

	defer response.Body.Close()

	databytes, err := io.ReadAll(response.Body) // ReadAll reads from r until an error or EOF and returns the data it read. A successful call returns err == nil, not err == EOF. Because ReadAll is defined to read from src until EOF, it does not treat an EOF from Read as an error to be reported.
	checkNil(err)
	fmt.Println("The Response's Body is\n", string(databytes))
	fmt.Printf("The Response's Body is of type %T\n", databytes)
}

func checkNil(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}
}