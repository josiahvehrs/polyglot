package main

import (
	"encoding/json"
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

	proj := projector.NewProjector(config)

	if config.Operation == projector.Print {
		if len(config.Args) == 0 {
			data := proj.GetValueAll()
			jsonString, err := json.Marshal(data)
			if err != nil {
				log.Fatalf("failed to unmarshal json")
			}

			fmt.Printf("json == %+v\n", jsonString)
		} else {
			if val, ok := proj.GetValue(config.Args[0]); ok {
				fmt.Printf("%+v\n", val)
			}
		}
	}

	if config.Operation == projector.Add {
		proj.SetValue(config.Args[0], config.Args[1])
		err := proj.Save()
		if err != nil {
			log.Fatalf("errored while saving: %v", err)
		}
	}

	if config.Operation == projector.Remove {
		proj.RemoveValue(config.Args[0])
		err := proj.Save()
		if err != nil {
			log.Fatalf("errored while saving: %v", err)
		}
	}
}
