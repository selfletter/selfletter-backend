# get_topics

Get an array of all topics that are available for user to subscribe.

**Endpoint**: *urlPrefix*/get_topics

**Request method:** GET

**Request queries:** None

**Possible HTTP response codes**: 500, 200

**Response JSON format:**
```json
{
  "error": "", // string, possible values: "", "database error"
  "topics": [ // array of topic objects or null in case or error
    {
      "name": "" // topic name
    }
  ]
}
```

**Example request:**
`http://localhost:8080/newsletter/api/get_topics`

**Example response:**
```json
{
  "error": "",
  "topics": [
    {
      "name": "rust"
    },
    {
      "name": "go"
    }
  ]
}
```