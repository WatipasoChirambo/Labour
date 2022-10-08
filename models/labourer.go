package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Labourer struct{
	Id primitive.ObjectID `json:"id,omitempty"`
	RegNo string `json:"RegNo"`
	FirstName string `json:"FirstName"`
	LastName string `json:"LastName"`
	Gender string `json:"Gender"`
	DistrictOfBirth string `json:"DistrictOfBirth"`
    Address string `json:"Address"`
	HomeAddress string `json:"HomeAddress"`
	Village string `json:"Village"`
	Ta string `json:"Ta"`
	District string `json:"District"`
	MaritalStatus string `json:"MaritalStatus"`
	PreviousResidence string `json:"PreviousResidence"`
	CurrentResidence string `json:"CurrentResidence"`
	Education string `json:"Education"`
	Institution string `json:"Institution"`
	Profession string `json:"Profession"`
	JobWanted string `json:"JobWanted"`
	PreviousJob string `json:"PreviousJob"`
	YearsUnemployed int64 `json:"YearsUnemployed"`
	ExperienceYears int64 `json:"ExperienceYears"`
	Training bool `json:"Training"`
	RegisteredElseWhere bool `json:"RegisteredElseWhere"`
	RegisteredNumber string `json:"RegisteredNumber"`
}


var Labourers = []Labourer{

}