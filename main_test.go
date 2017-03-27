package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/amitm1/go-microsvc-skel/misc"
)

func TestRequestId(t *testing.T) {

	//recorder := httptest.NewRecorder()

	req := httptest.NewRequest("GET", "http://fakeurl.com/test", nil)

	rid := RequestId(req)

	if !(len(rid) > 10) {
		t.Error("Length not greater than 10")
	}
}

func TestRequestIdWithHeader(t *testing.T) {

	//recorder := httptest.NewRecorder()

	rid := "123456abcde"
	req := httptest.NewRequest("GET", "http://fakeurl.com/test", nil)
	req.Header.Set("X-Request-Id", rid)

	rrid := RequestId(req)

	if !(rrid == rid) {
		t.Error("Didn't match the Request ID set in the header")
	}
}

func TestSetupRequestHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "http://fakeurl.com/test", nil)

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if val, ok := r.Context().Value("RequestHelper").(*misc.RequestHelpers); !ok {
			t.Errorf("Request helper not set in context %v", val)
		}

	})

	rr := httptest.NewRecorder()

	handler := SetupRequestHandler(testHandler, "test")
	handler.ServeHTTP(rr, req)
}

func TestGetCache(t *testing.T) {
	t.Error("Error")
}
