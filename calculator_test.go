package main

import "testing"

func TestFullCalc(t *testing.T) {

	run("8", 8.0, t)
	run("(1)", 1.0, t)
	run("(((3)))", 3.0, t)
	run("[((4))]", 4.0, t)
	run("([(6)])", 6.0, t)
	run("8/4", 2.0, t)
	run("100/5", 20.0, t)
	run("100/50", 2.0, t)
	run("(5/2)", 2.5, t)
	run("(100/10)", 10, t)
	run("([(6/2)])", 3.0, t)

	run("7*6", 42.0, t)
	run("40/4*9", 90.0, t)

	run("4-1", 3.0, t)
	run("4+1", 5.0, t)

	// run("[(5+1)/(4-1)]", 2.0, t)

}

func run(exp string, expect float64, t *testing.T) {
	result := calc(exp)

	if result != expect {
		t.Errorf("Error: result was %f and should have been %f", result, expect)
	}
}
