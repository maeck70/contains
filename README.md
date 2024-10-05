## contains

**A Go package that checks if a value exists in a slice of strings, handling numeric types.**

This package provides a versatile `Contains` function that efficiently determines if a value is present within a slice of strings. It gracefully handles potential numeric values by converting the slice to the corresponding numeric type (int or float64) for comparison.

### Installation

```bash
go get -u github.com/your-username/contains
```

Replace `your-username` with your actual GitHub username or the location of your package.

### Usage

Import the `contains` package and leverage the `Contains` function:

```go
package main

import (
    "fmt"
    "github.com/your-username/contains"
)

func main() {
    stringSlice := []string{"apple", "banana", "10", "3.14"}

    fmt.Println(contains.Contains(stringSlice, "apple"))        // true
    fmt.Println(contains.Contains(stringSlice, "orange"))      // false
    fmt.Println(contains.Contains(stringSlice, 10))             // true (converts "10" to int)
    fmt.Println(contains.Contains(stringSlice, 3.14159))        // false (converts to float64 with slight precision loss)
}
```

**Important Note:** Precision loss might occur when converting strings representing floating-point numbers due to the inherent limitations of data types.

### Functions

* **`Contains(s []string, val any) bool`**
  - This is the primary function, taking a slice of strings (`s`) and a value (`val`) of any type.
  - It switches on the type of `val` to perform type-specific checks:
    - String: Calls `containsString` for direct comparison.
    - Int: Converts the string slice to an int slice and calls `containsInt`.
    - Float64: Converts the string slice to a float64 slice and calls `containsFloat64`.
    - Any other type: Panics with a log message indicating unsupported type.
* **`containsString(s []string, str string) bool`**
  - Performs a simple string comparison within the string slice.
* **`containsInt(s []string, i int) bool`**
  - Converts the string slice to an int slice and checks if the provided int (`i`) exists within it.
* **`containsFloat64(s []string, f float64) bool`**
  - Converts the string slice to a float64 slice and checks if the provided float64 (`f`) exists within it.

**Error Handling:**
- The `containsInt` and `containsFloat64` functions utilize `strconv.Atoi` and `strconv.ParseFloat` respectively. These functions return an error if the conversion fails. In such cases, the code logs a fatal error indicating the conversion failure. 

This approach ensures that any potential errors during conversion are gracefully handled and reported.

Feel free to customize the package name, import path, and error handling behavior as needed for your specific project.