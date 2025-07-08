package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fname := os.Args[1]
	finfo, err := os.Stat(fname)
	if err != nil {
		fmt.Println("ERROR:")
		fmt.Println(err)
		if errors.Is( err, os.ErrNotExist ) {
			fmt.Printf("ERROR: El archivo [%s] NO EXISTE.\n", fname)
		}
		os.Exit(1)
	}
	fmt.Println( "=== FILE INFORMATION ===" )
	fmt.Println( "name:", finfo.Name() )
	fmt.Println( "size:", finfo.Size() )
	fmt.Println( "Mode:", finfo.Mode() )
	fmt.Println( "IsDir:", finfo.IsDir() )
	fmt.Println( "Sys:", finfo.Sys() )
}

