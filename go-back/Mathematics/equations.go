package Mathematics

import "math"

type Equations interface {
	SolveQuadraticEquation()
}

func (obj QuadraticEquation) SolveQuadraticEquation() QuadraticEquation {
	obj.D = obj.B*obj.B - 4*obj.A*obj.C

	if obj.D < 0 {
		return obj
	} else if obj.D == 0 {
		obj.X1 = -1 * obj.B / obj.A
		obj.X2 = obj.X1
		return obj
	} else {
		obj.X1 = (-1*obj.B + math.Sqrt(obj.D)) / obj.A
		obj.X2 = (-1*obj.B - math.Sqrt(obj.D)) / obj.A
		return obj
	}
}
