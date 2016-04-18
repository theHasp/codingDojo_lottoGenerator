package generators

import (
	"fmt"
	"math/rand"
	"time"
)

// NewGeneratorWithoutReuse will create a GeneratorWithoutReuse as Generator
func NewGeneratorWithoutReuse() Generator {
	tmp := GeneratorWithoutReuse{}
	tmp.resetUsedNumbers()
	return &tmp
}

// GeneratorWithoutReuse is a Generator which produces Fields were each number occurs
// only once
type GeneratorWithoutReuse struct {
	usableNumbers []int
}

// generate see Generator.generate
// it will return a bunch of fields, were every number only occurs once
// it will return an error if more than 8 fields are requested (because the 9 field)
// can't be filled with uniwue numbers
func (gen *GeneratorWithoutReuse) generate(amountOfFields int) ([]Field, error) {
	var fieldsToReturn []Field
	var emptyReturn []Field

	// reset, so the generator can be used again
	gen.resetUsedNumbers()

	for fieldNumber := 0; fieldNumber < amountOfFields; fieldNumber++ {
		field, err := gen.generateNextField()
		if err != nil {
			return emptyReturn, err
		}
		fieldsToReturn = append(fieldsToReturn, field)
	}

	return fieldsToReturn, nil
}

// generateNextField will generate the next field with unique numbers
func (gen *GeneratorWithoutReuse) generateNextField() (Field, error) {
	var returnField Field

	if len(gen.usableNumbers) < 6 {
		return returnField, fmt.Errorf("Can not generate another Field with unique numbers. Available numbers are %v", gen.usableNumbers)
	}

	var numbers []int
	usableNumbers := gen.usableNumbers
	for count := 0; count < 6; count++ {
		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(len(usableNumbers)) // rnd between 0 and len(usableNumbers)-1)
		numbers = append(numbers, gen.usableNumbers[index])
		usableNumbers = append(usableNumbers[:index], usableNumbers[index+1:]...)
	}

	returnField.Numbers = numbers
	gen.usableNumbers = usableNumbers

	return returnField, nil
}

// resetUsedNumbers will reset the numbers which have been used by the generator
func (gen *GeneratorWithoutReuse) resetUsedNumbers() {
	var usableNumbers []int
	for i := 1; i < 50; i++ {
		usableNumbers = append(usableNumbers, i)
	}
	gen.usableNumbers = usableNumbers
}
