# add_admin_key

Adds another key for admin operations.

**Endpoint**: *urlPrefix*/admin/add_admin_key

**Request method:** POST

**Request JSON format:**
```json 
{
  "key": "" // string, existing admin key
}
```

**Possible HTTP response codes**: 400, 403, 500, 207, 200

**Response JSON format:**
```json
{
  "error": "", // string, possible values: "", "bad json", "invalid admin key", "too many collisions", "database error", "warning: admin key added to database, but not saved in file on server"
  "key": "" // string, new admin key
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
  "error": "warning: admin key added to database, but not saved in file on server",
  "key": "391ca44adcc255fbff2947101e3f6dabf1a8bef2d3b681ff18e86beea71661ca968914ee8dbfd5c52e2527dc5134f998bff03ff1475bfaa2b66430ddc6a3b01a"
}
```

## Rate limiting notice:
If your request get rate limited, you will get ***different*** response JSON:
```json
{
  "error": "try again in 0s" // string, time of rate limiting may differ
}
```