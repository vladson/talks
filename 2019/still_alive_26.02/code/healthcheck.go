package code

import (
	"strings"
	"context"
	"github.com/pkg/errors"
)

type HealthCheck struct {
	Pingers []func(ctx context.Context) (map[string]interface{}, error)
}

//START OMIT
func (hc HealthCheck) Check(ctx context.Context) (map[string]interface{}, error) {
	response := map[string]interface{}{}
	var statusError error
	for _, pinger := range hc.Pingers { // HL
		resp, err := pinger(ctx)
		if err != nil {
			if statusError == nil {
				statusError = err
			} else {
				statusError = errors.Wrap(statusError, err.Error()) // HL
			}
		}
		for k, v := range resp {
			response[k] = v
		}
	}
	response["alive"] = true
	if statusError != nil {
		response["alive"] = false
		response["errors"] = strings.Replace(statusError.Error(), ": ", "; ", len(hc.Pingers))
	}
	return response, statusError // HL
	// END OMIT
}