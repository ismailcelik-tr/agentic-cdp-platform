package domain

import (
	"context"
	"time"
)

// Customer represents the core domain model for a customer in our CDP.
type Customer struct {
	ID        string            `json:"id"`
	Email     string            `json:"email"`
	FirstName string            `json:"first_name"`
	LastName  string            `json:"last_name"`
	Metadata  map[string]string `json:"metadata"` // Custom attributes
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

// Segment represents a group of customers defined by specific rules.
type Segment struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Rules       string    `json:"rules"` // JSON query representation
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CustomerRepository defines the contracts/interfaces for persistence.
// Following DDD, the domain layer owns the repository interfaces.
type CustomerRepository interface {
	Create(ctx context.Context, customer *Customer) error
	GetByID(ctx context.Context, id string) (*Customer, error)
	GetByEmail(ctx context.Context, email string) (*Customer, error)
	Update(ctx context.Context, customer *Customer) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, offset, limit int) ([]*Customer, error)
}
