package generators

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TEST-METHODS

func TestGenerator_WithoutReuse_Generate_Setup(t *testing.T) {
	iGen := NewGeneratorWithoutReuse()
	fields, err := iGen.generate(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(fields))
}

func TestGenerator_WithoutReuse_Generate_OneValidField(t *testing.T) {
	iGen := NewGeneratorWithoutReuse()
	fields, err := iGen.generate(1)
	require.NoError(t, err)
	require.Equal(t, 1, len(fields))
	err = fields[0].Validate()
	assert.NoError(t, err)
}

func TestGenerator_WithoutReuse_Generate_MoreThanOneValidFields(t *testing.T) {
	iGen := NewGeneratorWithoutReuse()
	for i := 0; i < 8; i++ {
		fields, err := iGen.generate(i)
		require.NoError(t, err)
		require.Equal(t, i, len(fields))
		for j := 0; j < i; j++ {
			err = fields[j].Validate()
			assert.NoError(t, err)
		}
	}
}

func TestGenerator_WithoutReuse_Generate_MoreThan8FieldsAreNotPossible(t *testing.T) {
	iGen := NewGeneratorWithoutReuse()

	fields, err := iGen.generate(9)
	require.Error(t, err)
	require.Equal(t, 0, len(fields))

	fields, err = iGen.generate(8)
	require.NoError(t, err)
	require.Equal(t, 8, len(fields))
}

func TestGenerator_WithoutReuse_Generate_NoDuplicatesInOneCall(t *testing.T) {
	iGen := NewGeneratorWithoutReuse()

	fields, err := iGen.generate(8)
	require.NoError(t, err)
	require.Equal(t, 8, len(fields))

	checkDuplicatesInFieldsForOneCall(t, fields)
}

func TestGenerator_WithoutReuse_Generate_NonDeterministic(t *testing.T) {
	iGen1 := NewGeneratorWithoutReuse()
	iGen2 := NewGeneratorWithoutReuse()

	for i := 1; i < 9; i++ {
		fields1, err := iGen1.generate(i)
		require.NoError(t, err)

		fields2, err := iGen2.generate(i)
		require.NoError(t, err)

		allFields := append(fields1, fields2...)
		checkNonDeterministic(t, allFields)
	}
}

// HELPER-METHODS

// checkDuplicatesInFieldsForOneCall asserts that every number occurs only once in all fields
func checkDuplicatesInFieldsForOneCall(t *testing.T, fields []Field) {
	allUsedNumbers := make(map[int]int)
	duplicatesExist := false
	for _, field := range fields {
		// check for duplicates within the field
		err := field.Validate()
		require.NoError(t, err)

		for _, number := range field.Numbers {
			// check duplicates in both fields
			count, exists := allUsedNumbers[number]
			duplicatesExist = duplicatesExist || exists
			allUsedNumbers[number] = count + 1
		}
	}
	assert.False(t, duplicatesExist, fmt.Sprintf("Duplicates exist %v", allUsedNumbers))
}

// checkNonDeterministic asserts that all fields are not equal
func checkNonDeterministic(t *testing.T, fields []Field) {
	allUsedCombinations := make(map[string]int)
	nonDeterministic := false
	for _, field := range fields {
		require.NoError(t, field.Validate())
		key := field.GetString()
		count, exists := allUsedCombinations[key]
		nonDeterministic = nonDeterministic || exists
		allUsedCombinations[key] = count + 1
	}
	assert.False(t, nonDeterministic, fmt.Sprintf("Calls are not nondeterministic %v", allUsedCombinations))
}
