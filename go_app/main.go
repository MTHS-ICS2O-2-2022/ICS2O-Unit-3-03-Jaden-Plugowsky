// Created by: Jaden Plugowsky
// Created on: April 2023
//
// This program finds the volume of a sphere

package main

import (
	"fmt"
	"math"
)

func main() {
	// This function finds the volume of a sphere
	var radius float64
	var volume float64

	// Input
	fmt.Println("\nThis program finds the volume of a Sphere using given values.")
	fmt.Println()
	fmt.Print("Enter the Radius value (cm): ")
	fmt.Scanln(&radius)

	// Process
	volume = 4 / 3 * math.Pi * math.Pow(radius, 3)

	// Output
	fmt.Println("The volume of the Sphere is: %.2f", volume, "cm³.")

	fmt.Println("\nDone.")
}
