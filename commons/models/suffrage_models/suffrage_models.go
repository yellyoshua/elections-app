package suffrage_models

import (
	gql "github.com/graphql-go/graphql"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

type SuffrageModel struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	CampaignID  primitive.ObjectID `bson:"campaignId,omitempty" json:"campaignId"`
	Endorsement primitive.ObjectID `bson:"endorsement,omitempty" json:"endorsement"`
	Suffrage    primitive.ObjectID `bson:"suffrage,omitempty" json:"suffrage"`
	Created     int64              `bson:"created" json:"created"`
}

var GraphqlBallotModel = gql.NewObject(gql.ObjectConfig{
	Name: "Suffrage",
	Fields: gql.Fields{
		"_id": &gql.Field{
			Type: gql.String,
		},
		"institute": &gql.Field{ // -> Instituto ID
			Type: gql.String,
		},
		"endorsement": &gql.Field{ // -> Aprobación ID
			Type: gql.String,
		},
		"suffrage": &gql.Field{ // -> Suffrage ID == Voto ID
			Type: gql.String,
		},
		"created": &gql.Field{
			Type: gql.Int,
		},
	},
})
