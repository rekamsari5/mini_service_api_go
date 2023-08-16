package customers

import "database/sql"

type Customers struct {
	ID        int64          `json:"id"`
	Name      string         `json:"name"`
	Address   string         `json:"address"`
	CreatedAt string         `json:"created_at"`
	UpdateAt  sql.NullString `json:"updated_at"`
}

type ResultRequest struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	CreatedAt string `json:"created_at"`
}

type ResultUpdate struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	UpdateAt string `json:"updated_at"`
}
