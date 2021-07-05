package cerrors

import "fmt"

type CError struct {
	Code int
	Err  error
}

func (c *CError) Error() string {
	return fmt.Sprintf("Error: %v", c.Err)
}
