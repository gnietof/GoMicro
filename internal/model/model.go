package model

type Member struct {
	Id        string
	FirstName string
	LastName  string
	EMail     string
}

type Location struct {
	WLC        string
	CampusId   string
	CampusName string
	Geo        string
	Country    string
	City       string
}
