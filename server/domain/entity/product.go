package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty"  bson:"name,omitempty"`
	Price       int64              `json:"price,omitempty"  bson:"price,omitempty"`
	Description string             `json:"description,omitempty"  bson:"description,omitempty"`
	Rating      int32              `json:"rating,omitempty"  bson:"rating,omitempty"`
	Images      []string           `json:"images,omitempty"  bson:"images,omitempty"`
	Seller      string             `json:"seller,omitempty"  bson:"seller,omitempty"`
	CreatedBy   string             `json:"createdBy,omitempty"  bson:"createdBy,omitempty"`
	UpdatedBy   string             `json:"updatedBy,omitempty"  bson:"updatedBy,omitempty"`
	UpdatedAt   time.Time          `json:"updatedAt,omitempty"  bson:"updatedAt,omitempty"`
	CreatedAt   time.Time          `json:"createdAt,omitempty"  bson:"createdAt,omitempty"`
}

func (p *Product) Validate() map[string]string {
	var errors map[string]string
	// price cannot be 0
	if p.Price <= 0 {
		errors["price"] = "Price cannot be 0 or less than 0"
	}
	// name cannot be null
	if p.Name == "" {
		errors["name"] = "Name cannot be null"
	}
	// seller cannot be null
	if p.Seller == "" {
		errors["seller"] = "Seller cannot be null"
	}
	// Description cannot be null
	if p.Description == "" {
		errors["description"] = "Description cannot be null"
	}
	// need atleast 1 image
	if len(p.Images) <= 0 {
		errors["images"] = "Provide atleast 1 image"
	}

	return errors
}

func (p *Product) PreSave(userid string) {
	p.CreatedBy = userid
	p.UpdatedAt = time.Now()
	p.CreatedAt = time.Now()
}

func (p *Product) PreUpdate(userid string) {
	p.UpdatedBy = userid
	p.UpdatedAt = time.Now()
}
