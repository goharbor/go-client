# go-client
Client library with golang for accessing Harbor API.

## Client Types

There are 3 swagger files in this repo.

```sh
api/
    v2.0/
        legacy_swagger.yaml    # legacy client
        swagger.yaml           # v2 client
    swagger.yaml               # assist client contains version and chart healthcheck

```

## Generate Clients

```sh
make gen-harbor-api
```

## To use the Clients

There is `ClientSet` defined inside `pkg/harbor/client.go` to construct Clients

Create a `Config` first then use it to create `ClientSet`

For Example:

```go
c := Config{
    URL: url,
	Transport: tr,
	AuthInfo: ai,
}

cs := NewClientSet(c)

cs.V2() // v2 client
cs.Legacy() // legacy client
cs.Assist() // assist client
```
