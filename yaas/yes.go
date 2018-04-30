package yaas

import "fmt"

const yesConst string = "yes"
const noConst string = "no"

var validPayloads = map[string]bool{
  yesConst: true,
  noConst: true,
}

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

  if (validPayload(fmt.Sprint(yes))) {
    return yesString(fmt.Sprint(y.payload)), nil
  } else {
    return yesString(""), fmt.Errorf("yaas: unable to stringify payload '%q'", y.payload)
  }
}

func newYes() yesStruct {
  return yesStruct{payload: yesString(yesConst)}
}

func newNo() yesStruct {
  return yesStruct{payload: yesString(noConst)}
}

func validPayload(s string) bool {
  return bool(validPayloads[s])
}
