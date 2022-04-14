package main

import "testing"

func TestCheckForWin(t *testing.T) {
	t.Run("Test return win", func(t *testing.T) {
		testdata := []int{0, 0, 0, 0, 0}
		want := true
		got := checkForWin(testdata)
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("Test return not won", func(t *testing.T) {
		testdata := []int{0, 1, 0, 0, 0}
		want := false
		got := checkForWin(testdata)
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
}

func TestCheckRightLetterWorngPos(t *testing.T) {
	word := "hello"
	t.Run("Test return right value for pos", func(t *testing.T) {
		guess := "hillo"
		want := []int{0, 2, 0, 0, 0}
		got := checkRightLetterWorngPos(guess, word)
		for i, val := range got {
			if val != want[i] {
				t.Errorf("got %d want %d", val, want[i])
			}
		}

	})
	t.Run("Test return right value for pos", func(t *testing.T) {
		guess := "smile"
		want := []int{2, 2, 2, 0, 1}
		got := checkRightLetterWorngPos(guess, word)
		for i, val := range got {
			if val != want[i] {
				t.Errorf("got %d want %d", val, want[i])
			}
		}

	})
}
