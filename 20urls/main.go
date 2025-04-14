package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://www.google.com/search?q=gg&oq=gg&gs_lcrp=EgZjaHJvbWUqDggAEEUYJxg7GIAEGIoFMg4IABBFGCcYOxiABBiKBTINCAEQLhiDARixAxiABDIQCAIQABiDARixAxiABBiKBTINCAMQLhiDARixAxiABDIHCAQQABiABDIGCAUQRRg8MgYIBhBFGDwyBggHEEUYPNIBBzY1M2owajeoAgCwAgA&sourceid=chrome&ie=UTF-8"

func main() {
	fmt.Println("Hello, to the URLs!")
	fmt.Println("The URL is: ", myurl)

	// Parsing the URL
	result, _ := url.Parse(myurl)
	fmt.Println("The Scheme of the result is: ", result.Scheme)
	fmt.Println("The Host of the result is: ", result.Host)
	fmt.Println("The Path of the result is: ", result.Path)
	fmt.Println("The Port of the result is: ", result.Port())
	fmt.Println("The RawQuery of the result is: ", result.RawQuery)

	queryParams := result.Query()
	fmt.Println("The Query Parameters are: ", queryParams)

}