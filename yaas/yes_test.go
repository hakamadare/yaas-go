package yaas

import "testing"

func TestYes(t *testing.T) {
  yes := NewYes()

  want := yesString("yes")
  got, err := yes.yes()

  if got != want {
    t.Errorf("got %q, want %q, err %q", got, want, err)
  }
}
