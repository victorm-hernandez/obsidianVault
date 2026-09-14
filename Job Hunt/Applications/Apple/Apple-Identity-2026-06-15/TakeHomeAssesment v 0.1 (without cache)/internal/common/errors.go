package common

import "fmt"

type CompoundError struct {
	Message string
	Err     error
}

func (c *CompoundError) Error() string {
	return fmt.Sprintf("Message: %v, Inner Error: %v", c.Message, c.Err.Error())
}

// Unwrap allows using errors.Is and errors.As to check for the inside error.
func (c *CompoundError) Unwrap() error {
	return c.Err
}
