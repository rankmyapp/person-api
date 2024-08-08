package usecase

import (
	"context"
	"time"

	"github.com/rpuglielli/person-api/internal/domain/person/entity"
	"github.com/rpuglielli/person-api/internal/domain/person/repository"
	"github.com/rpuglielli/person-api/pkg/errors"
)

type PersonUseCase struct {
	repo repository.PersonRepository
}

func NewPersonUseCase(repo repository.PersonRepository) *PersonUseCase {
	return &PersonUseCase{repo: repo}
}

func (uc *PersonUseCase) Create(ctx context.Context, person *entity.Person) error {
	if err := person.Validate(); err != nil {
		return err
	}

	existingPerson, err := uc.findPersonByEmail(ctx, person.Email)
	if err != nil {
		return err
	}
	if existingPerson != nil {
		return errors.NewConflictError("Person with this email already exists")
	}

	person.Created = time.Now()
	person.Updated = time.Now()

	return uc.repo.Create(ctx, person)
}

func (uc *PersonUseCase) Update(ctx context.Context, person *entity.Person) error {
	if err := person.Validate(); err != nil {
		return err
	}

	existingPerson, err := uc.findPersonByID(ctx, person.ID)
	if err != nil {
		return err
	}

	if person.Email != existingPerson.Email {
		emailCheck, err := uc.findPersonByEmail(ctx, person.Email)
		if err != nil {
			return err
		}
		if emailCheck != nil {
			return errors.NewConflictError("Email is already in use")
		}
	}

	person.Created = existingPerson.Created
	person.Updated = time.Now()

	return uc.repo.Update(ctx, person)
}

func (uc *PersonUseCase) Delete(ctx context.Context, id string) error {
	_, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		if errors.IsNotFoundError(err) {
			return errors.NewNotFoundError("Person not found")
		}
		return errors.NewInternalError("Failed to check existing person")
	}

	return uc.repo.Delete(ctx, id)
}

func (uc *PersonUseCase) FindByID(ctx context.Context, id string) (*entity.Person, error) {
	person, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		if errors.IsNotFoundError(err) {
			return nil, errors.NewNotFoundError("Person not found")
		}
		return nil, errors.NewInternalError("Failed to find person")
	}

	return person, nil
}

func (uc *PersonUseCase) FindAll(ctx context.Context, page, pageSize int) ([]*entity.Person, int64, error) {
	persons, total, err := uc.repo.FindAll(ctx, page, pageSize)
	if err != nil {
		return nil, 0, errors.NewInternalError("Failed to fetch persons")
	}

	return persons, total, nil
}

func (uc *PersonUseCase) findPersonByID(ctx context.Context, id string) (*entity.Person, error) {
	person, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		if errors.IsNotFoundError(err) {
			return nil, errors.NewNotFoundError("Person not found")
		}
		return nil, errors.NewInternalError("Failed to find person")
	}
	return person, nil
}

func (uc *PersonUseCase) findPersonByEmail(ctx context.Context, email string) (*entity.Person, error) {
	person, err := uc.repo.FindByEmail(ctx, email)
	if err != nil && !errors.IsNotFoundError(err) {
		return nil, errors.NewInternalError("Failed to check email uniqueness")
	}
	return person, nil
}
