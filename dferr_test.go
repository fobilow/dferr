package dferr_test

import (
	"fmt"
	"github.com/fobilow/dferr"
	"log"
	"os"
)

func ExampleThis() {
	f, err := os.OpenFile("dferr.go", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	defer dferr.This(f.Close)
	// Output:
}

func ExampleThisWithHandler() {
	f, err := os.OpenFile("dferr.go", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	defer dferr.ThisWithHandler(f.Close, func(err error) {
		log.Println(err)
	})
	// Output:
}