package pkg

import (
	"fmt"
)

func DomainForLocale(domain, locale string) string {
  result := ""
  if locale == "" {
    result = fmt.Sprintf("en.%s", domain)
  } else {
    result = fmt.Sprintf("%s.%s", locale, domain)
  }
  return result
}
