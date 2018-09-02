package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

func Before(flags *flag.FlagSet) error {
	if flag.NArg() != 2 {
		return fmt.Errorf("incorrect number of args %d != 2", flag.NArg())
	}
	uid := flag.Arg(1)
	if err := PutState(BeforeState{
		Time: time.Now(),
	}, uid); err != nil {
		log.Printf("prompt error: %s", err)
	}
	return nil
}
