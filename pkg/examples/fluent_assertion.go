package examples

type AssertedStruct struct {
	boolField   bool
	stringField string
	intField    int
	sliceField  []string
}

func (s *AssertedStruct) SetBoolField(newValue bool) {
	s.boolField = newValue
}

func (s *AssertedStruct) BoolField() bool {
	return s.boolField
}

func (s *AssertedStruct) SetStringField(newValue string) {
	s.stringField = newValue
}

func (s *AssertedStruct) StringField() string {
	return s.stringField
}

func (s *AssertedStruct) SetIntField(newValue int) {
	s.intField = newValue
}

func (s *AssertedStruct) IntField() int {
	return s.intField
}

func (s *AssertedStruct) SetSliceField(newValue []string) {
	s.sliceField = newValue
}

func (s *AssertedStruct) SliceField() []string {
	return s.sliceField
}
