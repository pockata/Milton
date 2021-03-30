package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/hashicorp/mdns"
	"net"
	"net/http"

	"milton/models"
)

type PumpController struct {
	Host string
	IP   net.IP
}

func queryActiveUnits(rw http.ResponseWriter, r *http.Request, db models.DB) {
	// Make a channel for results and start listening
	entriesCh := make(chan *mdns.ServiceEntry, 4)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		// var miltons []PumpController
		miltons := []PumpController{}

		defer wg.Done()

		for entry := range entriesCh {
			miltons = append(miltons, PumpController{
				Host: strings.Split(entry.Host, ".")[0],
				IP:   entry.AddrV4,
			})

			fmt.Printf("Got new entry: %v\n", entry)
		}

		res, err := json.Marshal(miltons)
		if err != nil {
			log.Fatalf("Error converting to json: %v\n", miltons)
		}

		fmt.Fprintf(rw, string(res))
	}()

	// Start the lookup
	mdns.Lookup("_MILTON._tcp", entriesCh)
	close(entriesCh)

	wg.Wait()
}
