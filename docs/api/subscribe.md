# subscribe

Subscribe a user to newsletter.

**Endpoint**: *urlPrefix*/subscribe

**Request method:** GET

**Request queries:**
```
email - user email
topics - topics that user wants to subscribe to sepparated by commas (,)
```

**Possible HTTP response codes**: 500, 400, 200

**Response JSON format:**
```json
{
  "error": "", // string, possible values: "", "too many collisions", "no topics chosen", "email is empty", "user already exists", "there is no such topic: %topicname%", "bad email", "database error"
  "token": "" // string, user manage token
}
```

**Example request:**
`http://localhost:8080/newsletter/api/subscribe?email=archhaze24@gmail.com&topics=go,rust`

**Example response:**
```json
{
  "error": "",
  "token": "f0aadecaa262de08e3520cfa573405fc023fb0d73be62b88b3508513e5a59b3ff66b9c74e856044cf1cee87a3b8229e955d7a64b4f37944857a3a82c765317b0"
}
```

## Rate limiting notice:
If your request get rate limited, you will get ***different*** response JSON:
```json
{
  "error": "try again in 0s" // string, time of rate limiting may differ
}
```