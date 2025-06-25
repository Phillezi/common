package api

type Command interface {
	Run(path string, options ...CommandOption) error
}
