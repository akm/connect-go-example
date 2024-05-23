# connect-go-example

This is an example of [Connect for Go](https://connectrpc.com/docs/go/getting-started).

## Run local process

```
go run ./cmd/server
```


## Run on local Docker

```
make run
```

## Call API

```
curl \
    --header "Content-Type: application/json" \
    --data '{"name": "Jane"}' \
    http://localhost:8080/greet.v1.GreetService/Greet
```

```
grpcurl \
    -protoset <(buf build -o -) -plaintext \
    -d '{"name": "Jane"}' \
    localhost:8080 greet.v1.GreetService/Greet
```

```
go run ./cmd/client/main.go
```
