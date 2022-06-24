package controller

type Action func(params map[string]interface{}) (result string, err error)
