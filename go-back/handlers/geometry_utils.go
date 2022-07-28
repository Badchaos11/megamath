package handlers

import (
	"fmt"
	"github.com/Badchaos11/megamath/go-back/Mathematics"
	"github.com/gin-gonic/gin"
)

func trianglePerimeter(c *gin.Context) (float64, error) {
	type Request struct {
		A float64 `json:"a,omitempty"`
		B float64 `json:"b,omitempty"`
		C float64 `json:"c,omitempty"`
	}
	var rt Request
	var tp Mathematics.Triangle
	if err := c.ShouldBindJSON(&rt); err != nil {
		fmt.Println(rt)
		fmt.Println(err)
		return 0, err
	}
	fmt.Println(rt)
	tp.A = rt.A
	tp.B = rt.B
	tp.C = rt.C

	return tp.GetPerimeter(), nil

}

func fourPerimeter(c *gin.Context) (float64, error) {
	type Request struct {
		Type string  `json:"type,omitempty"`
		A    float64 `json:"a,omitempty"`
	}

	var f Mathematics.FourSquare
	var r Request

	if err := c.ShouldBindJSON(&r); err != nil {
		return 0, err
	}
	f.A = r.A

	return f.GetPerimeter(), nil
}

func rectanglePerimeter(c *gin.Context) (float64, error) {
	type Request struct {
		Type string  `json:"type,omitempty"`
		A    float64 `json:"a,omitempty"`
		B    float64 `json:"b,omitempty"`
	}
	var r Request
	var rec Mathematics.Rectangle

	if err := c.ShouldBindJSON(&r); err != nil {
		return 0, err
	}
	rec.A = r.A
	rec.B = r.B

	return rec.GetPerimeter(), nil
}

func triangleSquare(c *gin.Context) (float64, error) {
	type Request struct {
		Type string  `json:"type,omitempty"`
		A    float64 `json:"a,omitempty"`
		B    float64 `json:"b,omitempty"`
		C    float64 `json:"c,omitempty"`
	}
	var rt Request
	var tp Mathematics.Triangle
	if err := c.ShouldBindJSON(&rt); err != nil {
		return 0, err
	}
	tp.A = rt.A
	tp.B = rt.B
	tp.C = rt.C

	return tp.GetSquare(), nil

}

func fourSquare(c *gin.Context) (float64, error) {
	type Request struct {
		Type string  `json:"type,omitempty"`
		A    float64 `json:"a,omitempty"`
	}

	var f Mathematics.FourSquare
	var r Request

	if err := c.ShouldBindJSON(&r); err != nil {
		return 0, err
	}
	f.A = r.A

	return f.GetSquare(), nil
}

func rectangleSquare(c *gin.Context) (float64, error) {
	type Request struct {
		Type string  `json:"type,omitempty"`
		A    float64 `json:"a,omitempty"`
		B    float64 `json:"b,omitempty"`
	}
	var r Request
	var rec Mathematics.Rectangle

	if err := c.ShouldBindJSON(&r); err != nil {
		return 0, err
	}
	rec.A = r.A
	rec.B = r.B

	return rec.GetSquare(), nil
}
