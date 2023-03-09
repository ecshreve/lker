package server

import (
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

type CloudEntry struct {
	Val   string
	Count int
	Rank  int
}

type ByCount []*CloudEntry

func (a ByCount) Len() int           { return len(a) }
func (a ByCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCount) Less(i, j int) bool { return a[i].Count < a[j].Count }

type Cloud struct {
	Entries []*CloudEntry
	Total   int
}

func (c *Cloud) UpdateRanks() {
	sort.Slice(c.Entries, func(i, j int) bool {
		return c.Entries[i].Count < c.Entries[j].Count
	})

	for ind := range c.Entries {
		if c.Entries[ind].Count == 0 {
			c.Entries[ind].Rank = 1
			continue
		}
		if ind > 19 {
			c.Entries[ind].Rank = 20
			continue
		}
		c.Entries[ind].Rank = ind + 1
	}

	sort.Slice(c.Entries, func(i, j int) bool {
		return c.Entries[i].Val < c.Entries[j].Val
	})
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

	ac := &Cloud{}
	alphabet := strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZ", "")
	for i := range alphabet {
		ac.Entries = append(ac.Entries, &CloudEntry{
			Val:   string(alphabet[i]),
			Count: 0,
			Rank:  1,
		})
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
			s.AlphaCloud.Entries[int(l)-int('A')].Count += 1
			s.AlphaCloud.Total += 1
		}
		s.AlphaCloud.UpdateRanks()
	}

	val := util.GetNearestMs()
	c.HTML(http.StatusOK, "index.html.tpl", gin.H{
		"B": val,
		"G": s.Guesses,
		"W": s.AlphaCloud.Entries,
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
