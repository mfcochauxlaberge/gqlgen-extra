// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type ArticleResult interface {
	IsArticleResult()
}

type CreateUserResult interface {
	IsCreateUserResult()
}

type LikeArticleResult interface {
	IsLikeArticleResult()
}

type MostLikedArticlesByTagResult interface {
	IsMostLikedArticlesByTagResult()
}

type PublishArticleResult interface {
	IsPublishArticleResult()
}

type TagArticleResult interface {
	IsTagArticleResult()
}

type UnlikeArticleResult interface {
	IsUnlikeArticleResult()
}

type UntagArticleResult interface {
	IsUntagArticleResult()
}

type UserResult interface {
	IsUserResult()
}

type Article struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    *User     `json:"author"`
	Tags      []Tag     `json:"tags"`
	Likes     []User    `json:"likes"`
	Comments  []Comment `json:"comments"`
}

func (Article) IsArticleResult()        {}
func (Article) IsPublishArticleResult() {}

type ArticleInput struct {
	ID string `json:"id"`
}

type ArticleList struct {
	Articles []*Article `json:"articles"`
}

func (ArticleList) IsMostLikedArticlesByTagResult() {}

type Comment struct {
	ID      string   `json:"id"`
	Content string   `json:"content"`
	Author  *User    `json:"author"`
	Article *Article `json:"article"`
}

type CreateUserInput struct {
	Username string `json:"username"`
}

type Like struct {
	User    *User    `json:"user"`
	Article *Article `json:"article"`
}

func (Like) IsLikeArticleResult() {}

type LikeArticleInput struct {
	ArticleID string `json:"articleId"`
	Label     string `json:"label"`
}

type MostLikedArticlesByTagInput struct {
	Limit int    `json:"limit"`
	Tag   string `json:"tag"`
}

type NotAuthorized struct {
	Reason *string `json:"reason"`
}

func (NotAuthorized) IsUserResult()                   {}
func (NotAuthorized) IsArticleResult()                {}
func (NotAuthorized) IsMostLikedArticlesByTagResult() {}
func (NotAuthorized) IsCreateUserResult()             {}
func (NotAuthorized) IsPublishArticleResult()         {}
func (NotAuthorized) IsTagArticleResult()             {}
func (NotAuthorized) IsUntagArticleResult()           {}
func (NotAuthorized) IsLikeArticleResult()            {}
func (NotAuthorized) IsUnlikeArticleResult()          {}

type PublishArticleInput struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

type Success struct {
	Reason string `json:"reason"`
}

func (Success) IsUntagArticleResult()  {}
func (Success) IsUnlikeArticleResult() {}

type Tag struct {
	Label    string    `json:"label"`
	Articles []Article `json:"articles"`
}

func (Tag) IsTagArticleResult() {}

type TagArticleInput struct {
	ArticleID string `json:"articleId"`
	Label     string `json:"label"`
}

type UnlikeArticleInput struct {
	ArticleID string `json:"articleId"`
	Label     string `json:"label"`
}

type UntagArticleInput struct {
	ArticleID string `json:"articleId"`
	Label     string `json:"label"`
}

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Articles  []Article `json:"articles"`
	Articles2 []Article `json:"articles2"`
	Comments  []Comment `json:"comments"`
	Likes     []Article `json:"likes"`
}

func (User) IsUserResult()       {}
func (User) IsCreateUserResult() {}

type UserInput struct {
	ID string `json:"id"`
}