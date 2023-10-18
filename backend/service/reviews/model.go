package reviews

type ReviewRequest struct {
	Name    string  `json:"name"`
	IdMovie int     `json:"id_movie"`
	Comment string  `json:"comment"`
	Score   float64 `json:"score"`
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
