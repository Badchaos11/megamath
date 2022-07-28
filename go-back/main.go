package main

import (
	"github.com/Badchaos11/megamath/go-back/handlers"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	l := log.New(os.Stdout, "Math Calc", log.LstdFlags)
	ml := handlers.NewLogger(l)
	router := gin.Default()

	router.POST("/triangle", ml.TriangleProcess)
	router.POST("/four", ml.FourProcess)
	router.POST("/rectangle", ml.RectangleProcess)
	router.POST("/circle", ml.CircleProcess)
	router.POST("/quadratic", ml.SolveQuadraticEquation)

	if err := router.Run("localhost:9090"); err != nil {
		l.Fatalf("Не удалось запустить сервер", err)
	}
}
