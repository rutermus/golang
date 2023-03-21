package pkg

const (
  OkMsg = "OK"
  CancelledMsg = "CANCELLED"
)

const (
  OkCode = iota
  CancelledCode
  UnknowCode
)

func ErrorMessageToCode(msg string) int {
  switch msg {
    case OkMsg:
      return OkCode
    case CancelledMsg:
      return CancelledCode
  }
  return UnknowCode
}