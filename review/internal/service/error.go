package service

import "fmt"

type CouponError struct {
	code string
	idx  int
	err  error
}

func (e CouponError) Error() string {
	return fmt.Sprintf("code: %s, index: %d, error: %v", e.code, e.idx, e.err)
}

func appendError(currentError error, nextError error) error {
	if currentError == nil {
		return nextError
	} else {
		return fmt.Errorf("%w; %v", currentError, nextError)
	}
}
