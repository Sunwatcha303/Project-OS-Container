-- FINAL DRAFT

DROP TABLE review;
DROP TABLE movie;
DROP TABLE actor;
DROP TABLE movie_actor;

CREATE TABLE IF NOT EXISTS actor(
    id_actor INT PRIMARY KEY,
    actor_name VARCHAR(30),
    sex VARCHAR(1)
);

CREATE TABLE IF NOT EXISTS movie (
    id_movie INT PRIMARY KEY,
    movie_name VARCHAR(50),
    movie_score FLOAT DEFAULT 0,
    category VARCHAR(30),
    image_movie LONGBLOB
);

CREATE TABLE IF NOT EXISTS movie_actor (
    id_movie INT,
    id_actor INT,
    PRIMARY KEY (id_movie,id_actor)
);

CREATE TABLE IF NOT EXISTS review (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(30),
    id_movie INT,
    comment VARCHAR(200),
    score FLOAT,
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id_movie) REFERENCES movie(id_movie)
);