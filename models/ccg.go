package models

type CCG struct {
	FirstName string `firestore:"FirstName"`
	LastName  string `firestore:"LastName"`
	Rank      string `firestore:"Rank"`
}
