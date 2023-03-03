package server

import (
	"html/template"
	"math/rand"
	"net/http"

	"github.com/ecshreve/lker/pkg/util"
	"github.com/samsarahq/go/oops"

	log "github.com/sirupsen/logrus"
)

type Server struct {
	ID        string
	Templates map[string]*template.Template
	Handlers  map[string]func()
}

func NewServer() *Server {
	log.Info("---> NewServer() - enter")
	defer log.Info("<--- NewServer() - exit")

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
	log.Info("---> buildHandlers() - enter")
	defer log.Info("<--- buildHandlers() - exit")

	handlers := make(map[string]func())

	indexHandler := func(w http.ResponseWriter, _ *http.Request) {
		log.Info("---> indexHandler() - enter")
		defer log.Info("<--- indexHandler() - exit")

		type tplArgs struct {
			B  int64
			RF int
		}

		args := tplArgs{
			B:  util.GetNearestMs(),
			RF: rand.Intn(10),
		}

		log.Infof("tpl args: %v", args)

		if err := s.Templates["index.html.tpl"].ExecuteTemplate(w, "index.html.tpl", args); err != nil {
			log.Error("", oops.Wrapf(err, "error executing template"))
		}
	}
	handlers["index"] = func() { http.HandleFunc("/", indexHandler) }
	handlers["static"] = func() { http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static")))) }
	s.Handlers = handlers
}

func (s *Server) registerHandlers() {
	log.Info("---> registerHandlers() - enter")
	defer log.Info("<--- registerHandlers() - exit")

	for route, handler := range s.Handlers {
		handler()
		log.Infof("registered handler for %s", route)
	}
}

func (s *Server) parseTemplateFiles() {
	log.Info("---> parseTemplateFiles() - enter")
	defer log.Info("<--- parseTemplateFiles() - exit")

	tpl := template.Must(template.ParseFiles("pkg/templates/index.html.tpl"))
	s.Templates[tpl.Name()] = tpl
}

func (s *Server) Serve() error {
	err := http.ListenAndServe(":8880", nil)
	if err != nil {
		return oops.Wrapf(err, "http server returned error")
	}
	return nil
}
