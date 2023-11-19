package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"net/http"
	time2 "time"
)

func main() {
	e := echo.New()

	log := logrus.New()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			log.WithFields(logrus.Fields{
				"URI":     values.URI,
				"status":  values.Status,
				"error":   values.Error,
				"latency": values.Latency,
				"method":  values.Method,
				"host":    values.Host,
				"ip":      values.RemoteIP,
			}).Info("request")

			return nil
		},
	}))

	e.Use(middleware.CORS())
	e.Use(middleware.CSRF())
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"data": "hello world"})
	})
	e.GET("/time", func(c echo.Context) error {
		time := time2.Now()
		return c.JSON(http.StatusOK, map[string]string{"data": time.String()})
	})
	e.Logger.Fatal(e.Start(":1323"))
}
