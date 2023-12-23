package contact

type Contact struct {
	Id    int    `query:"id"`
	First string `query:"first"`
	Last  string `query:"last"`
	Phone string `query:"phone"`
	Email string `query:"email"`
}

var joe = Contact{
	Id:    1,
	First: "John",
	Last:  "Doe",
	Phone: "555-555-5555",
	Email: "jdoe@example.com",
}

var jane = Contact{
	Id:    2,
	First: "Jane",
	Last:  "Doe",
	Phone: "555-555-6666",
	Email: "janedoe@example.com",
}

func Search(term string) (contacts []Contact, err error) {
	contacts = append(contacts, joe)
	return contacts, nil
}

func All() (contacts []Contact, err error) {
	contacts = append(contacts, joe)
	contacts = append(contacts, jane)
	return contacts, nil
}
