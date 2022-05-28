package controller

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)
func RequestFollowUser(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {}
func AcceptFollowRequest(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {}
func FollowRequests(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {}
