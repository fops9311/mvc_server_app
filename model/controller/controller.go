package controller

type Action func(params map[string]string) (result string, err error)
