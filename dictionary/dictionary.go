package dictionary

import "errors"

// Dictionary is a dictionary of strings
type Dictionary map[string]string

// ErrNotFound is an error for when the key is not found in the dictionary
var ErrNotFound = errors.New("could not find the word you were looking for")

// Search takes a dictionary and a search key and returns the value from the dictionary for that key
func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}

// Search searches the dictionary for the word key and returns the value
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add adds the given definition into the dictionary with word key
func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
