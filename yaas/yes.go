package yaas

import "fmt"

const yesConst string = "yes"
const noConst string = "no"

var validPayloads = map[string]bool{
  yesConst: true,
  noConst: true,
}

type YesString string

type yesStruct struct {
  payload YesString
}

type yes interface {
  yes() (YesString, error)
}

func Yes() (YesString, error) {
  y := newYes()
  yes := YesString(fmt.Sprint(y.payload))

  if (validPayload(fmt.Sprint(yes))) {
    return YesString(fmt.Sprint(y.payload)), nil
  } else {
    return YesString(""), fmt.Errorf("yaas: unable to stringify payload '%q'", y.payload)
  }
}

func newYes() yesStruct {
  return yesStruct{payload: YesString(yesConst)}
}

func newNo() yesStruct {
  return yesStruct{payload: YesString(noConst)}
}

func validPayload(s string) bool {
  return bool(validPayloads[s])
}
