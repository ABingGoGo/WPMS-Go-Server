package router

import (
	"SystemUserServer/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func RouterInit() {
	//routerBaseInfoInit(logfile)
	router := gin.New()
	//router.Use(gin.Recovery())
	router.Use(config.LogMiddleware())
	router.GET("/", func(c *gin.Context) {
		logrus.Info("123")
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	router.Run(fmt.Sprintf("%s:%d", config.YamlConfig.Server.Address, config.YamlConfig.Server.Port))
	//routerRunTimeInit(router)
}

//func routerBaseInfoInit(logfile *os.File) {
//	// 禁用控制台颜色
//	gin.DisableConsoleColor()
//	gin.SetMode(gin.ReleaseMode)
//	gin.DefaultWriter = io.MultiWriter(logfile)
//
//}

//func routerRunTimeInit(router *gin.Engine) {
//	srv := &http.Server{
//		Addr:    ":8080",
//		Handler: router,
//	}
//
//	go func() {
//		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
//			LogError.Fatalf("服务启动异常: %s\n", err)
//		}
//		LogInfo.Println("服务启动成功，监听8080")
//	}()
//
//	// 创建一个只接收一个参数的通道用于通知关闭服务器
//	stop := make(chan os.Signal, 1)
//	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
//	// 阻塞等待关闭信号
//	<-stop
//	LogInfo.Println("关闭Http服务 ...")
//
//	// 创建一个5秒的Context用于关闭服务器的等待时间
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	// 优雅关闭服务器
//	if err := srv.Shutdown(ctx); err != nil {
//		LogError.Fatal("Http服务关闭异常:", err)
//	}
//
//	LogInfo.Println("Http服务已经关闭")
//}
