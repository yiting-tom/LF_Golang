package main

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
)

func main() {
	p := Product{
		Name: "coffee",
		Cost: &Money{
			Amount:   100,
			Currency: "NTD",
		},
		l: log.New(),
	}

	p.AddCost(20)

	// just for demo
	m, _ := json.Marshal(p)
	fmt.Printf("P = %+v", string(m))
}
