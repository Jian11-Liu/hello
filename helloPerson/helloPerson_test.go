package main
import "testing"
func TestHelloPerson(t *testing.T) {
    got := HelloPerson("Chris")
    want := "Hello, Chris"
    if got != want {
        t.Errorf("got '%s' want '%s'", got, want)
    }
}