# README.md

## Running the Application

To start the application, use the following command:

```sh
go run main.go
```

## Running Tests

To run tests, use one of the following commands:

```sh
go test -v
```

or

```sh
make unit-test
```

## Running gRPC Services

To call gRPC services, follow these steps:

1. Navigate to the `/pie_fire_dire_grpc_server` folder and run:

    ```sh
    go run main.go
    ```

2. Navigate to the `/client_call_grpc` folder and run:

    ```sh
    go run main.go
    ```

