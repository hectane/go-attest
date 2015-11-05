package attest

import (
	"fmt"
	"net/http"
)

// Ensure that the request results in the specified status code.
func HttpStatusCode(req *http.Request, statusCode int) error {
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if r.StatusCode != statusCode {
		return fmt.Errorf("%d != %d", r.StatusCode, statusCode)
	}
	return nil
}
