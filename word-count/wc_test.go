package wordcount

import (
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	readFile := func(t testing.TB) []byte {
		t.Helper()

		byteData, err := os.ReadFile("../test.txt")
		if err != nil {
			t.Errorf("error reading file")
		}
		return byteData
	}
	t.Run("tests the reading of files", func(t *testing.T) {
		want, err := os.ReadFile("../test.txt")
		if err != nil {
			panic(err)
		}
		got := ReadFile("../test.txt")

		if len(got) != len(want) {
			t.Errorf("Lenghts of same files not equal want %d got %d", len(want), len(got))
		}
	})
	t.Run("tests the number of characters in the file", func(t *testing.T) {
		byteData := readFile(t)

		got := GetCharacterCount(byteData)
		want := 339292
		if int64(want) != got {
			t.Errorf("characters not equal want %d, got %d", want, got)
		}
	})

	t.Run("tests the number of bytes of the file", func(t *testing.T) {
		byteData := readFile(t)

		got := GetFileBytes(byteData)
		want := 342190
		if int64(want) != got {
			t.Errorf("bytes not equal want %d, got %d", want, got)
		}
	})

	t.Run("tests the number of lines of the file", func(t *testing.T) {
		byteData := readFile(t)

		got := GetLineCount(byteData)
		want := 7146
		if int64(want) != got {
			t.Errorf("lines not equal want %d, got %d", want, got)
		}
	})

	t.Run("tests the number of word of the file", func(t *testing.T) {
		byteData := readFile(t)

		got := GetWordCount(byteData)
		want := 58164
		if int64(want) != got {
			t.Errorf("number of words not equal want %d, got %d", want, got)
		}
	})
}
