package generators

import (
	"fmt"
	"sort"
)

// Generator is the interface for a lotto-number-generator
type Generator interface {
	// generate generates the requested amount of Fields
	// error is returned if the request can not be fullfilled
	generate(amountOfFields int) ([]Field, error)
}

// Field represents one lotto field with 6 numbers
type Field struct {
	Numbers []int
}

// Validate validates if the Numbers in the Field are
// - 6 in total
// - each number is 1 at min
// - each number is 49 at max
// - no duplicated numbers exist
func (f *Field) Validate() error {
	if len(f.Numbers) != 6 {
		return fmt.Errorf("A 6 out of 49 lotto field should have 6 numbers, but this one has %v", len(f.Numbers))
	}
	var alreadyCheckedNumbers []int
	for _, number := range f.Numbers {
		if number < 1 {
			return fmt.Errorf("A 6 out of 49 lotto field should have 6 numbers between 1 and 49, but this one has %v as a number", number)
		}
		if number > 49 {
			return fmt.Errorf("A 6 out of 49 lotto field should have 6 numbers between 1 and 49, but this one has %v as a number", number)
		}
		for _, knownNumber := range alreadyCheckedNumbers {
			if knownNumber == number {
				return fmt.Errorf("A 6 out of 49 lotto field should have 6 unique numbers between 1 and 49, but this one has %v duplicated", number)
			}
		}
		alreadyCheckedNumbers = append(alreadyCheckedNumbers, number)
	}
	return nil
}

// SortNumbers sorts the numbes of the Field in asc order
func (f *Field) SortNumbers() {
	sort.Ints(f.Numbers)
}

// GetString sorts the numbers and produces a readable layout (one line)
func (f *Field) GetString() string {
	f.SortNumbers()
	return fmt.Sprintf("%v", f.Numbers)
}

// GetMatrix sorts the numbers and produces a readable layout (multiple line)
func (f *Field) GetMatrix() string {
	var matrix [][]string
	matrix = append(matrix, []string{" 1", " 2", " 3", " 4", " 5", " 6", " 7"})
	matrix = append(matrix, []string{" 8", " 9", "10", "11", "12", "13", "14"})
	matrix = append(matrix, []string{"15", "16", "17", "18", "19", "20", "21"})
	matrix = append(matrix, []string{"22", "23", "24", "25", "26", "27", "28"})
	matrix = append(matrix, []string{"29", "30", "31", "32", "33", "34", "35"})
	matrix = append(matrix, []string{"36", "37", "38", "39", "40", "41", "42"})
	matrix = append(matrix, []string{"43", "44", "45", "46", "47", "48", "49"})

	f.SortNumbers()
	for _, number := range f.Numbers {
		line := number / 7
		position := number - (line * 7) - 1
		if number%7 == 0 {
			position = 6
			line = line - 1
		}
		matrix[line][position] = " X"
	}

	returnString := ""
	for _, line := range matrix {
		returnString = fmt.Sprintf("%v%v\n", returnString, line)
	}

	return returnString
}
