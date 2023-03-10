package server

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/ecshreve/lker/pkg/guess"
	"github.com/ecshreve/lker/pkg/util"

	"github.com/benbjohnson/clock"
	"github.com/gin-gonic/gin"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	ID          string
	Router      *gin.Engine
	Clock       clock.Clock
	Guesses     []string
	LetterCloud *guess.Cloud
}

func NewServer() *Server {
	s := &Server{
		ID:          "SERVER",
		Router:      gin.Default(),
		Clock:       clock.New(),
		Guesses:     []string{},
		LetterCloud: guess.DefaultCloud(),
	}

	s.registerHandlers()

	return s
}

func (s *Server) PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (s *Server) PostFormHandler(guessRaw string) error {
	g, err := guess.Sanitize(guessRaw)
	if err != nil {
		return oops.Wrapf(err, "unable to sanitize guess: %s", guessRaw)
	}

	if err = s.LetterCloud.ProcessGuess(g); err != nil {
		return oops.Wrapf(err, "unable to process guess for cloud: %v", g)
	}

	if len(s.Guesses) >= 100 {
		choice := rand.Intn(99)
		evict := s.Guesses[choice]
		s.Guesses[choice] = g
		log.Debug(fmt.Sprintf("evicted guess %s at index %d in list of %d guesses", evict, choice, len(s.Guesses)))
		return nil
	}

	s.Guesses = append(s.Guesses, g)
	return nil
}

func (s *Server) IndexHandler(c *gin.Context) {
	if guessRaw := c.PostForm("guessbox"); guessRaw != "" {
		log.Debug("guessRaw: ", guessRaw)
		if err := s.PostFormHandler(guessRaw); err != nil {
			log.Error(oops.Wrapf(err, "processing form"))
		}
	}

	nearestMs, err := util.GetNearestMs(s.Clock, nil)
	if err != nil {
		nearestMs = 86400000
		log.Error(oops.Wrapf(err, "getting nearest ms val"))
	}

	c.HTML(http.StatusOK, "index.html.tpl", gin.H{
		"B": nearestMs,
		"G": s.Guesses,
		"W": s.LetterCloud.Items,
		"D": "",
	})
}

func (s *Server) registerHandlers() {
	s.Router.StaticFile("style.css", "./static/style.css")
	s.Router.StaticFile("favicon.ico", "./static/favicon.ico")

	s.Router.LoadHTMLGlob("pkg/server/templates/*")
	s.Router.GET("/", s.IndexHandler)
	s.Router.POST("/", s.IndexHandler)
	s.Router.GET("/ping", s.PingHandler)

	if err := s.Router.SetTrustedProxies(nil); err != nil {
		log.Error(oops.Wrapf(err, "unable to set proxies"))
	}
}

func (s *Server) Serve() error {
	err := s.Router.Run(":8880")
	if err != nil {
		return oops.Wrapf(err, "gin server returned error")
	}
	return nil
}
