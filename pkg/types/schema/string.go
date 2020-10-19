package schema

type StringFunctions struct {
	ValidateFunc  func(string) bool
	TransformFunc func(string) string
}

func (funcs *StringFunctions) Validate(obj string) bool {
	return funcs == nil || funcs.ValidateFunc == nil || funcs.ValidateFunc(obj)
}

func (funcs *StringFunctions) Transform(obj string) string {
	if funcs == nil || funcs.TransformFunc == nil {
		return obj
	}

	return funcs.TransformFunc(obj)
}
