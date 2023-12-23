package contact

type Contact struct {
	Name string `query:"name"`
}

func Search(term string) (contacts []Contact, err error) {
	contacts = append(contacts, Contact{Name: "John"})
	return contacts, nil
}

func All() (contacts []Contact, err error) {
	contacts = append(contacts, Contact{Name: "John"})
	contacts = append(contacts, Contact{Name: "Jane"})
	return contacts, nil
}
