package model

type UserData struct {
	Id         string `bson:"_id,omitempty" json:"id"`
	Name       string `bson:"name" json:"name"`
	Number     int    `bson:"number" json:"number"`
	Age        int    `bson:"age" json:"age"`
	BirthMonth int    `bson:"birthMonth" json:"birthMonth"`
}
