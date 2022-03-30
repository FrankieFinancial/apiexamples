#!/bin/bash

CUSTOMER_ID="your customer id"
API_KEY="your api key"

curl --request POST \
     --url https://api.demo.frankiefinancial.io/compliance/v1.2/entity \
     --header "Accept: application/json" \
     --header "X-Frankie-CustomerID: $CUSTOMER_ID" \
     --header "api_key: $API_KEY" \
     --header "Content-Type: application/json" \
     --data '
{
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
}
'
