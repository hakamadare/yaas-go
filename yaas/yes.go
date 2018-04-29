package yaas

import "fmt"

const yesConst string = "yes"

type yesString string

type yesStruct struct {
  payload yesString
}

type yes interface {
  yes() (yesString, error)
}

func Yes() (yesString, error) {
  y := newYes()
  yes := yesString(fmt.Sprint(y.payload))

  if (fmt.Sprint(yes) == "yes") {
    return yesString(fmt.Sprint(y.payload)), nil
  } else {
    return yesString(""), fmt.Errorf("yaas: unable to stringify payload '%q'", y.payload)
  }
}

func newYes() yesStruct {
  return yesStruct{payload: yesString(yesConst)}
}
