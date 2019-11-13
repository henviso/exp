package main

import "testing"

func TestSum(t *testing.T) {
    want := 3
    if got := Sum(1, 2); got != want {
        t.Errorf("Sum(1, 2) = %q, want %q", got, want)
    }
}

func TestSub(t *testing.T) {
    want := -15
    if got := Sub(1, 16); got != want {
        t.Errorf("Sub(1, 16) = %q, want %q", got, want)
    }
}
