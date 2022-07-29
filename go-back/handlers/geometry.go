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
	var rs Response
	t.A = rq.A
	t.B = rq.B
	t.C = rq.C
	if rq.Type == "perimeter" {
		rs.Result = t.GetPerimeter()
		c.JSON(http.StatusOK, rs)
		//c.JSON(http.StatusOK, gin.H{"result": p})
	} else if rq.Type == "square" {
		s := t.GetSquare()
		c.JSON(http.StatusOK, gin.H{"result": s})
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
		c.JSON(http.StatusOK, gin.H{"result": p})
	} else if rq.Type == "square" {
		s := f.GetSquare()
		c.JSON(http.StatusOK, gin.H{"result": s})
	} else if rq.Type == "volume" {
		v := f.GetVolume()
		c.JSON(http.StatusOK, gin.H{"result": v})
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
		c.JSON(http.StatusOK, gin.H{"result": p})
	} else if rq.Type == "square" {
		s := r.GetSquare()
		c.JSON(http.StatusOK, gin.H{"result": s})
	} else {
		c.JSON(http.StatusOK, gin.H{"response": "Неправильный тип действия, выберите доступный."})
	}
}

func (ml *MathLogger) CircleProcess(c *gin.Context) {
	var rq Circle
	if err := c.ShouldBindJSON(&rq); err != nil {
		ml.l.Println(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
	}
	var ci Mathematics.Circle
	ci.R = rq.R

	if rq.Type == "longitude" {
		l := ci.GetPerimeter()
		ml.l.Println("Длина окружности найдена")
		c.JSON(http.StatusOK, gin.H{"result": l})
	} else if rq.Type == "square" {
		s := ci.GetSquare()
		ml.l.Println("Площадь найдена")
		c.JSON(http.StatusOK, gin.H{"result": s})
	} else if rq.Type == "volume" {
		v := ci.GetVolume()
		ml.l.Println("Объем найден")
		c.JSON(http.StatusOK, gin.H{"result": v})
	} else {
		ml.l.Println("Неизвестное действие с фигурой")
		c.JSON(http.StatusOK, gin.H{"response": "Invalid process"})
	}
}
