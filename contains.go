package contains

import (
	"fmt"
	"log"
	"strconv"
)

// Contains is used to check if the value is in the slice of strings
// If the value is numeric, the slice is converted to an int array
func Contains(s []string, val any) (bool, error) {
	var err error
	switch val := val.(type) {
	case string:
		r, err := containsString(s, val)
		return r, err
	case int:
		r, err := containsInt(s, val)
		return r, err
	case float64:
		r, err := containsFloat64(s, val)
		return r, err
	default:
		err = fmt.Errorf("unsupported type: %v", val)
	}
	return false, err
}

// MustContains is a wrapper for Contains that panics if an error is returned
func MustContains(s []string, val any) bool {
	r, err := Contains(s, val)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return r
}

// containsString checks if a string is in a slice of strings
func containsString(s []string, str string) (bool, error) {
	for _, v := range s {
		if v == str {
			return true, nil
		}
	}
	return false, nil
}

// containsInt checks if an int is in a slice of strings
func containsInt(s []string, i int) (bool, error) {
	sInt := []int{}

	// Convert the string array of ints to int array
	for _, v := range s {
		si, err := strconv.Atoi(v)
		if err != nil {
			return false, fmt.Errorf("error converting string to int: %v", err)
		}
		sInt = append(sInt, si)
	}

	// Check if the int is in the int array
	for _, v := range sInt {
		if v == i {
			return true, nil
		}
	}
	return false, nil
}

// containsFloat64 checks if an float64 value is in a slice of strings
func containsFloat64(s []string, f float64) (bool, error) {
	sFloat := []float64{}

	// Convert the string array of floats to float64 array
	for _, v := range s {
		sf, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return false, fmt.Errorf("error converting string to float64: %v", err)
		}
		sFloat = append(sFloat, sf)
	}

	// Check if the float64 is in the float64 array
	for _, v := range sFloat {
		if v == f {
			return true, nil
		}
	}
	return false, nil
}
