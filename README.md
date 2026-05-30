# Russ Bassett Orders API

A RESTful API built with Go and Gin for managing project orders.

## Features

- Create Orders
- Retrieve Orders
- Update Orders
- Delete Orders
- List Orders
- Filter Orders by Status
- API Key Authentication
- Input Validation
- In-Memory Data Storage

## Order Model

```json
{
  "id": 1,
  "client_name": "ABC Corporation",
  "project_type": "Control Room",
  "status": "Pending",
  "delivery_date": "2026-06-30"
}
```

Valid Status Values:

- Pending
- In Progress
- Completed

## Authentication

Include the API key in every request:

```http
x-api-key: russ123
```

## Run Application

Install dependencies:

```bash
go mod tidy
```

Run:

```bash
go run main.go
```

Server starts on:

```text
http://localhost:8080
```

## Endpoints

### Create Order

POST /orders

### Get Order

GET /orders/{id}

### Update Order

PUT /orders/{id}

### Delete Order

DELETE /orders/{id}

### List Orders

GET /orders

### Filter Orders

GET /orders?status=Pending

## Technologies

- Go
- Gin Framework
- Git
- GitHub

## Author

Balchander Reddy Savireddy
