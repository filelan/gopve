package schema

type IntFunctions struct {
	ValidateFunc  func(int) bool
	TransformFunc func(int) int
}

func (funcs *IntFunctions) Validate(obj int) bool {
	return funcs == nil || funcs.ValidateFunc == nil || funcs.ValidateFunc(obj)
}

func (funcs *IntFunctions) Transform(obj int) int {
	if funcs == nil || funcs.TransformFunc == nil {
		return obj
	}

	return funcs.TransformFunc(obj)
}
