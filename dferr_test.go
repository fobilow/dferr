package dferr_test

import (
	"fmt"
	"github.com/fobilow/dferr"
	"log"
	"os"
)

func ExampleFunc() {
	f, err := os.OpenFile("dferr.go", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	defer dferr.Func(f.Close)
	// Output:
}

func ExampleFuncWithHandler() {
	f, err := os.OpenFile("dferr.go", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	// Below is equivalent to defer log.Println(f.Close())
	defer dferr.FuncWithHandler(f.Close, func(err error) {
		log.Println(err)
	})
	// Output:
}
