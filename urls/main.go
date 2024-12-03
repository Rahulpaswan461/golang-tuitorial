package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("learning about urls")
	myURL := "https://example.com:8080/path/to/resource?key1=value1&key2=value2"
	fmt.Printf("the type is %T\n", myURL)

	parsedURL, err := url.Parse(myURL)
	if err != nil {
		fmt.Println("Error while parsing the URL", err)
		return
	}
	fmt.Println("parsed url is", parsedURL)
	fmt.Printf("type of the url is %T\n", parsedURL)

	fmt.Println("host is", parsedURL.Host)
	fmt.Println("scheme is", parsedURL.Scheme)
	fmt.Println("path is : ", parsedURL.Path)
	fmt.Println("query is : ", parsedURL.RawQuery)
	parsedURL.Host = "2000"

	newString := parsedURL.String()
	fmt.Println("new string is", newString)
}
