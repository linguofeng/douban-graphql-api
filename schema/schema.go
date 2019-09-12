package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/linguofeng/douban-graphql-api/types"
)

// New schema
func New() graphql.Schema {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: types.QueryType,
	})
	if err != nil {
		panic(err)
	}
	return schema
}
