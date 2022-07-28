package handlers

import (
	"github.com/Badchaos11/megamath/go-back/Mathematics"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ml *MathLogger) TriangleProcess(c *gin.Context) {
	var rq Triangle
	if err := c.ShouldBindJSON(&rq); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	var t Mathematics.Triangle
	t.A = rq.A
	t.B = rq.B
	t.C = rq.C
	if rq.Type == "perimeter" {
		p := t.GetPerimeter()
		c.JSON(http.StatusOK, p)
	} else if rq.Type == "square" {
		s := t.GetSquare()
		c.JSON(http.StatusOK, s)
	} else {
		c.JSON(http.StatusOK, gin.H{"response": "Неправильный тип действия, выберите доступный."})
	}
}

func (ml *MathLogger) FourProcess(c *gin.Context) {
	var rq Four
	if err := c.ShouldBindJSON(&rq); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	var f Mathematics.FourSquare
	f.A = rq.A

	if rq.Type == "perimeter" {
		p := f.GetPerimeter()
		c.JSON(http.StatusOK, p)
	} else if rq.Type == "square" {
		s := f.GetSquare()
		c.JSON(http.StatusOK, s)
	} else {
		c.JSON(http.StatusOK, gin.H{"response": "Неправильный тип действия, выберите доступный."})
	}
}

func (ml *MathLogger) RectangleProcess(c *gin.Context) {
	var rq Rectangle
	if err := c.ShouldBindJSON(&rq); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	var r Mathematics.Rectangle
	r.A = rq.A
	r.B = rq.B
	if rq.Type == "perimeter" {
		p := r.GetPerimeter()
		c.JSON(http.StatusOK, p)
	} else if rq.Type == "square" {
		s := r.GetSquare()
		c.JSON(http.StatusOK, s)
	} else {
		c.JSON(http.StatusOK, gin.H{"response": "Неправильный тип действия, выберите доступный."})
	}
}
