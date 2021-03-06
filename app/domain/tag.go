package domain

import (
	"context"
	"time"
)

// Tag ...
type Tag struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name,omitempty"`
	UpdatedAt time.Time `json:"-,omitempty"`
	CreatedAt time.Time `json:"-,omitempty"`
}

// TagUseCase represent the Tag's repository contract
type TagUseCase interface {
	GetAll(ctx context.Context, start int, limit int) ([]Tag, error)
	GetByID(ctx context.Context, id int64) (Tag, error)
	UpdateTag(ctx context.Context, Tag *Tag, id int64) error
	CreateTag(ctx context.Context, Tag *Tag) error
}

// TagRepository represent the Tag's repository
type TagRepository interface {
	GetAll(ctx context.Context, start int, limit int) ([]Tag, error)
	GetByID(ctx context.Context, id int64) (Tag, error)
	UpdateTag(ctx context.Context, Tag *Tag) error
	CreateTag(ctx context.Context, Tag *Tag) error
	DeleteTag(ctx context.Context, id int64) error
}
