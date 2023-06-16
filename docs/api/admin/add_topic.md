# add_topic

Adds a new topic for users to subscribe.

**Endpoint**: *urlPrefix*/admin/add_topic


**Request method:** POST

**Request JSON format:**
```json 
{
  "key": "", // string, admin key
  "name": "" // string, topic name. can't contain commas (",").
}
```

**Possible HTTP response codes**: 400, 403, 500, 200

**Response JSON format:**
```json
{
  "error": "", // string, possible values: "", "bad json", "invalid admin key", "name can't contain ","", "topic already exists", "database error"
  "topic": "" // string, "" if error happened, topic name if no errors or if topic already exists
}
```

**Example request:**
```json
{
  "key": "fef259760d03e89b33b5935af72da8585b801a2f86e3a91deee53ee073c7c84051ab43611308c8dfb7ad9ead00a5ff99eb033a6eac248d7d9e96b5fe5fc45d3e",
  "name": "go"
}
```

**Example response:**
```json
{
  "error": "",
  "topic": "go"
}
```

## Rate limiting notice:
If your request get rate limited, you will get ***different*** response JSON:
```json
{
  "error": "try again in 0s" // string, time of rate limiting may differ
}
```