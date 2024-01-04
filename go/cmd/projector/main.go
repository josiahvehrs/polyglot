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

	config, err := projector.NewConfig(opts)
	if err != nil {
		log.Fatalf("unable to get config %v", err)
	}

	fmt.Printf("Opts: %+v", opts)
	fmt.Printf("Opts: %+v", config)
}
