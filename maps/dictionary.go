package maps

const (
	//ErrNotFound is thrown when trying to search a nonexistent word
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	//ErrWordExists is thrown when trying to add an existing word
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	//ErrWordDoesNotExist is thrown when trying to update a nonexistent word
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

//DictionaryErr consolidates the erros thrown by Dictionary's operations
type DictionaryErr string

//Error allows DictionaryErr to implement error's interface
func (e DictionaryErr) Error() string {
	return string(e)
}

//Dictionary use maps to model a real world dictionary
type Dictionary map[string]string

//Search allows a client to search a word's definition
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

//Add enables a client to add a new word's definition
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

//Update enables a client to update a word's definition
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

//Update enables a client to delete a word's definition
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
