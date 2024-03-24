package api

import (
	"testing"
)

func TestCalculateCostEfficiency(t *testing.T) {
	wallTypes := []wallType{
		{thickness: 0.24, cost: 100, lambdaValue: 0.21, name: "Concrete Block"},
	}

	insulationTypes := []insulationType{
		{thickness: 0.10, cost: 50, lambdaValue: 0.04, name: "Styrofoam 0.10m"},
		{thickness: 0.20, cost: 100, lambdaValue: 0.04, name: "Styrofoam 0.20m"},
		{thickness: 0.30, cost: 150, lambdaValue: 0.04, name: "Styrofoam 0.30m"},
	}

	desiredUValue := 0.1

	expectedResult := Result{
		wallType:       wallType{thickness: 0.24, cost: 100, lambdaValue: 0.21, name: "Concrete Block"},
		insulationType: insulationType{thickness: 0.10, cost: 50, lambdaValue: 0.04, name: "Styrofoam 0.10m"},
		cost:           150,
		achievedUValue: 0.16666666666666666,
	}

	want := expectedResult
	got, _ := CalculateCostEfficiency(wallTypes, insulationTypes, desiredUValue)

	if want != got {
		t.Errorf("Expected: %v, but got: %v", want, got)
	}
}

func TestCalculateCostEfficiency_NoOptimalSolution(t *testing.T) {
	wallTypes := []wallType{
		{thickness: 0.24, cost: 100, lambdaValue: 0.21, name: "Concrete Block"},
	}

	insulationTypes := []insulationType{
		{thickness: 0.10, cost: 50, lambdaValue: 0.04, name: "Styrofoam 0.10m"},
		{thickness: 0.20, cost: 100, lambdaValue: 0.04, name: "Styrofoam 0.20m"},
		{thickness: 0.30, cost: 150, lambdaValue: 0.04, name: "Styrofoam 0.30m"},
	}

	desiredUValue := 0.1

	want := Result{
		wallType: wallType{
			thickness:   0.24,
			cost:        100,
			lambdaValue: 0.21,
			name:        "Concrete Block",
		},
		insulationType: insulationType{
			thickness:   0.10,
			cost:        50,
			lambdaValue: 0.04,
			name:        "Styrofoam 0.10m",
		},
		cost:           150,
		achievedUValue: 0.16666666666666666,
	}
	got, _ := CalculateCostEfficiency(wallTypes, insulationTypes, desiredUValue)

	if want != got {
		t.Errorf("Expected: %v, but got: %v", want, got)
	}
}
