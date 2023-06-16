# get_users

Get users according to query.

**Endpoint**: *urlPrefix*/admin/get_users

**Request method:** POST

**Request JSON format:**
```json 
{
  "key": "", // string, admin key
  "query": "" // string, search query
}
```

**Possible HTTP response codes**: 400, 403, 200, 500

**Response JSON format:**
```json
{
  "error": "", // string, possible values: "", "bad json", "invalid admin key", "database error", 
  "users": [ // array of found user objects or null in case of error
    {
      "token": "" // string, user manage token
      "email": "", // string, user email
    }
  ]
}
```

**Example request:**
```json
{
  "key": "fef259760d03e89b33b5935af72da8585b801a2f86e3a91deee53ee073c7c84051ab43611308c8dfb7ad9ead00a5ff99eb033a6eac248d7d9e96b5fe5fc45d3e",
  "query": "haze"
}
```

**Example response:**
```json
{
  "error": "",
  "users": [
    {
      "token": "c906d034f84d7a9385ebec08d6fdfff41a585825f657db2d095cb8d87a7fb335882b1fa3790753ec6c4e5da93f3a05bfbee05e581539d35867d04ed34952def7",
      "email": "archhazespam@gmail.com"
    },
    {
      "token": "485b70f5ccb949ce0e6d0fd16cda15a1955feb6bcde3160a6ca0424771621c287848d68074f5ca4859b00bb42ee9f91f5e91ad44b20404d883c396c172e63c8d",
      "email": "archhaze24spam@gmail.com"
    }
  ]
}
```

## Rate limiting notice:
If your request get rate limited, you will get ***different*** response JSON:
```json
{
  "error": "try again in 0s" // string, time of rate limiting may differ
}
```