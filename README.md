# go-client
Client library with golang for accessing Harbor API.

## Download swagger spec by version

Currently, the default Harbor version is `v2.8.2`.

*NOTE* Default version need to be updated manually inside [Makefile](Makefile)

To download swagger spec:

```sh
make update-spec
```

To download swagger spec from previous Harbor version, add `VERSION` as an argument:

```sh
make update-spec VERSION=v2.8.2
```

## Generate Clients

```sh
make gen-harbor-api VERSION=v2.8.2
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
```
