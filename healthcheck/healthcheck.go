package healthcheck

import (
	"errors"
	"net/http"
	"time"

	health "github.com/docker/go-healthcheck"
)

var updater = health.NewStatusUpdater()

func CheckNothing() error {
	return nil
}

func RegisterChecks() {
	health.Register("manual_http_status", updater)
	health.RegisterPeriodicFunc("check_nothing", 5*time.Second, CheckNothing)

}

// DownHandler registers a manual_http_status that always returns an Error
func DownHandler(w http.ResponseWriter, r *http.Request) {
	updater.Update(errors.New("Manual Check"))
	/*
		if r.Method == "POST" {
			updater.Update(errors.New("Manual Check"))
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	*/
}

// UpHandler registers a manual_http_status that always returns nil
func UpHandler(w http.ResponseWriter, r *http.Request) {
	updater.Update(nil)
	/*
		if r.Method == "POST" {
			updater.Update(nil)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	*/
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	health.StatusHandler(w, r)
	/*
		if r.Method == "GET" {
			checks := health.CheckStatus()
			status := http.StatusOK

			// If there is an error, return 503
			if len(checks) != 0 {
				status = http.StatusServiceUnavailable
			}

			health.statusResponse(w, r, status, checks)
		} else {
			http.NotFound(w, r)
		}
	*/
}

func init() {
	health.DefaultRegistry = health.NewRegistry()
	RegisterChecks()

}
