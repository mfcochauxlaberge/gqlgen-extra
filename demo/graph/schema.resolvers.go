package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/mfcochauxlaberge/gqlgen-extra/demo/graph/gen"
	"github.com/mfcochauxlaberge/gqlgen-extra/demo/graph/model"
	"github.com/mfcochauxlaberge/gqlgen-extra/store"
)

func (r *articleResolver) Author(ctx context.Context, obj *model.Article) (*model.User, error) {
	return nil, nil
}

func (r *articleResolver) Comments(ctx context.Context, obj *model.Article) ([]model.Comment, error) {
	return nil, nil
}

func (r *commentResolver) Author(ctx context.Context, obj *model.Comment) (*model.User, error) {
	return nil, nil
}

func (r *commentResolver) Article(ctx context.Context, obj *model.Comment) (*model.Article, error) {
	return nil, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (model.CreateUserResult, error) {
	return nil, nil
}

func (r *mutationResolver) PublishArticle(ctx context.Context, input model.PublishArticleInput) (model.PublishArticleResult, error) {
	return nil, nil
}

func (r *mutationResolver) TagArticle(ctx context.Context, input model.TagArticleInput) (model.TagArticleResult, error) {
	return nil, nil
}

func (r *mutationResolver) UntagArticle(ctx context.Context, input model.UntagArticleInput) (model.UntagArticleResult, error) {
	return nil, nil
}

func (r *mutationResolver) LikeArticle(ctx context.Context, input model.LikeArticleInput) (model.LikeArticleResult, error) {
	return nil, nil
}

func (r *mutationResolver) UnlikeArticle(ctx context.Context, input model.UnlikeArticleInput) (model.UnlikeArticleResult, error) {
	return nil, nil
}

func (r *queryResolver) User(ctx context.Context, input model.UserInput) (model.UserResult, error) {
	u := model.User{}

	err := store.With(ctx).GetOne(
		store.ByID(input.ID),
		store.Into(&u),
	)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *queryResolver) Article(ctx context.Context, input model.ArticleInput) (model.ArticleResult, error) {
	a := &model.Article{
		ID: "article1",
	}

	return a, nil
}

func (r *queryResolver) MostLikedArticlesByTag(ctx context.Context, input model.MostLikedArticlesByTagInput) (model.MostLikedArticlesByTagResult, error) {
	a := []model.Article{}

	err := store.With(ctx).GetMany(
		store.Collection("articles"),
		store.Into(&a),
	)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *userResolver) Articles(ctx context.Context, obj *model.User) ([]model.Article, error) {
	a := []model.Article{}

	log.Printf("obj: %+v\n", obj)

	err := store.With(ctx).GetOne(
		store.FromRelationship(obj),
		store.Into(&a),
	)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (r *userResolver) Articles2(ctx context.Context, obj *model.User) ([]model.Article, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Comments(ctx context.Context, obj *model.User) ([]model.Comment, error) {
	return nil, nil
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
