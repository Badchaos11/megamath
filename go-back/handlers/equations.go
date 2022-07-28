package handlers

import (
	"github.com/Badchaos11/megamath/go-back/Mathematics"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ml *MathLogger) SolveQuadraticEquation(c *gin.Context) {
	var rq InputQE
	if err := c.ShouldBindJSON(&rq); err != nil {
		ml.l.Println("Ошибка чтения тела запроса")
		c.JSON(http.StatusBadRequest, err.Error())
	}
	var qe Mathematics.QuadraticEquation
	qe.A = rq.A
	qe.B = rq.B
	qe.C = rq.C

	response := qe.SolveQuadraticEquation()
	if response.D < 0 {
		ml.l.Println("Уравнние успешно решено")
		c.JSON(http.StatusOK, gin.H{"discriminant": response.D, "result": "Корней нет"})
	} else {
		ml.l.Println("Уравнение успешно решено")
		c.JSON(http.StatusOK, gin.H{"discriminant": response.D, "X1": response.X1, "X2": response.X2})
	}
}
