package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var Instance *Server

type Server struct {
	HttpEngine *gin.Engine
}

var logger = logrus.WithField("server", "internal")

// Init 快速初始化
func Init() {
	gin.SetMode(gin.ReleaseMode)
	httpEngine := gin.New()
	httpEngine.Use(ginRequestLog(), gin.Recovery())
	Instance = &Server{
		HttpEngine: httpEngine,
	}
}

// Run 正式开启服务
func Run() {
	logger.Info("http engine starting...")
	err := Instance.HttpEngine.Run("127.0.0.1:9955")
	if err != nil {
		panic(err)
	}
	logger.Info("http engine running...")
}

// StartService 启动服务
// 根据 Module 生命周期 此过程应在Login前调用
// 请勿重复调用
func StartService() {
	logger.Infof("initializing modules ...")
	for _, mi := range modules {
		mi.Instance.Init()
	}
	for _, mi := range modules {
		mi.Instance.PostInit()
	}
	logger.Info("all modules initialized")

	logger.Info("registering modules serve functions ...")
	for _, mi := range modules {
		mi.Instance.Serve(Instance)
	}
	logger.Info("all modules serve functions registered")

	logger.Info("starting modules tasks ...")
	for _, mi := range modules {
		go mi.Instance.Start(Instance)
	}
	logger.Info("tasks running")
}