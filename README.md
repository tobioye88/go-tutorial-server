# Go Tutorial Server

## CRUD

## Create

```curl
curl --location --request POST 'http://localhost:8000' \
--header 'Content-Type: application/json' \
--data-raw '{
    "ISBN": "12342",
    "title": "New Book 2",
    "director": {
        "name": "John Doe"
    }
}'
```

## Read

- Read all

  ```curl
  curl --location --request GET 'http://localhost:8000/'
  ```

- Read Single
  ```curl
  curl --location --request GET 'http://localhost:8000/1'
  ```

## Update

```curl
curl --location --request PUT 'http://localhost:8000/81' \
--header 'Content-Type: application/json' \
--data-raw '{
    "ISBN": "12342",
    "title": "New Book 2000",
    "director": {
        "name": "John Doe"
    }
}'

```

## Delete

```curl
curl --location --request DELETE 'http://localhost:8000/81'

```
