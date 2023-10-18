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

type ReviewResponse struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	IdMovie int     `json:"id_movie"`
	Comment string  `json:"comment"`
	Score   float64 `json:"score"`
	Time    time.Time `json:"time"`
}

// CREATE TABLE IF NOT EXISTS review (
//     id INT PRIMARY KEY AUTO_INCREMENT,
//     name VARCHAR(30),
//     id_movie INT,
//     comment VARCHAR(200),
//     score FLOAT,
//     time TIMESTAMP,
//     FOREIGN KEY (id_movie) REFERENCES movie(id_movie)
// );
