/*
 * Package qeq provides functions to solve quadratic equations
 *
 */
package qeq

import (
	"errors"
	"math"
)

// Returns the real solutions for the quadratic equation defined by aX²+bX+c
func Solveqeq(a, b, c float64) (solution1, solution2 float64, err error) {

	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		err = errors.New("Does not have any real solutions")
		return solution1, solution2, err
	}
	d := math.Sqrt(discriminant)
	solution1 = ((-b + d) / (2 * a))
	solution2 = ((-b - d) / (2 * a))
	return solution1, solution2, err
}
