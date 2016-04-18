package generators

// NewGeneratorStandard will create a GeneratorStandard as Generator
func NewGeneratorStandard() Generator {
	tmp := GeneratorStandard{}
	intTmp := NewGeneratorWithoutReuse()
	tmp.intGenWithoutReuse = intTmp
	return &tmp
}

// GeneratorStandard is a Generator which produces random Fields
type GeneratorStandard struct {
	intGenWithoutReuse Generator
}

// generate see Generator.generate
// it will return a bunch of random fields
func (gen *GeneratorStandard) generate(amountOfFields int) ([]Field, error) {
	var fields []Field
	var emptyFields []Field
	for i := 0; i < amountOfFields; i++ {
		intFields, err := gen.intGenWithoutReuse.generate(1)
		if err != nil {
			return emptyFields, err
		}
		fields = append(fields, intFields...)
	}
	return fields, nil
}
