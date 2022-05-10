package client

type XMLReq interface {
	XMLRR() (User, error)
}

type User struct {
	Type    string `xml:"type"`
	Name    string `xml:"name"`
	ID      string `xml:"id"`
	Web     string `xml:"web"`
	Company string `xml:"company"`
}
