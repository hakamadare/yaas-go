package yaas

import "testing"

func TestInverseYes(t *testing.T) {
  want := yesString(noConst)
  got, err := Inverse(Yes())

  if got != want {
    t.Errorf("got %q, want %q, err %q", got, want, err)
  }
}

func TestInverseNo(t *testing.T) {
  want := yesString(yesConst)
  got, err := Inverse(Inverse(Yes()))

  if got != want {
    t.Errorf("got %q, want %q, err %q", got, want, err)
  }
}
