package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"

	"github.com/hashicorp/mdns"

	"milton/models"
)

type PumpController struct {
	// TODO: Check if we need Host or Name (https://github.com/hashicorp/mdns/blob/master/client.go#L17)
	Host string
	IP   net.IP
}

func QueryActiveUnits(rw http.ResponseWriter, r *http.Request, db models.DB) {
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
	// TODO: Move lookup name to config
	mdns.Lookup("_MILTON._tcp", entriesCh)
	close(entriesCh)

	wg.Wait()
}
