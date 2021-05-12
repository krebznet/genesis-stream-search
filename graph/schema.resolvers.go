package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"genesis-stream-search/graph/generated"
	"genesis-stream-search/graph/model"
	"time"
)

func (r *queryResolver) Search(ctx context.Context) ([]*model.Signal, error) {
	return r.signals, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return newQueryResolver(r) }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func newQueryResolver(r *Resolver) *queryResolver {
	// init the signals for the resolver
	r.signals = append(r.signals, []*model.Signal{
		{
			ID:        1,
			Timestamp: time.Now().Format("2006-01-02 15:04:05.000"),
			Name:      "GOOG",
			Vars: []*model.Variable{
				{
					ID: 1,
					Value: model.IntegerType{
						IntegerValue: 520,
					},
				},
				{
					ID: 2,
					Value: model.FloatType{
						FloatValue: 10520.42,
					},
				},
			},
		},
		{
			ID:        2,
			Timestamp: time.Now().Format("2006-01-02 15:04:05.000"),
			Name:      "FB",
			Vars: []*model.Variable{
				{
					ID: 1,
					Value: model.StringType{
						StringValue: "facebook",
					},
				},
				{
					ID: 2,
					Value: model.BoolType{
						BoolValue: true,
					},
				},
			},
		},
	}...)

	return &queryResolver{r}
}
