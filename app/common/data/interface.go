package data

type ParamsValidator interface {
	Validator() error
}
