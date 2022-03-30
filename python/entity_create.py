import requests
import json

url = "https://api.demo.frankiefinancial.io/compliance/v1.2/entity"
customerID = "your customer id"
apiKey = "your api key"


headers = {
    "Content-Type": "application/json; charset=utf-8",
    "Accept": "application/json",
    "X-Frankie-CustomerID": customerID,
    "api_key": apiKey
}

data = {
    "entity": {
        "addresses": [
            {
                "country": "AUS",
            },
        ],
        "dateOfBirth": {
            "dateOfBirth": "1999-11-12",
            "country": "AUS",
        },
        "gender": "F",
        "name": {
            "familyName": "TESTTWELVE",
            "middleName": "H",
            "givenName": "JUDY",
            "displayName": "JUDY H TESTTWELVE",
        },
    },
}

response = requests.post(url, headers=headers, json=data)

print("Status Code", response.status_code)
print("JSON Response ", response.json())
