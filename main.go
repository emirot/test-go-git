package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

func main() {

	_, err := git.PlainClone("/workspaces/toto/", false, &git.CloneOptions{
		URL:      "https://github.com/go-git/go-git",
		Progress: os.Stdout,
	})
	fmt.Println("er", err.Error())
}
