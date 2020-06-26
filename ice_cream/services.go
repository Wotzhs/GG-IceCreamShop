package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/oklog/ulid"
)

var iceCreamService *IceCreamService

type IceCreamService struct{}

func (s *IceCreamService) CreateIceam(iceCream *IceCream) error {
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
		fmt.Sprintf("{%s}", strings.Join(iceCream.SourcingValues, ",")),
		fmt.Sprintf("{%s}", strings.Join(iceCream.Ingredients, ",")),
		iceCream.AllergyInfo,
		iceCream.DietaryCertifications,
		iceCream.ProductID,
		iceCream.CreatedBy,
		iceCream.UpdatedBy,
	).Scan(&iceCream.ID)

	return err
}
