package dictionary

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr(
		"cannot perform operation on word because it does not exist",
	)
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	if d.exists(word) {
		return ErrWordExists
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	if !d.exists(word) {
		return ErrWordDoesNotExist
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Delete(word string) error {
	if !d.exists(word) {
		return ErrWordDoesNotExist
	}
	delete(d, word)
	return nil
}

func (d Dictionary) exists(word string) bool {
	_, exists := d[word]
	return exists
}
