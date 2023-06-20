# selfletter config

## Config fields:
`rateLimitingTimeoutSeconds` - positive integer, user actions rate limiting timeout in seconds

`adminRateLimitingTimeoutSeconds` - positive integer, admin actions rate limiting timeout in seconds

`DSN` - string, postgresql database config, refer to [wikipedia](https://en.wikipedia.org/wiki/Data_source_name)

`firstRun` - bool, generates admin key and puts it in `admin_keys.txt`. Will be automatically set to false after first
run

`tokenAndKeyLength` - positive integer, should be not 0 and should be divisible by 2, length for admin key and user
tokens

`urlPrefix` - string, api url prefix

`debug` - bool, debug mode

`domain` - string, your internet domain

`internalAddress` - string, internal address of your machine on which you host backend

`email` - object, email specific settings

`from` - string, from email field

`auth` - object, email auth specific settings

`identity` - string, email identity

`username` - string, email server username

`password` - string, email server password

`host` - string, email server address

`port` - string, email server port

`encryption` - string, email server encryption. Possible values: "SSL/TLS", "TLS", "STARTTLS", "SSL", "None".

**Default config**:
```json
{
  "rateLimitingTimeoutSeconds": 10,
  "adminRateLimitingTimeoutSeconds": 1,
  "DSN": "host=localhost port=5432 dbname=selfletter",
  "firstRun": true,
  "tokenAndKeyLength": 128,
  "urlPrefix": "/newsletter/api",
  "debug": false,
  "domain": "localhost:8080",
  "internalAddress": "localhost:8080",
  "email": {
    "from": "John Doe <myemail@myemailserver.com>",
    "auth": {
      "identity": "",
      "username": "myemail@myemailserver.com",
      "password": "mySmtpPassword",
      "host": "smtp.myemailserver.com",
      "port": 465,
      "encryption": "SSL"
    }
  }
}
```

**Config example:**
```json
{
  "rateLimitingTimeoutSeconds": 5,
  "adminRateLimitingTimeoutSeconds": 2,
  "DSN": "host=localhost port=5432 dbname=newsletter-dev",
  "firstRun": true,
  "tokenAndKeyLength": 128,
  "urlPrefix": "/newsletter/api",
  "debug": true,
  "domain": "localhost:8080",
  "internalAddress": "localhost:8080",
  "email": {
    "from": "Kirill Belolipetsky <archhaze24@gmail.com>",
    "auth": {
      "identity": "",
      "username": "archhaze24@gmail.com",
      "password": "mySmtpPassword",
      "host": "smtp.gmail.com",
      "port": 465,
      "encryption": "SSL"
    }
  }
}
```