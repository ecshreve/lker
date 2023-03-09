package server

import (
	"math/rand"
	"net/http"
	"sort"
	"strings"

	"github.com/ecshreve/lker/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/kr/pretty"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
)

type Guess struct {
	GuessVal string `form:"guessbox"`
}

type Entry struct {
	Val   string
	Count int
	Rank  int
}

type Cloud struct {
	Entries map[string]*Entry
	Total   int
}

func (c *Cloud) Mix() []*Entry {
	picked := make(map[string]bool)
	alphabet := strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZ", "")
	for _, a := range alphabet {
		picked[a] = false
	}

	var e []*Entry
	for len(e) < 26 {
		p := rand.Intn(len(alphabet))
		if picked[alphabet[p]] {
			continue
		}
		e = append(e, c.Entries[alphabet[p]])
		picked[alphabet[p]] = true
	}
	return e
}

func (c *Cloud) UpdateRanks() {
	var lst []*Entry
	for _, e := range c.Entries {
		lst = append(lst, e)
	}
	sort.Slice(lst, func(i, j int) bool {
		return lst[i].Count < lst[j].Count
	})
	for k, e := range lst {
		if e.Count == 0 {
			c.Entries[e.Val].Rank = 1
			continue
		}
		if k < 6 {
			c.Entries[e.Val].Rank = 1
			continue
		}
		c.Entries[e.Val].Rank = k - 5
	}
}

type Server struct {
	ID         string
	Router     *gin.Engine
	Guesses    []string
	AlphaCloud *Cloud
}

func NewServer() *Server {
	log.Info("---> NewServer() - enter")
	defer log.Info("<--- NewServer() - exit")

	ac := &Cloud{
		Entries: make(map[string]*Entry),
		Total:   0,
	}
	alphabet := strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZ", "")
	for _, i := range alphabet {
		ac.Entries[i] = &Entry{
			Val:   i,
			Count: 0,
			Rank:  1,
		}
	}

	s := &Server{
		ID:         "SERVER",
		Router:     gin.Default(),
		Guesses:    []string{},
		AlphaCloud: ac,
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
	var g Guess
	if c.ShouldBind(&g) == nil {
		g.GuessVal = strings.ToUpper(g.GuessVal)
		s.Guesses = append(s.Guesses, g.GuessVal)
		for _, l := range g.GuessVal {
			if int(l) < int('A') || int(l) > int('Z') {
				continue
			}
			keyval := string(l)

			s.AlphaCloud.Entries[keyval].Count += 1
			s.AlphaCloud.Total += 1
		}
		s.AlphaCloud.UpdateRanks()
	}

	val := util.GetNearestMs()
	c.HTML(http.StatusOK, "index.html.tpl", gin.H{
		"B": val,
		"G": s.Guesses,
		"W": s.AlphaCloud.Mix(),
	})
	pretty.Print(s.AlphaCloud)
}

func (s *Server) registerHandlers() {
	log.Info("---> registerHandlers() - enter")
	defer log.Info("<--- registerHandlers() - exit")

	s.Router.Static("/static", "./static")
	s.Router.LoadHTMLGlob("pkg/templates/*")
	s.Router.GET("/index", s.IndexHandler)
	s.Router.GET("/ping", s.PingHandler)
}

func (s *Server) Serve() error {
	err := s.Router.Run(":8880")
	if err != nil {
		return oops.Wrapf(err, "gin server returned error")
	}
	return nil
}
