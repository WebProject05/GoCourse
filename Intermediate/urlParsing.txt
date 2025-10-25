package main

import (
	"fmt"
	"net/url"
)

func main() {
	// [scheme://][userindo@]host[:port][/path][?query][#fragment]
	rawUrl := "https://example.com:8000/path?query=param#fragment"
	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Parsed Url:", parsedUrl)
	fmt.Println("Parsed Url scheme:", parsedUrl.Scheme)
	fmt.Println("Parsed Url Host:", parsedUrl.Host)
	fmt.Println("Parsed Url Port:", parsedUrl.Port())
	fmt.Println("Parsed Url Path:", parsedUrl.Path)
	fmt.Println("Parsed Url Query:", parsedUrl.Query())
	fmt.Println("Parsed Url Fragment:", parsedUrl.Fragment)

	url1 := "https://example.com:8000/path?name=Santosh&source=git#fragment"
	parsedUrl2, err1 := url.Parse(url1)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(parsedUrl2)
	fmt.Println("Parsed Url scheme:", parsedUrl2.Scheme)
	fmt.Println("Parsed Url Host:", parsedUrl2.Host)
	fmt.Println("Parsed Url Port:", parsedUrl2.Port())
	fmt.Println("Parsed Url Path:", parsedUrl2.Path)
	fmt.Println("Parsed Url Query:", parsedUrl2.Query())
	fmt.Println("Parsed Url Fragment:", parsedUrl2.Fragment)

	queryParams := parsedUrl2.Query()
	fmt.Println(queryParams)
	fmt.Println(queryParams.Get("name"))
	fmt.Println(queryParams.Get("source"))

	// Building a URL
	baseUrl := &url.URL{
		Scheme: "https",
		Host: "udemy.com",
		Path: "/path",
	}

	fmt.Println(baseUrl)
	query := baseUrl.Query()
	query.Set("name", "santosh")
	baseUrl.RawQuery = query.Encode()

	fmt.Println("Final URL built:", baseUrl.String())

	values := url.Values{}

	// Adding the key-value pairs to the values object
	values.Add("q", "games")
	// values.Add("age", "20")

	// making a new query variable and assigning the values object to it
	encodedQuery := values.Encode()
	fmt.Println(encodedQuery)

	// Build a new url with this new query
	// Custom url to go to reddit and search games
	baseUrl1 := "https://reddit.com/search"
	fullQuery := baseUrl1 + "?" + encodedQuery
	fmt.Println("Full Query:", fullQuery)
}