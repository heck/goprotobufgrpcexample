package main

import (
	"github.com/heck/gomoduleexample/internal/myintpkg"
	"github.com/heck/gomoduleexample/myextpkg"
)

// Main func
func main() {
	myintpkg.Run()
	myextpkg.Run()
}
