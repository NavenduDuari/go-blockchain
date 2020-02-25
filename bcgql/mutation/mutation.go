package mutation

import (
	"sync"

	"github.com/NavenduDuari/go-blockchain/blockchain"
	"github.com/graphql-go/graphql"
)

var mutex = &sync.Mutex{}

var RootMutation = graphql.ObjectConfig(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"MineBlock": &graphql.Field{
			Type: nil,
			Args: graphql.FieldConfigArgument{
				"Data": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				mutex.Lock()
				preBlock := blockchain.Blockchain[len(blockchain.Blockchain)-1]
				newBlock := blockchain.GenerateBlock(preBlock, params.Args["Data"].(int))
				if blockchain.IsBlockValid(newBlock, preBlock) {
					blockchain.Blockchain = append(blockchain.Blockchain, newBlock)
				}
				mutex.Unlock()

				return nil, nil
			},
		},
	},
})
