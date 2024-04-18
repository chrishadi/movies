# Movies

## Requirements:
- Go 1.22
- PostgreSQL

## How to
Run the `cmd/server/main.go`. The server will listen on port 8080.

Use these payloads for testing:
- POST /directors
```json
{
    "data": {
        "type": "directors",
        "attributes": {
            "name": "Alex Garland",
            "birthdate": 12502800
        }
    }
}
```
- POST /movies
```json
{
    "data": {
        "type": "movies",
        "attributes": {
            "title": "Civil War",
            "genre": "Action",
            "releaseYear": 2024
        },
        "relationships": {
            "director": {
                "data": {
                    "type": "directors",
                    "id": "<director-id>"
                }
            }
        }
    }
}
```

Run unit tests:
```shell
go install github.com/onsi/ginkgo/v2/ginkgo
ginkgo ./...
```
