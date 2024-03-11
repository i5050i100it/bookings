package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurve(t *testing.T) {
	var myH myHandler

	h := NoSurf(&myH)

	switch v := h.(type) {
	case http.Handler:
		//do nothing
	default:
		t.Errorf(fmt.Sprintf("NoSurf: type is not http.Handler, but is %T", v))
	}

}

func TestSessionLoad(t *testing.T) {
	var myH myHandler

	h := SessionLoad(&myH)

	switch v := h.(type) {
	case http.Handler:
		//do nothing
	default:
		t.Errorf(fmt.Sprintf("SessionLoad: type is not http.Handler, but is %T", v))
	}

}
