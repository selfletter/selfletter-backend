# unsubscribe

Unsubscribe user from newsletter.

**Endpoint**: *urlPrefix*/unsubscribe

**Request method:** GET

**Request queries:**
```
token - user manage token
```

**Possible HTTP response codes**: 400, 500, 200

**Response JSON format:**
```json
{
  "error": "" // string, possible values: "", "empty token", "invalid token", "database error"
}
```

**Example request:**
`http://localhost:8080/newsletter/api/unsubscribe?token=f0aadecaa262de08e3520cfa573405fc023fb0d73be62b88b3508513e5a59b3ff66b9c74e856044cf1cee87a3b8229e955d7a64b4f37944857a3a82c765317b0`

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