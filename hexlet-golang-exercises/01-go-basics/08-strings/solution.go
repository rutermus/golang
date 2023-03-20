package pkg

import (
	"fmt"
	"strings"
)

func Greetings(name string) string {
  name = strings.Trim(name, " ")
  name = strings.ToLower(name)
  name = strings.Title(name)
	
  return fmt.Sprintf("Привет, %s!", name)
}