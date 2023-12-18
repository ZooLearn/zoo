package domain

import (
	"context"
	"time"
)

type Tag struct {
	ID        *uint64
	Key       *string
	Alias     *string
	Color     *string
	ImageUrl  *string
	CreatedAt *time.Time
	UpdateAt  *time.Time
}

type TagRepository interface {
	Create(c context.Context, tag Tag) error
	Update(c context.Context, tag Tag) error
	Delete(c context.Context, tagID uint64) error
	GetByID(c context.Context, tagID uint64) (*Tag, error)
	Fetch(c context.Context, r FetchTagRequest) ([]Tag, *Pagination, error)
	FetchAll(c context.Context) ([]Tag, error)
}

type TagInfo struct {
	ID        *uint64   `json:"id"`
	Key       *string   `json:"key"`
	Alias     *string   `json:"alias"`
	Color     *string   `json:"color"`
	ImageUrl  *string   `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}

type CreateTagRequest struct {
	Key      string `json:"key"`
	Alias    string `json:"alias"`
	Color    string `json:"color"`
	ImageUrl string `json:"image_url"`
}

type UpdateTagRequest struct {
	ID       uint64  `json:"id"`
	Key      *string `json:"key"`
	Alias    *string `json:"alias"`
	Color    *string `json:"color"`
	ImageUrl *string `json:"image_url"`
}

type DeleteTagRequest struct {
	ID uint64 `path:"id"`
}

type GetTagByIDRequest struct {
	ID uint64 `path:"id"`
}

type FetchTagRequest struct {
	Page     uint64  `form:"page"`
	PageSize uint64  `form:"page_size"`
	Key      *string `form:"key"`
}

type FetchTagResponse struct {
	Response
	Data       []TagInfo `json:"data"`
	Pagination `json:"pagination"`
}

type FetchAllResponse struct {
	Response
	Data []TagInfo `json:"data"`
}

type GetByIDResponse struct {
	Response
	Data TagInfo `json:"data"`
}

type TagUsecase interface {
	Create(c context.Context, tag Tag) error
	Update(c context.Context, tag Tag) error
	Delete(c context.Context, tagID uint64) error
	GetByID(c context.Context, tagID uint64) (*Tag, error)
	Fetch(c context.Context, r FetchTagRequest) ([]Tag, *Pagination, error)
	FetchAll(c context.Context) ([]Tag, error)
}
