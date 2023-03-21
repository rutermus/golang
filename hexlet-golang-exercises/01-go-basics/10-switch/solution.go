package pkg

import "strings"

func ModifySpaces(s, mode string) string {
  replacement := ""
  switch mode {
    default:
      replacement = "*"
    case "dash":
      replacement = "-"
    case "underscore":
      replacement = "_"
  }
  return strings.ReplaceAll(s, " ", replacement)
}
