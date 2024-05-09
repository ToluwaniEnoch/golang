package main

import "testing"

var justATestString = "this is just a test"


func TestDelete(t *testing.T) {
	t.Run("delete existing word", func(t *testing.T) {
		word := "test"
		definition := justATestString

		dictionary := Dictionary{word: definition}

		err := dictionary.Delete(word)

		assertError(t, err, nil)
	})
	
	t.Run("delete non existing word", func(t *testing.T) {
		word := "test"

		dictionary := Dictionary{}

		err := dictionary.Delete(word)

		assertError(t, err, ErrCouldNotFindWord)
	})
}


func TestUpdate(t *testing.T) {
	t.Run("update existing word", func(t *testing.T) {
		word := "test"
		definition := justATestString

		dictionary := Dictionary{word: definition}

		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertDefinition(t, dictionary, word, newDefinition)
		assertError(t, err, nil)
	})
	
	t.Run("update new word", func(t *testing.T) {
		word := "test"
		definition := justATestString

		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		// assertDefinition(t, dictionary, word, ErrCouldNotFindWord.Error())
		assertError(t, err, ErrCouldNotFindWord)
	})
}


func TestAdd(t *testing.T) {

	t.Run("new word", func (t *testing.T)  {
		dictionary := Dictionary{}

		word := "test"
		definition := justATestString

		err := dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)
		assertError(t, err, nil)
	})

	t.Run("existing word", func (t *testing.T)  {
		word := "test"
		definition := justATestString

		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
	
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": justATestString}


	t.Run("known word", func (t *testing.T)  {
		got, err := dictionary.Search("test")
		want := justATestString

		assertStrings(t, got, want)
		assertError(t, err, nil)
	})
	
	t.Run("unknown word", func (t *testing.T)  {
		got, err := dictionary.Search("unknown")
		
		assertError(t, err, ErrCouldNotFindWord)
		assertStrings(t, got, "")
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition( t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(t, got, definition)
}
