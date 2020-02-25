package model

import (
	"github.com/graphql-go/graphql"
)

var Block = graphql.NewObject(graphql.ObjectConfig{
	Name: "Block",
	Fields: graphql.Fields{
		"Index": &graphql.Field{
			Type: graphql.String,
		},
		"Data": &graphql.Field{
			Type: graphql.Int,
		},
		"Timestamp": &graphql.Field{
			Type: graphql.String,
		},
		"Hash": &graphql.Field{
			Type: graphql.String,
		},
		"PrevHash": &graphql.Field{
			Type: graphql.String,
		},
	},
})
