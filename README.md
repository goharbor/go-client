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

## Download swagger spec by version

Currently, the default Harbor version is `v2.5.0`.

*NOTE* Default version need to be updated manually inside [Makefile](Makefile)

To download swagger spec:

```sh
make update-spec
```

To download swagger spec from previous Harbor version, add `VERSION` as an argument:

```sh
make update-spec VERSION=v2.5.0
```

## Generate Clients

```sh
make gen-harbor-api VERSION=v2.5.0
```

## Check all available make commands

```sh
make
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
