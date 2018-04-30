package yaas

import "testing"

func TestYes(t *testing.T) {
  want := YesString("yes")
  got, err := Yes()

  if got != want {
    t.Errorf("got %q, want %q, err %q", got, want, err)
  }
}

func TestYesError(t *testing.T) {
  discard, got := Yes()

  if got != nil {
    t.Errorf("got %q, want %q, discard %q", got, nil, discard)
  }
}
