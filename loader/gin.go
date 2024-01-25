package loader

import (
	"auth-server/config"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func InitGin() {
	// init the gin instance (engine)
	engine := gin.New()
	engine.Use(gin.Recovery())

	InitRouter(engine)

	port := config.ServerPort
	runParams := config.ServerHost + ":" + strconv.Itoa(port)
	log.Println("master server at ", runParams)

	if err := engine.Run(runParams); err != nil {
		log.Error(err)
		return
	}
}
