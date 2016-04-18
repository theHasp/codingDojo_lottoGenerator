package generators

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TEST-METHODS

func TestGenerator_Standard_Generate_Setup(t *testing.T) {
	iGen := NewGeneratorStandard()
	fields, err := iGen.generate(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(fields))
}

func TestGenerator_Standard_Generate_OneValidField(t *testing.T) {
	iGen := NewGeneratorStandard()
	fields, err := iGen.generate(1)
	require.NoError(t, err)
	require.Equal(t, 1, len(fields))
	err = fields[0].Validate()
	assert.NoError(t, err)
}

func TestGenerator_Standard_Generate_MoreThanOneValidFields(t *testing.T) {
	iGen := NewGeneratorStandard()
	for i := 0; i < 10; i++ {
		fields, err := iGen.generate(i)
		require.NoError(t, err)
		require.Equal(t, i, len(fields))
		for j := 0; j < i; j++ {
			err = fields[j].Validate()
			assert.NoError(t, err)
		}
	}
}

func TestGenerator_Standard_Generate_InternalGeneratorFailure(t *testing.T) {
	iGen := NewGeneratorStandard()
	failGen := GeneratorFailiure{}
	stanGen := iGen.(*GeneratorStandard)
	stanGen.intGenWithoutReuse = &failGen
	fields, err := iGen.generate(1)
	assert.Error(t, err)
	assert.Equal(t, 0, len(fields))
}

func TestGenerator_Standard_Generate_MoreThan8FieldsArePossible(t *testing.T) {
	iGen := NewGeneratorStandard()

	fields, err := iGen.generate(9)
	require.NoError(t, err)
	require.Equal(t, 9, len(fields))
}

func TestGenerator_Standard_Generate_NonDeterministic(t *testing.T) {
	iGen1 := NewGeneratorStandard()
	iGen2 := NewGeneratorStandard()

	for i := 1; i < 9; i++ {
		fields1, err := iGen1.generate(i)
		require.NoError(t, err)

		fields2, err := iGen2.generate(i)
		require.NoError(t, err)

		allFields := append(fields1, fields2...)
		checkNonDeterministic(t, allFields)
	}
}

// HELPER-STRUCTS
// GeneratorFailiure is a Generator which throws an error on each generate call
type GeneratorFailiure struct {
}

// generate see Generator.generate
// will fail on every call
func (gen *GeneratorFailiure) generate(amountOfFields int) ([]Field, error) {
	var fields []Field
	return fields, errors.New("This call should fail always")
}
