# CastAI API Client for Go

This client is built using [oapi-codegen](https://github.com/deepmap/oapi-codegen)
against the [CastAI OpenAPI specification](https://api.cast.ai/v1/spec/openapi.json).
The provided `Makefile` will download the latest specification and re-run the
code generator if there are any changes.

```sh
$ go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.1.0
$ make
Checking for newer openapi.json file... found!
Generating castai package...
$ make
Checking for newer openapi.json file... up-to-date
$
```
