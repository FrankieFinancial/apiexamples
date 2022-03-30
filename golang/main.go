package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://api.demo.frankiefinancial.io/compliance/v1.2/entity"
	customerID := "your customer id"
	apiKey := "your api key"

	var jsonStr = []byte(`{
    "entity": {
        "addresses": [
            {
                "country": "AUS"
            }
        ],
        "dateOfBirth": {
            "dateOfBirth": "1999-11-12",
            "country": "AUS"
        },
        "gender": "F",
        "name": {
            "familyName": "TESTTWELVE",
            "middleName": "H",
            "givenName": "JUDY",
            "displayName" : "JUDY H TESTTWELVE"
        }
    }
}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Frankie-CustomerID", customerID)
	req.Header.Set("api_key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
