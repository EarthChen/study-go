package dictionary

type Dict map[string]string

type DictErr string

const (
	ErrNotFound         = DictErr("could not find the word you were looking for")
	ErrWordExists       = DictErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictErr("cannot update word because it does not exist")
)

func (e DictErr) Error() string {
	return string(e)
}

func (d Dict) Update(word, definition string) error {
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

func (d Dict) Search(word string) (string, error) {
	val, ok := d[word]
	if ok {
		return val, nil
	}
	return val, ErrNotFound
}

func (d Dict) Add(key string, val string) error {
	_, err := d.Search(key)
	if err == ErrNotFound {
		d[key] = val
		return nil
	} else if err == nil {
		return err
	} else {
		return err
	}
}

func (d Dict) Delete(word string) {
	delete(d, word)
}

func Search(dict map[string]string, word string) string {
	return dict[word]
}
