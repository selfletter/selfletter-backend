# get_server_config

Get current server config.

**Endpoint**: *urlPrefix*/admin/get_server_config

**Request method:** POST

**Request JSON format:**

```json 
{
  "key": ""
  // string, admin key
}
```

**Possible HTTP response codes**: 400, 403, 200

**Response JSON format:**

```json
{
  "error": "",
  // string, possible values: "", "bad json", "invalid admin key"
  "config": {
    // object, current config, if there is an error config will be with empty (default) fields presented below
    "rateLimitingTimeoutSeconds": 0,
    // positive integer, user actions rate limiting timeout in seconds
    "adminRateLimitingTimeoutSeconds": 0,
    // positive integer, admin actions rate limiting timeout in seconds
    "DSN": "",
    // string, postgresql database config
    "firstRun": false,
    // bool, generates admin key and puts it in `admin_keys.txt`. will be automatically set to false after first run
    "tokenAndKeyLength": 0,
    // positive integer, should be not 0 and should be divisible by 2, length for admin key and user tokens
    "urlPrefix": "",
    // string, api url prefix
    "debug": false,
    // bool, debug mode
    "domain": "",
    // string, server internet domain.
    "internalAddress": "",
    // string, internal address of machine on which backend is hosted.
    "email": {
      // object, email specific settings
      "from": "",
      // string, from email field
      "auth": {
        // object, email auth specific settings
        "identity": "",
        // string, email identity
        "username": "",
        // string, email server username
        "password": "",
        // string, email server password
        "host": "",
        // string, email server address
        "port": 0,
        // string, email server port
        "encryption": ""
        // string, email server encryption. Possible values: "SSL/TLS", "TLS", "STARTTLS", "SSL", "None"
      }
    }
  }
}
```

**Example request:**

```json
{
  "key": "fef259760d03e89b33b5935af72da8585b801a2f86e3a91deee53ee073c7c84051ab43611308c8dfb7ad9ead00a5ff99eb033a6eac248d7d9e96b5fe5fc45d3e"
}
```

**Example response:**

```json
{
  "error": "",
  "config": {
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
}
```

## Rate limiting notice:

If your request get rate limited, you will get ***different*** response JSON:

```json
{
  "error": "try again in 0s"
  // string, time of rate limiting may differ
}
```