package api

import (
	"errors"
	"math"
)

type wallType struct {
	thickness   float64
	cost        int
	lambdaValue float64
	name        string
}

type insulationType struct {
	thickness   float64
	cost        int
	lambdaValue float64
	name        string
}

type Result struct {
	wallType       wallType
	insulationType insulationType
	cost           int
	achievedUValue float64
}

func CalculateCostEfficiency(wallTypes []wallType, insulationTypes []insulationType, desiredUValue float64) (result Result, err error) {
	// Initialize variables to track the optimal solution
	var optimalWallType string
	var optimalInsulationType string
	var minCost float64 = math.Inf(1)
	var achievedUValue float64 = math.Inf(1)

	// Iterate through the list of wall types
	for _, wt := range wallTypes {
		// Iterate through the list of insulation types
		for _, it := range insulationTypes {
			// Calculate the R value for the current combination of wall and insulation types
			currentRValue := float64(wt.thickness) / it.lambdaValue

			// Calculate the U value for the current combination of wall and insulation types
			currentUValue := 1 / currentRValue

			// Check if the current U value is closer to the desired U value than the previous optimal solution
			if math.Abs(currentUValue-desiredUValue) < math.Abs(achievedUValue-desiredUValue) {
				// Update the optimal solution
				optimalWallType = wt.name
				optimalInsulationType = it.name
				minCost = float64(wt.cost) + float64(it.cost)
				achievedUValue = currentUValue
			}
		}
	}

	// Check if an optimal solution was found
	if math.IsInf(minCost, 1) {
		return result, errors.New("no optimal solution found")
	}

	// Set the result with the optimal solution
	result.wallType = wallType{name: optimalWallType, thickness: wallTypes[0].thickness, cost: wallTypes[0].cost, lambdaValue: wallTypes[0].lambdaValue}
	result.insulationType = insulationType{name: optimalInsulationType, thickness: insulationTypes[0].thickness, cost: insulationTypes[0].cost, lambdaValue: insulationTypes[0].lambdaValue}
	result.cost = int(minCost)
	result.achievedUValue = achievedUValue

	return result, nil
}
