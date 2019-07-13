package api

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if _, ok := fmt.Fprint(w, "Welcome!\n"); ok != nil {
		// todo
	}
}

func hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if _, ok := fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name")); ok != nil {
		// err handle
	}
}

func getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	if _, ok := fmt.Fprintf(w, "you are get user %s", uid); ok != nil {
		// err handle
	}
}

func modifyUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	if _, ok := fmt.Fprintf(w, "you are modify user %s", uid); ok != nil {

	}
}

func deleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	if _, ok := fmt.Fprintf(w, "you are delete user %s", uid); ok != nil {
		// err handle
	}

}

func addUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// uid := r.FormValue("uid")
	uid := ps.ByName("uid")
	if _, ok := fmt.Fprintf(w, "you are add user %s", uid); ok != nil {
		// err handle
	}
	// fmt.Fprintf(w, "you are add user %s", r)
}