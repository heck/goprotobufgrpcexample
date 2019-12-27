# from a combo of: 
#     https://blog.gopheracademy.com/advent-2017/make/
#     https://tutorialedge.net/golang/go-protocol-buffer-tutorial/
.PHONY: compile
PROTOC_GEN_GO := $(GOPATH)/bin/protoc-gen-go

# If $GOPATH/bin/protoc does not exist, we'll run this command to install it.
$(PROTOC_GEN_GO):
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u google.golang.org/grpc

*.proto: ./api/*.proto | $(PROTOC_GEN_GO)
	protoc --go_out=. ./api/*.proto

# This is a "phony" target - an alias for the above command, so "make compile"
# still works.
compile: *.proto

#
# CLEAN
#
.PHONY: clean
clean:
	-rm ./api/*.pb.go
