package db

func IterRowsFn(fn func() error) error {
	return fn()
}
