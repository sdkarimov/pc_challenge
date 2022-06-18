package storage

type Storage interface {
	Get(string) (interface{}, bool)
	Set(string, interface{})
	RunGC()
}
