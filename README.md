# Noteworthy

* `internal` is a Go tool keyword.  Anything under a directory named `internal` (anywhere in the code tree) is visible only to code within the module
* `cmd` as the directory to put all executables is arbitrary, though it is a common Golang convention
* This repo ignores a semi-common convention of putting all packages under a `pkg` directory
    - I don't like it because it pollutes import statements: `import github/heck/gomoduleexample/pkg/myextpkg`
* `gomoduleexample.go` does not have to be the same name as the repo, but this seems like a common convention
* `gomoduleexample.go`'s line `package gomoduleexample` is, I believe, arbitrary.  I think it can be any name you wish (doesn't have to be the repo name)
    - This naming especially makes sense if `gomoduleexample.go` has a `func init() { ... }` defined (called automatically at startup) that inits module stuff
* `myextpkg/myextpkg.go` does not have to have the same name as the package
    - This naming especially makes sense if `myextpkg.go` has a `func init() { ... }` defined (called automatically at startup) that inits package stuff
* `internal/myintpkg/myintpkg.go` does not have to have the same name as the package
    - This naming especially makes sense if `myintpkg.go` has a `func init() { ... }` defined (called automatically at startup) that inits package stuff
* `myextpkg` and `internal/myintpkg` can have any number of `.go` files
* `myextpkg` and `internal/myintpkg` can have any number of subdirectories (which would then be submodules. e.g.: `github.com/heck/gomoduleexample/myextpkg/subpkg`)

For an exhaustive example of a Go project structure see: https://github.com/golang-standards/project-layout

# build

```bash
$ cd path/to/gomoduleexample
$ go build ./...     # build recursively (NO OUTPUT - just verifies build works). -n => dry run print of build steps
$ go build -o path/to/output ./path/to/cmd/dir  # build a command/executable.  w/o -o writes to pwd
$ env GOOS=windows GOARCH=amd64 go build ./path/to/cmd/dir  # build a command/executable for a target OS/CPU (here Windows/AMD64)
$ go install ./...   # build and copy results to GOPATH - includes executables and packages
$ go clean -i ./...  # cleans out results of install (-i).  -n => dry run print of files affected
$ git tag -a v0.1.0 -m "tag comment"  # set module version
```

# use this module within an (external/3rd party) executable

```bash
$ cat > extcmd.go << EOF
package main

// // The following import fails because myintpkg visible only by module members
// import "github.com/heck/gomoduleexample/internal/myintpkg"
import "github.com/heck/gomoduleexample/myextpkg"

func main() {
    // // myintpkg isn't visible outside of module, so this isn't possible:
    // myintpkg.Run()
    myextpkg.Run()
}
EOF
$ go build ./extcmd.go  # `go build` will install `gomoduleexample` under `$GOPATH/pkg` (if it's not there already)
$ ./extcmd              # prints "Hello, external example!"
```

# how to recreate this example

```bash
# prep the development directory
$ mkdir path/to/gomoduleexample
$ cd path/to/gomoduleexample
$ git init
$ go mod init github.com/heck/gomoduleexample  # creates the `go.mod` file
$ cat > gomoduleexample.go << EOF
package gomoduleexample
EOF
# create the directory tree
$ mkdir cmd internal
$ mkdir cmd/mycmd internal/myintpkg myextpkg
# create the source files
$ cat > myintpkg/myintpkg.go << EOF
package myintpkg

import "fmt"

// Run func
func Run() {
    fmt.Printf("Hello, internal example!\n")
}
EOF
$ cat > myextpkg/myextpkg.go << EOF
package myextpkg

import "fmt"

// Run func
func Run() {
    fmt.Printf("Hello, external example!\n")
}
EOF
$ cat > cmd/mycmd/main.go << EOF
package main

import "github.com/heck/gomoduleexample/internal/myintpkg"
import "github.com/heck/gomoduleexample/myextpkg"

// Main func
func main() {
    myintpkg.Run()
    myextpkg.Run()
}
EOF
```

## build and run the `mycmd` executable

```bash
$ cd path/to/gomoduleexample
$ go build ./cmd/mycmd
$ ./mycmd
Hello, internal example!
Hello, external example!
```

# how to setup a Mac

```bash
$ brew install go
# one time setup of the Go workspace (used for all projects)
$ cat >> ~/.zshrc << EOF
# Set the Go workspace to $HOME/go
export GOPATH="$HOME/go"
export PATH="$PATH:$GOPATH/bin"
EOF
$ mkdir ~/go
$ cd ~/go
$ mkdir src pkg bin  # create dirs used by go commands
```