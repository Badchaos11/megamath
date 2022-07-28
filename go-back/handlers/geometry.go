package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ml *MathLogger) GetPerimeter(c *gin.Context) {
	figureType := c.Params.ByName("type")
	if figureType == "triangle" {
		p, err := trianglePerimeter(c)
		if err != nil {
			ml.l.Println("Ошибка при получении периметра")
			c.JSON(http.StatusBadRequest, err.Error())
		} else {
			ml.l.Println("Периметр успешно рассчитан")
			c.JSON(http.StatusOK, gin.H{"perimeter": p})
		}
	} else if figureType == "four" {
		p, err := fourPerimeter(c)
		if err != nil {
			ml.l.Println("Ошибка при получении периметра")
			c.JSON(http.StatusBadRequest, err.Error())
		} else {
			ml.l.Println("Периметр успешно рассчитан")
			c.JSON(http.StatusOK, gin.H{"perimeter": p})
		}
	} else if figureType == "rectangle" {
		p, err := rectanglePerimeter(c)
		if err != nil {
			ml.l.Println("Ошибка при получении периметра")
			c.JSON(http.StatusBadRequest, err.Error())
		} else {
			ml.l.Println("Периметр успешно рассчитан")
			c.JSON(http.StatusOK, gin.H{"perimeter": p})
		}
	}
}

func (ml *MathLogger) GetSquare(c *gin.Context) {
	figureType := c.Params.ByName("type")
	if figureType == "triangle" {
		p, err := triangleSquare(c)
		if err != nil {
			ml.l.Println("Ошибка при получении площади")
			c.JSON(http.StatusBadRequest, err.Error())
		} else {
			ml.l.Println("Площадь успешно рассчитан")
			c.JSON(http.StatusOK, gin.H{"perimeter": p})
		}
	} else if figureType == "four" {
		p, err := fourSquare(c)
		if err != nil {
			ml.l.Println("Ошибка при получении площади")
			c.JSON(http.StatusBadRequest, err.Error())
		} else {
			ml.l.Println("Площадь успешно рассчитан")
			c.JSON(http.StatusOK, gin.H{"perimeter": p})
		}
	} else if figureType == "rectangle" {
		p, err := rectangleSquare(c)
		if err != nil {
			ml.l.Println("Ошибка при получении площади")
			c.JSON(http.StatusBadRequest, err.Error())
		} else {
			ml.l.Println("Площадь успешно рассчитан")
			c.JSON(http.StatusOK, gin.H{"perimeter": p})
		}
	}
}
