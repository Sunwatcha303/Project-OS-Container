-- DRAFT 1

CREATE TABLE IF NOT EXISTS actor(
    id_actor INT PRIMARY KEY,
    actor_name VARCHAR(30)
);

CREATE TABLE IF NOT EXISTS category(
    id_category INT PRIMARY KEY,
    th_name VARCHAR(30),
    en_name VARCHAR(30)
);

CREATE TABLE IF NOT EXISTS movie (
    id_movie INT PRIMARY KEY,
    movie_name VARCHAR(50),
    id_actor INT,
    id_category INT,
    movie_score FLOAT,
    FOREIGN KEY (id_actor) REFERENCES actor(id_actor),
    FOREIGN KEY (id_category) REFERENCES category(id_category)
);

CREATE TABLE IF NOT EXISTS review (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(30),
    id_movie INT,
    comment VARCHAR(200),
    score FLOAT,
    time TIMESTAMP,
    FOREIGN KEY (id_movie) REFERENCES movie(id_movie)
);