package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
)

func main() {

	r, err := git.PlainClone("/tmp/toto", false, &git.CloneOptions{
		URL:      "https://github.com/go-git/go-git",
		Progress: os.Stdout,
	})
	if err != nil {
		log.Fatal(err)
	}
	w, err := r.Worktree()
	sparseCheckoutDirectories := []string{"config"}
	w.Checkout(&git.CheckoutOptions{
		Branch:                    "refs/heads/master",
		SparseCheckoutDirectories: sparseCheckoutDirectories,
	})
	entries, err := os.ReadDir("/tmp/toto")
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		fmt.Println(e.Name())
	}
	fmt.Println("---")
	entries, err = os.ReadDir("/tmp/toto/config")
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		fmt.Println(e.Name())
	}
}
