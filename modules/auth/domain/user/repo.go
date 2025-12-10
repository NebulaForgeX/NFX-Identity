package user

import (
	"context"

	"github.com/google/uuid"
)

type Repo interface {
	Create(ctx context.Context, u *User) error
	Update(ctx context.Context, u *User) error
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
	GetByUsername(ctx context.Context, username string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByPhone(ctx context.Context, phone string) (*User, error)
	Exists(ctx context.Context, id uuid.UUID) (bool, error)
	ExistsByUsername(ctx context.Context, username string) (bool, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	ExistsByPhone(ctx context.Context, phone string) (bool, error)
	UpdatePassword(ctx context.Context, userID uuid.UUID, hashedPassword string) error
	UpdateStatus(ctx context.Context, userID uuid.UUID, status string) error
	Delete(ctx context.Context, userID uuid.UUID) error
}

