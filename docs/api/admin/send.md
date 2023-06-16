# send

Sends a requested email to all users that are subscribed to requested topic.
All emails will automatically contain unsubscribe link for specific user at the bottom of email.

**Endpoint**: *urlPrefix*/admin/send

**Request method:** POST

**Request JSON format:**

```json 
{
  "key": "", // string, admin key
  "topic": "", // string, topic name
  "message": "", // string, html email content
  "subject": "" // string, email subject
}
```

**Possible HTTP response codes**: 400, 403, 500, 207, 200

**Response JSON format:**

```json
{
  "errors": [], // null or array containing non-critical mail server errors,
  "criticalError": "" // string, possible values: "", "bad json", "invalid admin key", "bad message html", "database error", "mail server auth encryption set incorrectly"
}
```

**Example request:**

```json
{
  "key": "fef259760d03e89b33b5935af72da8585b801a2f86e3a91deee53ee073c7c84051ab43611308c8dfb7ad9ead00a5ff99eb033a6eac248d7d9e96b5fe5fc45d3e",
  "topic": "rust",
  "message": "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <title>seems like aita</title>\n</head>\n<body>\n<p>i just did a super cool thing</p>\n</body>\n</html>",
  "subject": "seems that someone will hate me for this"
}
```

**Example response:**

```json
{
  "errors": null,
  "criticalError": ""
} 
```

## Rate limiting notice:
If your request get rate limited, you will get ***different*** response JSON:
```json
{
  "error": "try again in 0s" // string, time of rate limiting may differ
}
```