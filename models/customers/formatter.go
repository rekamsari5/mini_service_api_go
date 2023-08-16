package customers

import "time"

type FormatResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"updated_at"`
}

func Formatter(data []Customers) []FormatResponse {
	results := []FormatResponse{}

	for _, v := range data {
		createdAt, _ := time.Parse("2006-01-02T15:04:05Z", v.CreatedAt)
		updateAt, _ := time.Parse("2006-01-02T15:04:05Z", v.UpdateAt.String)

		result := FormatResponse{
			ID:        v.ID,
			Name:      v.Name,
			Address:   v.Address,
			CreatedAt: createdAt.Format("2006-01-02 15:04:05"),
			UpdateAt:  updateAt.Format("2006-01-02 15:04:05"),
		}
		results = append(results, result)

	}

	return results
}
