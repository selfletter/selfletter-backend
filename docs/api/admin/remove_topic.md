# remove_topic

Remove topic and unsubscribe all users from this topic.

**Endpoint**: *urlPrefix*/admin/remove_topic

**Request method:** POST

**Request JSON format:**
```json 
{
  "key": "", // string, admin key
  "topic": "" // string, topic name
}
```

**Possible HTTP response codes**: 400, 403, 500, 200

**Response JSON format:**
```json
{
  "error": "", // string, possible values: "", "bad json", "invalid admin key", "topic doesn't exist", "database error"
  "topic": "" // string, topic name, empty if error happened
}
```

**Example request:**
```json
{
  "key": "fef259760d03e89b33b5935af72da8585b801a2f86e3a91deee53ee073c7c84051ab43611308c8dfb7ad9ead00a5ff99eb033a6eac248d7d9e96b5fe5fc45d3e",
  "topic": "go"
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