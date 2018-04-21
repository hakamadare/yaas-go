package yaas

import "errors"
import "fmt"

const yesConst string = "yes"

type yesString string

type yesStruct struct {
  payload yesString
}

type yes interface {
  yes() (yesString, error)
}

func (y yesStruct) yes() (yesString, error) {
  return yesString(fmt.Sprint(y.payload)), errors.New("yaas: unable to stringify payload")
}

func NewYes() yesStruct {
  return yesStruct{payload: yesString(yesConst)}
}
