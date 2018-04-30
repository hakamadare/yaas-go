package yaas

import "fmt"

func Inverse(y YesString, err error) (YesString, error) {
  err = fmt.Errorf("unable to invert string '%s'", y)

  switch {
  case yesConst == string(y):
    y = YesString(noConst)
    err = nil
  case noConst == string(y):
    y = YesString(yesConst)
    err = nil
  }

  return y, err
}
