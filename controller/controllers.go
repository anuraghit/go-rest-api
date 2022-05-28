package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Welcome(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// rw.Write("Welcome")
	rw.Write([]byte("Welcome"))
}

func GetFriendsNearBy(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {}