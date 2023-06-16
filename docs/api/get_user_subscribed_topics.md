# get_user_subscribed_topics

Get a list of topics that user is subscribed to.

**Endpoint**: *urlPrefix*/get_user_subscribed_topics

**Request method:** GET

**Request queries:**
```
token - user manage token
```


**Possible HTTP response codes**: 400, 200, 500

**Response JSON format:**

```json
{
  "error": "", // string, possible values: "", "empty token", "invalid token", "database error"
  "topics": [ // array of user subscribed topic objects or null in case of error
    {
      "name": "" // string, topic name
    }
  ]
}
```

**Example request:**
`http://localhost:8080/newsletter/api/get_user_subscribed_topics?token=9a95efbfbb3edeeb17da01f92185925d73982ddcf4ecaf04d74dda48c346cf7c545532b08e4534ddd132f87c350a2073608090bae9890b76c6fa5d0239254a6e`

**Example response:**
```json
{
  "error": "",
  "topics": [
    {
      "name": "rust"
    }
  ]
}
```