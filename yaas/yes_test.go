package yaas

import "testing"

func TestYes(t *testing.T) {
  want := yesString("yes")
  got, err := Yes()

  if got != want {
    t.Errorf("got %q, want %q, err %q", got, want, err)
  }
}
