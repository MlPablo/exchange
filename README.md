# BTC Application

This repository contains a Golang server that implements an API for retrieving the current BTC to UAH exchange rate, managing email subscriptions for rate change notifications, and sending rate to subscribed users.

## API Specification

The API follows the Swagger 2.0 specification. You can find the detailed API documentation in the api/openapi/swagger.yaml file.

### Endpoints

- `/rate`: Retrieves the current BTC to UAH exchange rate.
- `/subscribe`: Subscribes an email address to receive rate notifications.
- `/sendEmails`: Sends rate BTC to UAH to all subscribed email addresses.

## Requirements

- Golang
- Docker

## Getting Started

1. Clone the repository:

```shell
git clone https://github.com/MlPablo/exchange.git
cd exchange
```

2. Build the Docker image:
```shell
docker build -t exchange .
```

3. Run the Docker container:
```shell
docker run -p 8080:8080 exchange
```

4. Or run locally:
```shell
go run cmd/web/main.go
```

Your server will start on port 8080


