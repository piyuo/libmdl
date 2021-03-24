# libmdl

Contain model for all service project

## Git

clone source code to local.

``` bash
git clone git@github.com:piyuo/libmdl.git
```

## Development

write test file and using go extension

``` bash
run test | debug test
```

## Test

unit test using go test

``` bash
go test ./... -parallel 16
```

## Update go.mod

To upgrade all dependencies at once for a given module, just run the following from the root directory of your module

This upgrades to the latest or minor patch release

```bash
go get -u ./...
```
