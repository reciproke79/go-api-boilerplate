package main

import (
	"fmt"
	"net/http"
)

type Controller struct{}

// Index Doc
// @Summary Get Default
// @Description Get values from default location
// @Tags GetDefault Example
// @Accept json
// @Produce json
// Param Authorization header string true "Authentication header" - disabled
// @Success 200 {string} string "Hi from Index"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Unauthorized"
// @Failure 500 {string} string "ok"
// Security ApiKeyAuth - diabled
// Security OAuth2Implicit[admin, write] - disabled
// @Router / [get]
func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", `{"msg":"Hi from Index"}`)
}
