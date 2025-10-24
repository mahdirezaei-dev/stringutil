package stringutil

import "testing"

func TestReverse(t *testing.T) {
	got := Reverse("hello")
	want := "olleh"

	if got != want {
		t.Errorf("Reverse(\"hello\") = %q; want %q", got, want)
	}
}
