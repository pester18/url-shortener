package controller

type Context interface {
	JSON(code int, i interface{})
	Redirect(code int, location string)
}
