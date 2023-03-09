package server

import (
	"net/http"

	"github.com/ecshreve/lker/pkg/guess"
	"github.com/ecshreve/lker/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
)

type Server struct {
	ID          string
	Router      *gin.Engine
	Guesses     []string
	LetterCloud *guess.Cloud
}

func NewServer() *Server {
	log.Trace("---> - enter")
	defer log.Trace("<--- - exit")

	s := &Server{
		ID:          "SERVER",
		Router:      gin.Default(),
		Guesses:     []string{},
		LetterCloud: guess.NewCloud(),
	}

	s.registerHandlers()

	return s
}

func (s *Server) PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (s *Server) IndexHandler(c *gin.Context) {
	log.Trace("---> - enter")
	defer log.Trace("<--- - exit")

	if guessRaw := c.PostForm("guessbox"); guessRaw != "" {
		log.Debug("guessRaw: ", guessRaw)

		g, err := guess.Sanitize(guessRaw)
		if err != nil {
			log.Error(oops.Wrapf(err, "unable to sanitize guess: %s", guessRaw))
			return
		}

		if err := s.LetterCloud.ProcessGuess(g); err != nil {
			log.Error(oops.Wrapf(err, "unable to process guess for cloud: %v", g))
			return
		}
		log.Debug("sanitized: ", g)
		s.Guesses = append(s.Guesses, g)
	}

	c.HTML(http.StatusOK, "index.html.tpl", gin.H{
		"B": util.GetNearestMs(),
		"G": s.Guesses,
		"W": s.LetterCloud.Items,
	})
}

func (s *Server) registerHandlers() {
	log.Trace("---> - enter")
	defer log.Trace("<--- - exit")

	s.Router.Static("/static", "./static")
	s.Router.StaticFile("favicon.ico", "./static/favicon.ico")

	s.Router.LoadHTMLGlob("pkg/templates/*")
	s.Router.GET("/", s.IndexHandler)
	s.Router.POST("/", s.IndexHandler)
	s.Router.GET("/ping", s.PingHandler)

}

func (s *Server) Serve() error {
	err := s.Router.Run(":8880")
	if err != nil {
		return oops.Wrapf(err, "gin server returned error")
	}
	return nil
}
