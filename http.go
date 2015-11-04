package attest

import (
	"fmt"
	"net/http"
)

// Ensure that a request results in the specified status code.
func HttpStatusCode(req *http.Request, statusCode int) error {
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if r.StatusCode != statusCode {
		fmt.Errorf("received %d; expected %d", r.StatusCode, statusCode)
	}
	return nil
}
