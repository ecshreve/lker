package main

import (
	"github.com/ecshreve/lker/pkg/server"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("---> main() - enter")
	defer log.Info("<--- main() - exit")

	s := server.NewServer()
	if err := s.Serve(); err != nil {
		log.Error("error returned from server", oops.Wrapf(err, "wrapped error from server"))
	}
}
