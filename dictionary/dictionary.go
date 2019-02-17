package dictionary

// Dictionary is a dictionary of strings
type Dictionary map[string]string

const (
	// ErrNotFound is an error for when the word can't be found
	ErrNotFound = DictErr("could not find the word you were looking for")

	//ErrWordExists is an error for when a word is already in the dictionary
	ErrWordExists = DictErr("cannot add word because it already exists")

	//ErrWordDoesNotExist is an error for when a key doesn't exist
	ErrWordDoesNotExist = DictErr("cannot update word because it does not exist")
)

// DictErr is an error message
type DictErr string

// Error returns a error
func (e DictErr) Error() string {
	return string(e)
}

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
func (d Dictionary) Add(word, definition string) error {

	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

// Update will update the definition for a given word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

//Delete deletes an entry from the dicitonary using the given key
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
