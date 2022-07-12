package delivery

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/itsapep/golang-with-mongodb/config"
	"github.com/itsapep/golang-with-mongodb/delivery/controller"
	"github.com/itsapep/golang-with-mongodb/manager"
)

type appServer struct {
	usecaseManager manager.UsecaseManager
	engine         *gin.Engine
	host           string
}

func (a *appServer) initHandlers() {
	controller.NewProductController(a.engine, a.usecaseManager.ProductRegistrationUsecase())
}

func (a *appServer) Run() {
	a.initHandlers()
	err := a.engine.Run(a.host)
	if err != nil {
		return
	}
}

func NewServer() *appServer {
	r := gin.Default()
	c := config.NewConfig()
	infraManager := manager.NewInfraManager(c)
	repoManager := manager.NewRepositoryManager(infraManager)
	usecaseManager := manager.NewUsecaseManager(repoManager)
	host := fmt.Sprintf("%s:%s", c.APIHost, c.APIPort)
	return &appServer{
		usecaseManager: usecaseManager,
		engine:         r,
		host:           host,
	}
}
