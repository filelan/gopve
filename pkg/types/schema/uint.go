package schema

type UintFunctions struct {
	DefaultValue  *uint
	ValidateFunc  func(uint) bool
	TransformFunc func(uint) uint
}

func (funcs *UintFunctions) Validate(obj uint) bool {
	return funcs == nil || funcs.ValidateFunc == nil || funcs.ValidateFunc(obj)
}

func (funcs *UintFunctions) Transform(obj uint) uint {
	if funcs == nil || funcs.TransformFunc == nil {
		return obj
	}

	return funcs.TransformFunc(obj)
}
