# edit_subscribed_topics

Edit topics of subscribed user.

**Endpoint**: *urlPrefix*/edit_subscribed_topics

**Request method:** GET

**Request queries:**

```
topics - topics that user wants to subscribe to sepparated by commas (,)
token - user manage token
```

**Possible HTTP response codes**: 500, 400, 200

**Response JSON format:**

```json
{
  "error": "" // string, possible values: "", "no topics chosen", "there is no such user", "there is no such topic: %topicName%", "database error" 
}
```

**Example request:**
`http://localhost:8080/newsletter/api/edit_subscribed_topics?topics=rust,go&token=c04504032f32b2dc0779737ed8966e638398e920623121ec8a8e9f445ab6c68209260f100c06e90083706fcd04f47ff754ee71278c19fa6cafaf760ac2edf315`

**Example response:**

```json
{
  "error": ""
}
```

## Rate limiting notice:

If your request get rate limited, you will get ***different*** response JSON:

```json
{
  "error": "try again in 0s" // string, time of rate limiting may differ
}
```