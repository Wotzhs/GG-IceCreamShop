package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/oklog/ulid"
)

type IceCream struct {
	ID                    ulid.ULID
	Name                  string
	ImageClosed           string
	ImageOpen             string
	Description           string
	Story                 string
	SourcingValues        []string
	Ingredients           []string
	AllergyInfo           string
	DietaryCertifications string
	ProductID             string
	CreatedBy             string
	UpdatedBy             string
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

func (i *IceCream) Validate() error {
	errStr := []string{}

	if i.Name == "" {
		errStr = append(errStr, "name must not be empty")
	}

	if i.ImageClosed == "" {
		errStr = append(errStr, "image_closed must not be empty")
	}

	if i.ImageOpen == "" {
		errStr = append(errStr, "image_open must not be empty")
	}

	if i.Description == "" {
		errStr = append(errStr, "description must not be empty")
	}

	if i.Story == "" {
		errStr = append(errStr, "story must not be empty")
	}

	if len(errStr) > 0 {
		return fmt.Errorf("%v", strings.Join(errStr[:], ", "))
	}

	return nil
}
