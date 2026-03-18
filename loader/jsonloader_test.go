package loader

import (
    "testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
    name := "Gladys"
    want := "Gladyss"
    if want != name  {
        t.Errorf(`fouts %s`, want)
    }
}
