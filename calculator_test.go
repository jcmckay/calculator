package main

import "testing"

func TestCalculator(t *testing.T) {

	result0 := calculator("")

	if result0 != 0 {
		t.Errorf("Error: result was %d and should have been %d", result0, 0)
	}

	result1 := calculator("1")

	if result1 != 1 {
		t.Errorf("Error: result was %d and should have been %d", result1, 1)
	}

	result2 := calculator("1+1")

	if result2 != 2 {
		t.Errorf("Error: result was %d and should have been %d", result2, 2)
	}

	result3 := calculator("1+9+6")

	if result3 != 16 {
		t.Errorf("Error: result was %d and should have been %d", result3, 16)
	}

	result4 := calculator("8-1")

	if result4 != 7 {
		t.Errorf("Error: result was %d and should have been %d", result4, 7)
	}

	result5 := calculator("9*2")

	if result5 != 18 {
		t.Errorf("Error: result was %d and should have been %d", result5, 18)
	}

	result6 := calculator("4/2")

	if result6 != 2 {
		t.Errorf("Error: result was %d and should have been %d", result6, 2)
	}

	result7 := calculator("9-3/3*2+1")

	if result7 != 5 {
		t.Errorf("Error: result was %d and should have been %d", result7, 5)
	}

	result8 := calculator("100+5")

	if result8 != 105 {
		t.Errorf("Error: result was %d and should have been %d", result8, 105)
	}

	result9 := calculator("100+51")

	if result9 != 151 {
		t.Errorf("Error: result was %d and should have been %d", result9, 151)
	}

	result10 := calculator("12+5")

	if result10 != 17 {
		t.Errorf("Error: result was %d and should have been %d", result10, 17)
	}

	result11 := calculator("12+617")

	if result11 != 629 {
		t.Errorf("Error: result was %d and should have been %d", result11, 629)
	}
}
