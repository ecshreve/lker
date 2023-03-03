package server

import (
	"html/template"
	"math/rand"
	"net/http"

	"github.com/ecshreve/lker/pkg/util"
	"github.com/samsarahq/go/oops"

	"golang.org/x/exp/slog"
)

type Server struct {
	ID        string
	Templates map[string]*template.Template
	Handlers  map[string]func()
}

func NewServer() *Server {
	s := &Server{
		ID:        "SERVER",
		Templates: make(map[string]*template.Template),
	}
	s.parseTemplateFiles()
	s.buildHandlers()
	s.registerHandlers()

	return s
}

func (s *Server) buildHandlers() {
	handlers := make(map[string]func())

	indexHandler := func(w http.ResponseWriter, _ *http.Request) {
		slog.Info("---> indexHandler() - enter")
		defer slog.Info("<--- indexHandler() - exit")

		type tplArgs struct {
			B  int
			RF int
		}

		args := tplArgs{
			B:  util.GetNearestMs(),
			RF: rand.Intn(10) + 20,
		}

		slog.Info("tpl args: %s", args)

		if err := s.Templates["index.html.tpl"].ExecuteTemplate(w, "index.html.tpl", args); err != nil {
			slog.Error("", oops.Wrapf(err, "error executing template"))
		}
	}
	handlers["index"] = func() { http.HandleFunc("/", indexHandler) }

	handlers["static"] = func() { http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static")))) }
	s.Handlers = handlers
}

func (s *Server) registerHandlers() {
	for route, handler := range s.Handlers {
		handler()
		slog.Info("registered handler for %s", route)
	}
}

func (s *Server) parseTemplateFiles() {
	tpl := template.Must(template.ParseFiles("/home/eric/github/lker/pkg/templates/index.html.tpl"))
	s.Templates[tpl.Name()] = tpl
}

func (s *Server) Serve() error {
	err := http.ListenAndServe(":8880", nil)
	if err != nil {
		return oops.Wrapf(err, "http server returned error")
	}
	return nil
}
