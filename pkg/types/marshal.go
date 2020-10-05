package types

type Marshaler interface {
	Marshal() (string, error)
}

type Unmarshaler interface {
	Unmarshal(s string) error
}
