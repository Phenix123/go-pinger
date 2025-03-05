package pinger

import (
	"log"
	"net/http"
)

type Pinger struct {
	host string
}

func NewPinger(host string) *Pinger {
	return &Pinger{host: host}
}

func (p *Pinger) Ping() int {
	resp, err := http.Get(p.host)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	return resp.StatusCode
}
