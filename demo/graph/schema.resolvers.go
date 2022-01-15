package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/mfcochauxlaberge/gqlgen-extra/demo/graph/gen"
	"github.com/mfcochauxlaberge/gqlgen-extra/demo/graph/model"
	"github.com/mfcochauxlaberge/gqlgen-extra/store"
)

func (r *articleResolver) Author(ctx context.Context, obj *model.Article) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *articleResolver) Comments(ctx context.Context, obj *model.Article) ([]model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *commentResolver) Author(ctx context.Context, obj *model.Comment) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *commentResolver) Article(ctx context.Context, obj *model.Comment) (*model.Article, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (model.CreateUserResult, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, input model.UserInput) (model.UserResult, error) {
	u := model.User{}

	err := store.Builder(ctx).Select("col1", "col2").From("users").Into(&u)
	if err != nil {
		return nil, err
	}

	// u := &model.User{
	// 	ID:       "user1",
	// 	Username: "user1",
	// 	Articles: []model.Article{
	// 		{
	// 			ID:      "article1",
	// 			Title:   "My Article",
	// 			Content: "This is the content.",
	// 			Author: &model.User{
	// 				ID: "user1",
	// 			},
	// 		},
	// 	},
	// }

	return u, nil
}

func (r *queryResolver) Article(ctx context.Context, input model.ArticleInput) (model.ArticleResult, error) {
	a := &model.Article{
		ID: "article1",
	}

	return a, nil
}

func (r *userResolver) Articles(ctx context.Context, obj *model.User) ([]model.Article, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Comments(ctx context.Context, obj *model.User) ([]model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

// Article returns gen.ArticleResolver implementation.
func (r *Resolver) Article() gen.ArticleResolver { return &articleResolver{r} }

// Comment returns gen.CommentResolver implementation.
func (r *Resolver) Comment() gen.CommentResolver { return &commentResolver{r} }

// Mutation returns gen.MutationResolver implementation.
func (r *Resolver) Mutation() gen.MutationResolver { return &mutationResolver{r} }

// Query returns gen.QueryResolver implementation.
func (r *Resolver) Query() gen.QueryResolver { return &queryResolver{r} }

// User returns gen.UserResolver implementation.
func (r *Resolver) User() gen.UserResolver { return &userResolver{r} }

type articleResolver struct{ *Resolver }
type commentResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
