package yaas

import "fmt"

func Inverse(y yesString, err error) (yesString, error) {
  err = fmt.Errorf("unable to invert string '%s'", y)

  switch {
  case yesConst == string(y):
    y = yesString(noConst)
    err = nil
  case noConst == string(y):
    y = yesString(yesConst)
    err = nil
  }

  return y, err
}
