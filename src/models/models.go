package models


type Course struct {
	Title			string
	Url				string
	StartedAt		string
	FinishedAt		string
	Image			string
	TeachersName	[]string
	TeachersImages	[]string
	TeachersDescriptions	[]string
	Skills			string
	Description		string
	Requirements	string
	DurationInWeek	int64
	LecturesCount	int64
	HasCertificate	bool
}
