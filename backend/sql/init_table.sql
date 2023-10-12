CREATE TABLE actor(
    id_actor INT PRIMARY KEY,
    actor_name VARCHAR(30)
);

CREATE TABLE category(
    id_category INT PRIMARY KEY,
    th_name VARCHAR(30),
    en_name VARCHAR(30)
)

CREATE TABLE movie(
    id_moive INT PRIMARY KEY,
    movie_name VARCHAR(50),
    id_actor INT FOREIGN KEY REFERENCES actor(id_actor),
    id_category INT FOREIGN KEY REFERENCES category(id_category),
    movie_score FLOAT
);

CREATE TABLE review(
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(30),
    id_movie INT FOREIGN KEY REFERENCES movie(id_moive)
    comment VARCHAR(200),
    score FLOAT,
    time TIMESTAMP
)