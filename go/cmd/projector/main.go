package main

import (
	"fmt"
	"log"

	"github.com/josiahvehrs/polyglot/pkg/projector"
)

func main() {
	opts, err := projector.GetOpts()
	if err != nil {
		log.Fatalf("unable to get options %v", err)
	}

	fmt.Printf("Opts: %+v", opts)
}
