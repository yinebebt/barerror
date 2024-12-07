package main

import (
	"errors"
	"fmt"
	"github.com/yinebebt/barerror"
)

func main() {
	// Example 1: Using predefined error types
	err := someFunction()
	if err != nil {
		var appErr *barerror.AppError
		if errors.As(err, &appErr) {
			fmt.Printf("error Code: %d\nmsg: %v\nerr: %v\n\n", appErr.Code, appErr.Message, appErr.Err)
		}
	}

	// Example 2: Creating a new error
	customErr := barerror.ErrValidation.New("username is required")
	fmt.Printf("custom error\nerror Code: %d\nmsg: %v\nerr: %v\n\n", customErr.Code, customErr.Message, customErr.Err)

	// Example 3: Wrapping an error
	wrappedErr := barerror.ErrDBWrite.Wrap(fmt.Errorf("connection timeout"), "Failed to insert user")
	fmt.Printf("wrapped error\nerror Code: %d\nmsg: %v\nerr: %v\n\n", wrappedErr.Code, wrappedErr.Message, wrappedErr.Err)
}

func someFunction() error {
	// Simulating an error condition
	return barerror.ErrNotFound.New("user not found")
}
