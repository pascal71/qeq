package qeq

import (
	"testing"
)

func TestSolvqeq(t *testing.T) {

	var has_real_solutions bool

	tables := []struct {
		a  float64
		b  float64
		c  float64
		s1 float64
		s2 float64
		rs bool
	}{
		{1, 1, 0, 0, -1, true},
		{1, 1, 1, 0, 0, false},
		{4, 8, 3, -0.5, -1.5, true},
		{3, 12, 0, 0, -4, true},
		{3, 12, 12, -2, -2, true},
		{5, 1, 3, 0, 0, false},
	}

	for _, table := range tables {

		s1, s2, err := Solveqeq(table.a, table.b, table.c)
		has_real_solutions = (err == nil)

		if has_real_solutions != table.rs {
			t.Errorf("Solvqeq(%f,%f,%f)==(%f,%f,%v) should not have any real solutions.", table.a, table.b, table.c, s1, s2, err)
		} else {
			if s1 != table.s1 || s2 != table.s2 {
				t.Errorf("Solvqeq(%f,%f,%f)==(%f,%f,%v) wanted (%f,%f).", table.a, table.b, table.c, s1, s2, err, table.s1, table.s2)
			}
		}

	}
}
