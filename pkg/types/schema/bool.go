package schema

type BoolFunctions struct {
	ValidateFunc  func(bool) bool
	TransformFunc func(bool) bool
}

func (funcs *BoolFunctions) Validate(obj bool) bool {
	return funcs == nil || funcs.ValidateFunc == nil || funcs.ValidateFunc(obj)
}

func (funcs *BoolFunctions) Transform(obj bool) bool {
	if funcs == nil || funcs.TransformFunc == nil {
		return obj
	}

	return funcs.TransformFunc(obj)
}
