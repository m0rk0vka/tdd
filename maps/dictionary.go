package maps

import "errors"

var ErrNotFound = errors.New("could not find the word you are looking for")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	val, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return val, nil
}
