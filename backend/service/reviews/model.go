package reviews

import "time"

type HealthResponse struct {
	Code        int     `json:"code"`
	Status      string  `json:"status"`
	Description *string `json:"description"`
	Message     *string `json:"message"`
}

type ReviewRequest struct {
	Name    string  `json:"name"`
	IdMovie int     `json:"id_movie"`
	Comment string  `json:"comment"`
	Score   float64 `json:"score"`
}

type ReviewUpdateRequest struct {
	Name    *string  `json:"name"`
	Comment *string  `json:"comment"`
	Score   *float64 `json:"score"`
}

type ReviewResponse struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	IdMovie  int       `json:"id_movie"`
	Comment  string    `json:"comment"`
	Score    float64   `json:"score"`
	CreateAt time.Time `json:"create_at"`
}
