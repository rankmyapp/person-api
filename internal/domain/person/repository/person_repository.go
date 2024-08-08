package repository

import (
	"context"

	"github.com/rpuglielli/person-api/internal/domain/person/entity"
)

type PersonRepository interface {
	Create(ctx context.Context, person *entity.Person) error
	Update(ctx context.Context, person *entity.Person) error
	Delete(ctx context.Context, id string) error
	FindByID(ctx context.Context, id string) (*entity.Person, error)
	FindAll(ctx context.Context, page, pageSize int) ([]*entity.Person, int64, error)
	FindByEmail(ctx context.Context, email string) (*entity.Person, error)
}
