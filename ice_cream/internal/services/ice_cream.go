package services

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"proto/ice_cream"

	"GG-IceCreamShop/ice_cream/internal/db"
	"GG-IceCreamShop/ice_cream/internal/models"

	"github.com/lib/pq"
	"github.com/oklog/ulid"
)

var IceCream *iceCreamService

type iceCreamService struct{}

func (s *iceCreamService) GetIceCreamByID(ID ulid.ULID, iceCream *models.IceCream) error {
	query := "SELECT * FROM ice_creams WHERE id = $1"
	return db.Conn.QueryRow(query, ID).Scan(
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
}

func (s *iceCreamService) GetIceCreams(q *ice_cream.IceCreamQuery, iceCreams *[]models.IceCream) error {
	baseQuery := "SELECT * FROM ice_creams\n"
	defaultLimit := "LIMIT 10\n"

	var queries []string
	var variables []interface{}

	if q.After != "" {
		queries = append(queries, fmt.Sprintf("id < $%d", len(queries)+1))
		ID, err := ulid.Parse(q.After)
		if err != nil {
			return err
		}
		variables = append(variables, ID)
	}

	if q.Name != "" {
		queries = append(queries, fmt.Sprintf("name LIKE '%%' || $%d || '%%'", len(queries)+1))
		variables = append(variables, q.Name)
	}

	if len(q.SourcingValues) > 0 {
		queries = append(queries, fmt.Sprintf("sourcing_values @> $%d", len(queries)+1))
		variables = append(variables, pq.Array(q.SourcingValues))
	}

	if len(q.Ingredients) > 0 {
		queries = append(queries, fmt.Sprintf("ingredients @> $%d", len(queries)+1))
		variables = append(variables, pq.Array(q.Ingredients))
	}

	query := baseQuery
	if len(queries) > 0 {
		query += "WHERE " + strings.Join(queries, " AND ") + "\n"
	}

	var sort string
	switch q.SortCol {
	case ice_cream.SortColumn_NAME:
		sort = "ORDER BY name "
	case ice_cream.SortColumn_CREATED_AT:
		sort = "ORDER BY created_at "
	case ice_cream.SortColumn_UPDATED_AT:
		sort = "ORDER BY udpated_at "
	}

	if len(sort) > 0 {
		if q.SortDir == ice_cream.SortDir_ASC {
			sort += "asc\n"
		} else {
			sort += "desc\n"
		}
	}

	limit := defaultLimit
	if q.First != 0 {
		limit = fmt.Sprintf("LIMIT $%d\n", len(variables)+1)
		variables = append(variables, q.First)
	}

	rows, err := db.Conn.Query(query+sort+limit, variables...)
	if err != nil {
		return err
	}

	for rows.Next() {
		iceCream := models.IceCream{}
		err := rows.Scan(
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

		if err != nil {
			return err
		}
		*iceCreams = append(*iceCreams, iceCream)
	}

	return nil
}

func (s *iceCreamService) GetIceCreamsCount(totalCount *int32) error {
	totalCountQuery := "SELECT count(id) from ice_creams;"
	if err := db.Conn.QueryRow(totalCountQuery).Scan(totalCount); err != nil {
		return err
	}
	return nil
}

func (s *iceCreamService) HasNextIceCreams(lastID ulid.ULID, hasNext *bool) error {
	hasNextQuery := "SELECT count(id)>0 from ice_creams WHERE id < $1"
	if err := db.Conn.QueryRow(hasNextQuery, lastID).Scan(hasNext); err != nil {
		return err
	}
	return nil
}

func (s *iceCreamService) CreateIceCream(iceCream *models.IceCream) error {
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

	err := db.Conn.QueryRow(
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

func (s *iceCreamService) UpdateIceCream(iceCream *models.IceCream) error {
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

	err := db.Conn.QueryRow(
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

func (s *iceCreamService) DeleteIceCream(iceCream *models.IceCream) error {
	query := `DELETE from ice_creams WHERE id = $1`
	_, err := db.Conn.Exec(query, iceCream.ID)
	return err
}

func (s *iceCreamService) Import(iceCreams []models.IceCream) (int64, error) {
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
		VALUES

	`
	var variables []interface{}
	cols := 13

	for i, iceCream := range iceCreams {
		t := time.Unix(time.Now().Unix(), time.Now().UnixNano())
		entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
		id := ulid.MustNew(ulid.Timestamp(t), entropy)

		query += fmt.Sprintf(
			"($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d),\n",
			cols*i+1,
			cols*i+2,
			cols*i+3,
			cols*i+4,
			cols*i+5,
			cols*i+6,
			cols*i+7,
			cols*i+8,
			cols*i+9,
			cols*i+10,
			cols*i+11,
			cols*i+12,
			cols*i+13,
		)
		variables = append(
			variables,
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
			"admin",
			"admin",
		)
	}

	query = strings.TrimSuffix(query, ",\n")
	query += "\nON CONFLICT (name) DO NOTHING"

	res, err := db.Conn.Exec(query, variables...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
