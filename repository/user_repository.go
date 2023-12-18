package repository

import (
	"context"

	"github.com/ZooLearn/zoo/domain"
	"github.com/ZooLearn/zoo/ent"
	"github.com/ZooLearn/zoo/ent/user"
	"github.com/google/uuid"
)

type userRepository struct {
	ctx      context.Context
	database *ent.Client
}

func NewUserRepository(db *ent.Client) domain.UserRepository {
	return &userRepository{
		ctx:      context.Background(),
		database: db,
	}
}

func (ur *userRepository) Create(c context.Context, user *domain.User) error {
	_, err := ur.database.User.Create().
		SetEmail(user.Email).
		SetName(user.Name).
		SetPassword(user.Password).
		Save(ur.ctx)
	return err
}

func (ur *userRepository) GetByEmail(c context.Context, email string) (domain.User, error) {
	user, err := ur.database.User.Query().Where(user.EmailEQ(email)).First(ur.ctx)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		ID:       user.ID.String(),
		Name:     user.Email,
		Email:    user.Email,
		Password: user.Password,
	}, nil

}

func (ur *userRepository) GetByID(c context.Context, id string) (domain.User, error) {
	user, err := ur.database.User.Query().Where(user.IDEQ(uuid.MustParse(id))).First(ur.ctx)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		ID:    user.ID.String(),
		Name:  user.Email,
		Email: user.Email,
	}, nil
}
