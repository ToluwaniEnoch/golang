package main

import "errors"

var (
	ErrCouldNotFindWord = errors.New("could not find the word you were looking for")
 	ErrWordExists = errors.New("word already exists") 
)


func(d Dictionary) Delete (word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrCouldNotFindWord:
		return ErrCouldNotFindWord
	case nil:
		delete(d, word)
		return nil
	default:
		return err
	}
}

func (d Dictionary) Update (word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrCouldNotFindWord:
		return ErrCouldNotFindWord
	case nil:
		d[word] = newDefinition
		return nil
	default:
		return err
	}

}

func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}

type Dictionary map[string]string

type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

func (d Dictionary) Search (word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrCouldNotFindWord
	}

	return definition, nil
}

func (d Dictionary) Add (word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrCouldNotFindWord:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}