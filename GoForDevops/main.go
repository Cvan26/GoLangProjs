package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Response interface { // Create response interface, anything that implemented Getresponse function will be considered as Response interface
	GetResponse() string
}

func (w Words) GetResponse() string { // Get response function of Words type, Words receiver

	return strings.Join(w.Words, ",")
}

func (o Occurrence) GetResponse() string { // Same with Occurrence type
	out := []string{}
	for word, occurrence := range o.Words {
		out = append(out, fmt.Sprintf("%s(%d)", word, occurrence))
	}
	return fmt.Sprintf("%s", strings.Join(out, ",")) // merge a slice of string into just only one string
}

func doRequest(requestURL string) (Response, error) { // HTTP call function for GET request

	if _, err := url.ParseRequestURI(requestURL); err != nil { // check whether or not the URL is valid
		fmt.Printf("fake url %s", err)
		os.Exit(1)
		return nil, fmt.Errorf("Parse Error: %s,", err) // return value of function
	}
	resp, err := http.Get(requestURL) // http.Get function return respone and error
	if err != nil {
		fmt.Printf("can not receive resp, error:%s", err)
		os.Exit(1)
		return nil, fmt.Errorf("HTTP get error :%s", err)
	}
	bs := make([]byte, 9999)     // create variable bs as a slice of byte with max range is 9999
	n, err := resp.Body.Read(bs) // return also the actual length of resp.Body and stored the value in variable "n"
	fmt.Println(string(bs))
	if resp.StatusCode != 200 {
		fmt.Printf("Invalid HTTP output, status code: %d\n", resp.StatusCode)
		os.Exit(1)
		return nil, fmt.Errorf("Read Error: %s", err)
	}

	var page Page
	err = json.Unmarshal(bs[:n], &page) // must specify the slice with actual length, thats why use bs[:n]
	if err != nil {
		fmt.Printf("Unmarshal error: %s", err)
		// os.Exit(1)
		return nil, RequestError{
			HTTPCode: resp.StatusCode,
			Body:     string(bs),
			Err:      fmt.Sprintf("Unmarshall Err: %s", err),
		}
	}
	switch page.Name {
	case "words":
		var word Words
		err = json.Unmarshal(bs[:n], &word)
		if err != nil {
			return nil, fmt.Errorf("Unmarshall Error: %s", err)
		}
		return word, nil
	case "occurrence":
		var occurrence Occurrence
		err = json.Unmarshal(bs[:n], &occurrence)
		// for word, occoccurrence := range occurrence.Words {
		// 	fmt.Printf("%s:%d\n", word, occoccurrence)
		// }
		if err != nil {
			return nil, fmt.Errorf("Unmarshall Error: %s", err)
		}
		return occurrence, nil
	}
	return nil, nil // there's no response
}

type Words struct {
	Input string   `json:"input"`
	Words []string `json:"words"`
}
type Page struct {
	Name string `json:"page"`
}
type Occurrence struct {
	Words map[string]int `json:"words"`
}

func main() {
	args := os.Args
	//JSON to struct golang

	if len(args) < 2 {
		fmt.Println("enter 'go run main.go [arg1]'")
		os.Exit(1)
	}
	res, err := doRequest(args[1]) // call to the doRequest function that was created above with return value is Resposne(interface) and error respectively
	if err != nil {
		if RequestErr, ok := err.(RequestError); ok {
			fmt.Printf("Error: %s (HTTP code: %d, Body: %s)\n", RequestErr.Err, RequestErr.HTTPCode, RequestErr.Body)
			os.Exit(1)
		}
		// fmt.Printf("Error: %s\n", err)
		// os.Exit(1)
	}
	if res == nil {
		fmt.Printf("No Response\n")
		os.Exit(1)
	}
	fmt.Printf("Response :%s", res.GetResponse())
}
