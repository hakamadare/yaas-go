package yaas

import "testing"

func TestInverseYes(t *testing.T) {
  want := YesString(noConst)
  got, err := Inverse(Yes())

  if got != want {
    t.Errorf("got %q, want %q, err %q", got, want, err)
  }
}

func TestInverseNo(t *testing.T) {
  want := YesString(yesConst)
  got, err := Inverse(Inverse(Yes()))

  if got != want {
    t.Errorf("got %q, want %q, err %q", got, want, err)
  }
}
