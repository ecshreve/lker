package main

import (
	"golang.org/x/exp/slog"

	"github.com/ecshreve/lker/pkg/server"
	"github.com/samsarahq/go/oops"
)

func main() {
	slog.Info("---> main() - enter")
	defer slog.Info("<--- main() - exit")

	s := server.NewServer()
	if err := s.Serve(); err != nil {
		slog.Error("error returned from server", oops.Wrapf(err, "wrapped error from server"))
	}
}
