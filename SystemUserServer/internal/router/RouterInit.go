package router

import (
	"SystemUserServer/config"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
)

func GinRouterInit() {
	routerBaseInfoInit()
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.LoggerWithFormatter(config.LoggerWithFormatter))
	// 服务启动检测rul
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, config.YamlConfig.Server.ApplicationName)
	})

	routerV1 := router.Group(config.YamlConfig.Server.Servlet.ContextPath + "/api/v1")
	systemUserQueryRouter(routerV1)

	routerRunTimeInit(router)
}

func routerBaseInfoInit() {
	// 禁用控制台颜色
	gin.DisableConsoleColor()
	// 日志配置
	gin.SetMode(gin.ReleaseMode)
}

func routerRunTimeInit(router *gin.Engine) {
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.YamlConfig.Server.Address, config.YamlConfig.Server.Port),
		Handler: router,
	}
	// 检测服务是否启动
	go func() {
		for {
			time.Sleep(time.Second)
			resp, err := http.Get(fmt.Sprintf("http://%s:%d", config.YamlConfig.Server.Address, config.YamlConfig.Server.Port))
			if err != nil {
				logrus.Println("服务启动异常:", err)
				continue
			}
			if resp.StatusCode != http.StatusOK {
				logrus.Println("服务自检异常，访问自检地址返回状态码:", resp.StatusCode)
				continue
			}
			body, _ := io.ReadAll(resp.Body)
			isCheck := strings.Replace(string(body), "\"", "", -1) != config.YamlConfig.Server.ApplicationName
			if isCheck {
				logrus.Errorf("服务自检异常，请检查端口 %d是否被 %s服务占用", config.YamlConfig.Server.Port, config.YamlConfig.Server.ApplicationName)
			}
			break
		}
		logrus.Infof(fmt.Sprintf("服务启动成功，监听%s:%d", config.YamlConfig.Server.Address, config.YamlConfig.Server.Port))
	}()

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logrus.Errorf("服务启动异常: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logrus.Info("开始关闭服务...")

	//等待已连接请求处理，超时时间15秒
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logrus.Error("关闭服务异常:", err)
	}
	logrus.Info("服务已关闭，3秒后退出程序")
	time.Sleep(3 * time.Second)
}
