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

func TestParsedYesYes(t *testing.T) {
  want := YesString("yes")
  got, err := ParsedYes("yes")

  if got != want {
    t.Errorf("got %q, want %q, err %q", got, want, err)
  }
}

func TestParsedYesNo(t *testing.T) {
  want := YesString("no")
  got, err := ParsedYes("no")

  if got != want {
    t.Errorf("got %q, want %q, err %q", got, want, err)
  }
}

func TestParsedYesInvalid(t *testing.T) {
  want := YesString("")
  got, err := ParsedYes("indubitably")

  if got != want {
    t.Errorf("got %q, want %q, err %q", got, want, err)
  }

  if err == nil {
    t.Errorf("%q should throw an error, got %q", got, err)
  }
}

