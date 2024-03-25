package main

import (
	"fmt"

	"ginblog/internal/middleware"
	"ginblog/pkg/config"
	"ginblog/pkg/http"
	"ginblog/pkg/log"
	"go.uber.org/zap"
)

func main() {
	conf := config.NewConfig()
	logger := log.NewLog(conf)
	jwt := middleware.NewJWT(conf)

	logger.Info("server start", zap.String("host", "http://127.0.0.1:"+conf.GetString("http.port")))

	app, cleanup, err := newApp(conf, logger, jwt)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	http.Run(app, fmt.Sprintf(":%d", conf.GetInt("http.port")))
}
