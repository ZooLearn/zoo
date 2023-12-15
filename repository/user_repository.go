package repository

import (
	"context"

	"github.com/ZooLearn/zoo/domain"
	"github.com/ZooLearn/zoo/ent"
	"github.com/ZooLearn/zoo/ent/user"
)

type userRepository struct {
	ctx        context.Context
	database   *ent.Client
	collection string
}

func NewUserRepository(db *ent.Client, collection string) domain.UserRepository {
	return &userRepository{
		ctx:        context.Background(),
		database:   db,
		collection: collection,
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

func (ur *userRepository) Fetch(c context.Context) ([]domain.User, error) {
	// collection := ur.database.Collection(ur.collection)

	// opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	// cursor, err := collection.Find(c, bson.D{}, opts)

	// if err != nil {
	// 	return nil, err
	// }

	// var users []domain.User

	// err = cursor.All(c, &users)
	// if users == nil {
	// 	return []domain.User{}, err
	// }

	// return users, err
	return nil, nil
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
	// collection := ur.database.Collection(ur.collection)

	// var user domain.User

	// idHex, err := primitive.ObjectIDFromHex(id)
	// if err != nil {
	// 	return user, err
	// }

	// err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&user)
	// return user, err
	return domain.User{}, nil
}
