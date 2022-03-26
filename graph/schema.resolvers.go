package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/alfiancikoa/graphql-server/graph/generated"
	"github.com/alfiancikoa/graphql-server/graph/model"
)

// Fungsi untuk menambahkan buku baru
func (r *mutationResolver) CreateBook(ctx context.Context, input model.NewBook) (*model.Book, error) {
	// Fetch data new Book
	book := &model.Book{
		ID:     fmt.Sprintf("%d", len(r.books)+1),
		Code:   input.Code,
		Title:  input.Title,
		Author: &model.Author{Name: input.AuthorName, Country: input.AuthorCountry},
	}
	// Tambahkan ke dalam slice dengan menggunakan append
	r.books = append(r.books, book)
	return book, nil
}

// Fungsi untuk meng-update data buku
func (r *mutationResolver) UpdateBook(ctx context.Context, input model.NewBook, id string) (*model.Book, error) {
	var idBook int = -1
	for i := 0; i < len(r.books); i++ {
		if r.books[i].ID == id {
			idBook = i
			break
		}
	}
	if idBook == -1 {
		return nil, fmt.Errorf("book not found")
	}
	r.books[idBook] = &model.Book{
		ID:     id,
		Code:   input.Code,
		Title:  input.Title,
		Author: &model.Author{Name: input.AuthorName, Country: input.AuthorCountry},
	}
	return r.books[idBook], nil
}

// Fungsi untuk meenampilkan seluruh buku
func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	return r.books, nil
}

// Fungsi untuk menampilkan informasi buku pada id buku tertentu
func (r *queryResolver) BookID(ctx context.Context, id string) (*model.Book, error) {
	for i := 0; i < len(r.books); i++ {
		if r.books[i].ID == id {
			return r.books[i], nil
		}
	}
	return nil, fmt.Errorf("file book not found")
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
