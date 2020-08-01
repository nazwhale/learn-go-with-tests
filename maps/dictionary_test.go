package maps

import (
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"cheese": "yellow and tasty"}

	t.Run("Word exists", func(t *testing.T) {
		got, err := dictionary.Search("cheese")
		assertError(t, err, nil)
		assertStrings(t, got, "yellow and tasty")
	})

	t.Run("Word doesn't exist", func(t *testing.T) {
		_, err := dictionary.Search("melon")
		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	word := "apple"
	definition := "green and crunchy"

	t.Run("New word", func(t *testing.T) {
		dictionary := make(Dictionary)
		err := dictionary.Add(word, definition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("Word exists, doesn't overwrite", func(t *testing.T) {
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "red and sticky")
		assertError(t, err, ErrExistingWord)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	word := "apple"
	definition := "green and crunchy"

	t.Run("Existing word", func(t *testing.T) {
		dictionary := Dictionary{word: definition}
		newDefinition := "red and sticky"
		err := dictionary.Update(word, newDefinition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("Nothing to update", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Update(word, definition)
		assertError(t, err, ErrWordDoesntExist)
	})
}

func TestDelete(t *testing.T) {
	word := "apple"
	definition := "green and crunchy"
	dictionary := Dictionary{word: definition}

	dictionary.Delete(word)
	emptyDictionary := Dictionary{}

	if !reflect.DeepEqual(dictionary, emptyDictionary) {
		t.Errorf("got: %q\nwant: %q", dictionary, emptyDictionary)
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	assertError(t, err, nil)
	assertStrings(t, got, definition)
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got: %q\nwant: %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("Got err: %q but expected err: %q", got, want)
	}

	if got == nil && want != nil {
		t.Fatal("Expected to get an error")
	}
}
