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

	router.POST("/perimeter", ml.GetPerimeter)
	router.POST("/square", ml.GetSquare)

	if err := router.Run("localhost:9090"); err != nil {
		l.Fatalf("Не удалось запустить сервер", err)
	}
}
