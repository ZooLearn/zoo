package repository

import (
	"context"
	"time"

	"github.com/ZooLearn/zoo/domain"
	"github.com/ZooLearn/zoo/ent"
	"github.com/ZooLearn/zoo/ent/predicate"
	"github.com/ZooLearn/zoo/ent/tag"
)

type tagRepository struct {
	ctx      context.Context
	database *ent.Client
}

func NewTagRepository(db *ent.Client) domain.TagRepository {
	return &tagRepository{
		ctx:      context.Background(),
		database: db,
	}
}

func (tr *tagRepository) Create(c context.Context, tag domain.Tag) error {
	_, err := tr.database.Tag.Create().
		SetAlias(*tag.Alias).
		SetColor(*tag.Color).
		SetImageUrl(*tag.ImageUrl).
		SetKey(*tag.Key).Save(tr.ctx)
	return err
}
func (tr *tagRepository) Update(c context.Context, tag domain.Tag) error {
	err := tr.database.Tag.UpdateOneID(uint64(*tag.ID)).
		SetNillableAlias(tag.Alias).
		SetNillableColor(tag.Color).
		SetNillableImageUrl(tag.ImageUrl).
		SetNillableKey(tag.Key).
		Exec(tr.ctx)
	return err
}
func (tr *tagRepository) Delete(c context.Context, tagID uint64) error {
	err := tr.database.Tag.UpdateOneID(tagID).
		SetDeleteAt(time.Now()).
		Exec(tr.ctx)
	return err
}
func (tr *tagRepository) GetByID(c context.Context, tagID uint64) (*domain.Tag, error) {
	tag, err := tr.database.Tag.Query().Where(tag.IDEQ(uint64(tagID))).First(tr.ctx)
	if err != nil {
		return nil, err
	}
	return &domain.Tag{
		ID:       &tagID,
		Alias:    &tag.Alias,
		Color:    &tag.Color,
		ImageUrl: &tag.ImageUrl,
	}, nil
}

func (tr *tagRepository) Fetch(c context.Context, r domain.FetchTagRequest) ([]domain.Tag, *domain.Pagination, error) {
	predicate := []predicate.Tag{}
	predicate = append(predicate, tag.DeleteAtIsNil())
	if r.Key != nil {
		predicate = append(predicate, tag.KeyEQ(*r.Key))
	}
	tags, err := tr.database.Tag.Query().Where(predicate...).Page(tr.ctx, r.Page, r.PageSize, func(tp *ent.TagPager) {
		tp.Order = ent.Desc(tag.FieldCreatedAt)
	})
	if err != nil {
		return nil, nil, err
	}
	listTag := make([]domain.Tag, 0, len(tags.List))
	for _, tag := range tags.List {
		listTag = append(listTag, domain.Tag{
			ID:        &tag.ID,
			Key:       &tag.Key,
			Alias:     &tag.Alias,
			Color:     &tag.Color,
			ImageUrl:  &tag.ImageUrl,
			CreatedAt: &tag.CreatedAt,
			UpdateAt:  &tag.UpdatedAt,
		})
	}
	return listTag, &domain.Pagination{
		Page:     tags.PageDetails.Page,
		PageSize: tags.PageDetails.Size,
		Total:    tags.PageDetails.Total,
	}, nil
}

func (tr *tagRepository) FetchAll(c context.Context) ([]domain.Tag, error) {
	tags, err := tr.database.Tag.Query().Where(tag.DeleteAtIsNil()).All(tr.ctx)
	if err != nil {
		return nil, err
	}
	listTag := make([]domain.Tag, 0, len(tags))
	for _, tag := range tags {
		listTag = append(listTag, domain.Tag{
			ID:        &tag.ID,
			Key:       &tag.Key,
			Alias:     &tag.Alias,
			Color:     &tag.Color,
			CreatedAt: &tag.CreatedAt,
			UpdateAt:  &tag.UpdatedAt,
		})
	}
	return listTag, nil
}
