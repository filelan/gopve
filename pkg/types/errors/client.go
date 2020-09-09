package errors

type ClientError string

func (err ClientError) Error() string {
	return string(err)
}

func (err ClientError) Is(target error) bool {
	ts := target.Error()
	es := string(err)
	return ts == es
}
