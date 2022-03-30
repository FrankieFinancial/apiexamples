const https = require("https");

const customerID = "your customer id";
const apiKey = "your api key";

const payload = JSON.stringify({
  entity: {
    addresses: [
      {
        country: "AUS",
      },
    ],
    dateOfBirth: {
      dateOfBirth: "1999-11-12",
      country: "AUS",
    },
    gender: "F",
    name: {
      familyName: "TESTTWELVE",
      middleName: "H",
      givenName: "JUDY",
      displayName: "JUDY H TESTTWELVE",
    },
  },
});

const options = {
  hostname: "api.demo.frankiefinancial.io",
  port: 443,
  path: "/compliance/v1.2/entity",
  method: "POST",
  headers: {
    "Content-Type": "application/json",
    Accept: "application/json",
    "X-Frankie-CustomerID": customerID,
    api_key: apiKey,
  },
};

const req = https.request(options, (res) => {
  console.log(`statusCode: ${res.statusCode}`);

  res.on("data", (d) => {
    process.stdout.write(d);
  });
});

req.on("error", (error) => {
  console.error(error);
});

req.write(payload);
req.end();
