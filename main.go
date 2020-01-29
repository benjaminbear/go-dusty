package main

import (
	"fmt"
	"os"

	"github.com/schlund/go-dusty/cmd"
)

func main() {
	rootCmd, err := cmd.NewCommandTree()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
