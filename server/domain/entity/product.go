package entity

import "time"

type Product struct {
	ID          string    `json:"id" bson:"id"`
	Name        string    `json:"name"  bson:"name"`
	Price       int64     `json:"price"  bson:"price"`
	Description string    `json:"description"  bson:"description"`
	Rating      int32     `json:"rating"  bson:"rating"`
	Images      []string  `json:"images"  bson:"images"`
	Seller      string    `json:"seller"  bson:"seller"`
	CreatedBy   string    `json:"createdBy"  bson:"createdBy"`
	UpdatedBy   string    `json:"updatedBy"  bson:"updatedBy"`
	UpdatedAt   time.Time `json:"updatedAt"  bson:"updatedAt"`
	CreatedAt   time.Time `json:"createdAt"  bson:"createdAt"`
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
