## CLI

1. Start server

- Bind to `127.0.0.1:8080` and serve `HTTP`
- Allow interacting with secrets engine through REST API

```bash
secretstore server
```

2. Store/Update secret
```bash
secretstore put --key="key-name" --value="contents" --encoding-key="passphrase"
```

3. Get secret
```bash
secretstore get --key="key-name" --encoding-key="passphrase"
```

4. List secrets
```bash
secretstore list-secret-keys
```

## Library
```go
// Use a library-provided KV store
// or create your own that implements the secretstore.KVStore interface
volatileKVStore := secretstore.NewInMemoryKVStore()
store, err := secretstore.NewSecretStore(volatileKVStore)

err, ok := store.Put("secret-key", "secret-value", "encoding-key")
value, err := store.Get("secret-key", "encoding-key")
secrets, err := store.List()
```

## REST API
1. Store/Update secret

`POST` `/secret/{secret_name}`

Request to this endpoint must include the `X-SECRETSTORE-ENCODING-KEY` header.

2. Get secret

`GET` `/secret/{secret_name}`

Request to this endpoint must include the `X-SECRETSTORE-ENCODING-KEY` header.

3. List all secrets

`GET` `/secret`
