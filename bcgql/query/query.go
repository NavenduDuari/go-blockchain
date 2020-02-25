package query

import (
	"github.com/NavenduDuari/go-blockchain/bcgql/model"
	"github.com/NavenduDuari/go-blockchain/blockchain"
	"github.com/graphql-go/graphql"
)

var RootQuery = graphql.ObjectConfig(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"DisplayChain": &graphql.Field{
			Type: graphql.NewList(model.Block),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return blockchain.Blockchain, nil
			},
		},
	},
})
