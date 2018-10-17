package controller

import "net/http"

type Controller struct {
	W http.ResponseWriter
	R *http.Request
}

func (this *Controller) init(w http.ResponseWriter, r *http.Request) {
	this.W = w
	this.R = r
}
