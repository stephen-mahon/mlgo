package warmUpExercise

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func warmup() {
	a := mat.NewDiagDense(5, []float64{1, 1, 1, 1, 1})
	fmt.Println(a)
}
