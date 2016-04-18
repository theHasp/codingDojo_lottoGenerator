package generators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TEST-METHODS

func TestField_Validate_AmountOfNumbers(t *testing.T) {
	f := Field{}
	err := f.Validate()
	assert.Error(t, err)

	for index := 1; index < 6; index++ {
		f.Numbers = append(f.Numbers, index)
		err = f.Validate()
		assert.Error(t, err)
	}

	f.Numbers = append(f.Numbers, 6)
	err = f.Validate()
	assert.NoError(t, err)

	for index := 7; index < 1000; index++ {
		f.Numbers = append(f.Numbers, 1)
		err = f.Validate()
		assert.Error(t, err)
	}
}

func TestField_Validate_NumbersGt0AndLower50(t *testing.T) {
	// check > 0
	f := Field{}
	for index := 0; index < 6; index++ {
		f.Numbers = append(f.Numbers, index)
	}
	err := f.Validate()
	assert.Error(t, err)

	// check < 50
	f = Field{}
	for index := 45; index < 51; index++ {
		f.Numbers = append(f.Numbers, index)
	}
	err = f.Validate()
	assert.Error(t, err)
}

func TestField_Validate_Duplicates(t *testing.T) {
	f := Field{}
	for index := 1; index < 6; index++ {
		f.Numbers = append(f.Numbers, index)
	}
	f.Numbers = append(f.Numbers, 5)
	err := f.Validate()
	assert.Error(t, err)
}

func TestField_SortNumbers(t *testing.T) {
	f := Field{}
	for index := 6; index > 0; index-- {
		f.Numbers = append(f.Numbers, index)
	}
	f.SortNumbers()
	for index := 1; index < 7; index++ {
		assert.Equal(t, f.Numbers[index-1], index)
	}
}

func TestField_GetString(t *testing.T) {
	f := Field{}
	for index := 6; index > 0; index-- {
		f.Numbers = append(f.Numbers, index)
	}
	txt := f.GetString()
	assert.Equal(t, "[1 2 3 4 5 6]", txt)
}

func TestField_GetMatrix(t *testing.T) {
	f := Field{}
	f.Numbers = append(f.Numbers, 41)
	f.Numbers = append(f.Numbers, 33)
	f.Numbers = append(f.Numbers, 25)
	f.Numbers = append(f.Numbers, 17)
	f.Numbers = append(f.Numbers, 9)
	f.Numbers = append(f.Numbers, 1)

	txt := f.GetMatrix()
	expected := "[ X  2  3  4  5  6  7]\n[ 8  X 10 11 12 13 14]\n[15 16  X 18 19 20 21]\n[22 23 24  X 26 27 28]\n[29 30 31 32  X 34 35]\n[36 37 38 39 40  X 42]\n[43 44 45 46 47 48 49]\n"
	assert.Equal(t, expected, txt)

	f = Field{}
	f.Numbers = append(f.Numbers, 7)
	f.Numbers = append(f.Numbers, 14)
	f.Numbers = append(f.Numbers, 21)
	f.Numbers = append(f.Numbers, 28)
	f.Numbers = append(f.Numbers, 35)
	f.Numbers = append(f.Numbers, 42)

	txt = f.GetMatrix()
	expected = "[ 1  2  3  4  5  6  X]\n[ 8  9 10 11 12 13  X]\n[15 16 17 18 19 20  X]\n[22 23 24 25 26 27  X]\n[29 30 31 32 33 34  X]\n[36 37 38 39 40 41  X]\n[43 44 45 46 47 48 49]\n"
	assert.Equal(t, expected, txt)

	f = Field{}
	f.Numbers = append(f.Numbers, 14)
	f.Numbers = append(f.Numbers, 21)
	f.Numbers = append(f.Numbers, 28)
	f.Numbers = append(f.Numbers, 35)
	f.Numbers = append(f.Numbers, 42)
	f.Numbers = append(f.Numbers, 49)

	txt = f.GetMatrix()
	expected = "[ 1  2  3  4  5  6  7]\n[ 8  9 10 11 12 13  X]\n[15 16 17 18 19 20  X]\n[22 23 24 25 26 27  X]\n[29 30 31 32 33 34  X]\n[36 37 38 39 40 41  X]\n[43 44 45 46 47 48  X]\n"
	assert.Equal(t, expected, txt)

	f = Field{}
	f.Numbers = append(f.Numbers, 1)
	f.Numbers = append(f.Numbers, 8)
	f.Numbers = append(f.Numbers, 15)
	f.Numbers = append(f.Numbers, 22)
	f.Numbers = append(f.Numbers, 29)
	f.Numbers = append(f.Numbers, 36)

	txt = f.GetMatrix()
	expected = "[ X  2  3  4  5  6  7]\n[ X  9 10 11 12 13 14]\n[ X 16 17 18 19 20 21]\n[ X 23 24 25 26 27 28]\n[ X 30 31 32 33 34 35]\n[ X 37 38 39 40 41 42]\n[43 44 45 46 47 48 49]\n"
	assert.Equal(t, expected, txt)

	f = Field{}
	f.Numbers = append(f.Numbers, 8)
	f.Numbers = append(f.Numbers, 15)
	f.Numbers = append(f.Numbers, 22)
	f.Numbers = append(f.Numbers, 29)
	f.Numbers = append(f.Numbers, 36)
	f.Numbers = append(f.Numbers, 43)

	txt = f.GetMatrix()
	expected = "[ 1  2  3  4  5  6  7]\n[ X  9 10 11 12 13 14]\n[ X 16 17 18 19 20 21]\n[ X 23 24 25 26 27 28]\n[ X 30 31 32 33 34 35]\n[ X 37 38 39 40 41 42]\n[ X 44 45 46 47 48 49]\n"
	assert.Equal(t, expected, txt)
}
