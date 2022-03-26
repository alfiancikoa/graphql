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
	// Fetch data new Book
	var bookId int
	if len(r.books) > 0 {
		bookId, _ = strconv.Atoi(r.books[len(r.books)-1].ID)
	}
	book := &model.Book{
		ID:     fmt.Sprintf("%d", bookId+1),
		Code:   input.Code,
		Title:  input.Title,
		Author: &model.Author{Name: input.AuthorName, Country: input.AuthorCountry},
	}
	// Tambahkan ke dalam slice dengan menggunakan append
	r.books = append(r.books, book)
	return book, nil
}

func (r *mutationResolver) UpdateBook(ctx context.Context, input model.NewBook, id string) (*model.Book, error) {
	var idBook int = -1
	// Cari buku yang id bukunya sama dengan id
	for i := 0; i < len(r.books); i++ {
		if r.books[i].ID == id {
			idBook = i
			break
		}
	}
	// Jika id buku tidak ditemukan
	if idBook == -1 {
		return nil, fmt.Errorf("book not found")
	}
	// Jika buku ditemukan, perbaharui isi informasi buku
	r.books[idBook] = &model.Book{
		ID:     id,
		Code:   input.Code,
		Title:  input.Title,
		Author: &model.Author{Name: input.AuthorName, Country: input.AuthorCountry},
	}
	return r.books[idBook], nil
}

func (r *mutationResolver) DeleteBook(ctx context.Context, id string) (*model.Book, error) {
	for i := 0; i < len(r.books); i++ {
		if r.books[i].ID == id {
			delBook := r.books[i]
			r.books = append(r.books[:i], r.books[i+1:]...)
			return delBook, nil
		}
	}
	return nil, fmt.Errorf("book not found")
}

func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	return r.books, nil
}

func (r *queryResolver) BookID(ctx context.Context, id string) (*model.Book, error) {
	// Cari buku dengan id buku sama dengan id
	for i := 0; i < len(r.books); i++ {
		if r.books[i].ID == id {
			// Jika ketemu, kembalikan informasi buku
			return r.books[i], nil
		}
	}
	// Jika buku tidak ditemukan, kembalikan  pesan error
	return nil, fmt.Errorf("file book not found")
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
