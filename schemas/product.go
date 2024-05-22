package schemas

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID              uuid.UUID
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Deleted         bool
	Name            string
	Type            string
	Price           float32
	DiscountedPrice float32
	DiscountFlag    bool
	Image           string
	Quantity        int
}
