
# prerequisites

## Install the Protobuf binary and includes for your platform

> see: https://github.com/protocolbuffers/protobuf/releases

### macOS

```bash
$ cp protoc-vers-plat-arch/bin/protoc /usr/local/bin
$ cp -R protoc-vers-plat-arch/include /usr/local/include
```

## Install the Go stuff

### macOS

```bash
$ brew install go
```

## Install Golang support for Protobuf and gPRC

```bash
$ go get -u github.com/golang/protobuf/protoc-gen-go
$ go get -u google.golang.org/grpc
```

# building this example
```bash
$ cd path/to/goprotobufgrpcexample
$ go build ./...     # build recursively (NO OUTPUT - just verifies build works). -n => dry run print of build steps
$ go build -o path/to/output ./path/to/cmd/dir  # build a command/executable.  w/o -o writes to pwd
$ env GOOS=windows GOARCH=amd64 go build ./path/to/cmd/dir  # build a command/executable for a target OS/CPU (here Windows/AMD64)
$ go install ./...   # build and copy results to GOPATH - includes executables and packages
$ go clean -i ./...  # cleans out results of install (-i).  -n => dry run print of files affected
$ git tag -a v0.1.0 -m "tag comment"  # set module version
```
