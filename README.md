# BTC Application

This repository contains a Golang server that implements an API for retrieving the current BTC to UAH exchange rate, managing email subscriptions for rate change notifications, and sending rate to subscribed users. App is using <https://currencyapi.com/> for getting BTC/UAH rate.

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
go run ./...
```

Your server will start on port 8080

### Project structure
/cmd/web/main.go - entry point. Here all modules are setups.
/pkg
    /domain - all data models and their communication API
    /http - routes. they have access to services.
    /infrastructure - third part API logic. Mail send logic and get currency logic. Thay must implement interface defined by services that are using them.
    /repository - logic for storing data. It only depends on domain.
    /services - business logic of application. They depends on domain, repository. (sometimes other services)

.env - all enviroment variables. here you can define server port, email user etc...
