package main

import (
	"math/rand"
	"time"

	"github.com/lib/pq"
	"github.com/oklog/ulid"
)

var iceCreamService *IceCreamService

type IceCreamService struct{}

func (s *IceCreamService) CreateIceCream(iceCream *IceCream) error {
	query := `
		INSERT INTO ice_creams (
			id,
			name,
			image_closed,
			image_open,
			description,
			story,
			sourcing_values,
			ingredients,
			allergy_info,
			dietary_certifications,
			product_id,
			created_by,
			updated_by
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		RETURNING id
	`

	t := time.Unix(time.Now().Unix(), time.Now().UnixNano())
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)

	err := db.QueryRow(
		query,
		id,
		iceCream.Name,
		iceCream.ImageClosed,
		iceCream.ImageOpen,
		iceCream.Description,
		iceCream.Story,
		pq.Array(iceCream.SourcingValues),
		pq.Array(iceCream.Ingredients),
		iceCream.AllergyInfo,
		iceCream.DietaryCertifications,
		iceCream.ProductID,
		iceCream.CreatedBy,
		iceCream.UpdatedBy,
	).Scan(&iceCream.ID)

	return err
}

func (s *IceCreamService) UpdateIceCream(iceCream *IceCream) error {
	query := `
		UPDATE ice_creams SET
			name = $1,
			image_closed = $2,
			image_open = $3,
			description = $4,
			story = $5,
			sourcing_values = $6,
			ingredients = $7,
			allergy_info = $8,
			dietary_certifications = $9,
			product_id = $10,
			updated_by = $11,
			updated_at = $12
		WHERE ID = $13
		RETURNING *
	`

	err := db.QueryRow(
		query,
		iceCream.Name,
		iceCream.ImageClosed,
		iceCream.ImageOpen,
		iceCream.Description,
		iceCream.Story,
		pq.Array(iceCream.SourcingValues),
		pq.Array(iceCream.Ingredients),
		iceCream.AllergyInfo,
		iceCream.DietaryCertifications,
		iceCream.ProductID,
		iceCream.UpdatedBy,
		time.Now(),
		iceCream.ID,
	).Scan(
		&iceCream.ID,
		&iceCream.Name,
		&iceCream.ImageClosed,
		&iceCream.ImageOpen,
		&iceCream.Description,
		&iceCream.Story,
		pq.Array(&iceCream.SourcingValues),
		pq.Array(&iceCream.Ingredients),
		&iceCream.AllergyInfo,
		&iceCream.DietaryCertifications,
		&iceCream.ProductID,
		&iceCream.CreatedBy,
		&iceCream.UpdatedBy,
		&iceCream.CreatedAt,
		&iceCream.UpdatedAt,
	)

	return err
}
