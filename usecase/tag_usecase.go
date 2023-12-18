package usecase

import (
	"context"
	"time"

	"github.com/ZooLearn/zoo/domain"
)

type tagUsecase struct {
	tagRepository  domain.TagRepository
	contextTimeout time.Duration
}

func NewTagUsecase(tagRepository domain.TagRepository, timeout time.Duration) domain.TagUsecase {
	return &tagUsecase{
		tagRepository:  tagRepository,
		contextTimeout: timeout,
	}
}

func (tu *tagUsecase) Create(c context.Context, tag domain.Tag) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.tagRepository.Create(ctx, tag)
}

func (tu *tagUsecase) Update(c context.Context, tag domain.Tag) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.tagRepository.Update(ctx, tag)
}

func (tu *tagUsecase) Delete(c context.Context, tagID uint64) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.tagRepository.Delete(ctx, tagID)
}

func (tu *tagUsecase) GetByID(c context.Context, tagID uint64) (*domain.Tag, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.tagRepository.GetByID(ctx, tagID)
}

func (tu *tagUsecase) Fetch(c context.Context, r domain.FetchTagRequest) ([]domain.Tag, *domain.Pagination, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.tagRepository.Fetch(ctx, r)
}

func (tu *tagUsecase) FetchAll(c context.Context) ([]domain.Tag, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.tagRepository.FetchAll(ctx)
}
