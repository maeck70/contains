package contains

import (
	"log"
	"strconv"
)

// Contains is used to check if the value is in the slice of strings
// If the value is numeric, the slice is converted to an int array
func Contains(s []string, val any) bool {
	switch val := val.(type) {
	case string:
		return containsString(s, val)
	case int:
		return containsInt(s, val)
	case float64:
		return containsFloat64(s, val)
	default:
		log.Panicf("Unsupported type: %v\n", val)
	}
	return false
}

// containsString checks if a string is in a slice of strings
func containsString(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

// containsInt checks if an int is in a slice of strings
func containsInt(s []string, i int) bool {
	sInt := []int{}

	// Convert the string array of ints to int array
	for _, v := range s {
		si, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("Error converting string to int: %v", err)
		}
		sInt = append(sInt, si)
	}

	// Check if the int is in the int array
	for _, v := range sInt {
		if v == i {
			return true
		}
	}
	return false
}

// containsFloat64 checks if an float64 value is in a slice of strings
func containsFloat64(s []string, f float64) bool {
	sFloat := []float64{}

	// Convert the string array of floats to float64 array
	for _, v := range s {
		sf, err := strconv.ParseFloat(v, 64)
		if err != nil {
			log.Fatalf("Error converting string to float64: %v", err)
		}
		sFloat = append(sFloat, sf)
	}

	// Check if the float64 is in the float64 array
	for _, v := range sFloat {
		if v == f {
			return true
		}
	}
	return false
}
