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

	// router := gin.Default()
	// router.Static("/static", "./static")
	// router.LoadHTMLGlob("pkg/templates/*")
	// router.GET("/index", func(c *gin.Context) {
	// 	var g Guess
	// 	c.ShouldBind(&g)
	// 	pretty.Print(g)
	// 	val := util.GetNearestMs()
	// 	c.HTML(http.StatusOK, "index.html.tpl", gin.H{
	// 		"B":        val,
	// 		"guessbox": g.GuessVal,
	// 	})
	// })
	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// router.Run(":8880") // listen and serve on 0.0.0.0:8080

}
