package movies

type HealthResponse struct {
	Code        int     `json:"code"`
	Status      string  `json:"status"`
	Description *string `json:"description"`
	Message     *string `json:"message"`
}

type MovieRequest struct {
	IdMovie    int    `json:"id_movie"`
	MovieName  string `json:"movie_name"`
	Category   string `json:"category"`
	ImageMovie string `json:"image_movie"` // Base64-encoded image data
}

type MovieResponse struct {
	IdMovie    int     `json:"id_movie"`
	MovieName  string  `json:"movie_name"`
	MovieScore float64 `json:"movie_score"`
	Category   string  `json:"category"`
	ImageMovie string  `json:"image_movie"`
}

type MovieUpdateRequest struct {
	MovieName  *string `json:"movie_name"`
	Category   *string `json:"category"`
	ImageMovie *string `json:"image_movie"`
}

type ScoreResponse struct {
	MovieScore float64 `json:"movie_score"`
}
