package storage

type Storage interface {
	Name() string
	Kind() (string, error)
	Content() (Content, error)
}
