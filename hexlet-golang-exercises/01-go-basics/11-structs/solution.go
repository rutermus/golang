package pkg

import (
	"strings"
)

type UserCreateRequest struct {
	FirstName string
	Age       int
}

// BEGIN (write your solution here)
func Validate(req UserCreateRequest) string {
  name := req.FirstName
  age := req.Age
  result := ""
  err := "invalid request"

  if name == "" || strings.Contains(name, " ") || age <= 0 || age > 150 {
    return err
  }

  return result
}
// END
