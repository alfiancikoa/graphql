package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/alfiancikoa/graphql-server/graph/generated"
	"github.com/alfiancikoa/graphql-server/graph/model"
)

func (r *mutationResolver) CreateBook(ctx context.Context, input model.NewBook) (*model.Book, error) {
	var AuthorId int
	if len(r.books) > 0 {
		AuthorId, _ = strconv.Atoi(r.books[len(r.books)-1].Author.ID)
	}
	book := &model.Book{
		ID:     fmt.Sprintf("%d", len(r.books)+1),
		Code:   input.Code,
		Title:  input.Title,
		Author: &model.Author{ID: fmt.Sprintf("%d", AuthorId+1), Name: input.AuthorName, Country: input.AuthorCountry},
	}
	r.books = append(r.books, book)
	return book, nil
}

func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	return r.books, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
