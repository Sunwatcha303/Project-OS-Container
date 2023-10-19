-- create test data in database --


INSERT INTO movie (id_movie,movie_name,category) VALUES (1,'SUN','TEST');
INSERT INTO movie (id_movie,movie_name,category) VALUES (2,'DEAR','NAN');

INSERT INTO review (name,id_movie,comment,score,time) VALUES ('SAVE',1,'กาก',1,NOW());
INSERT INTO review (name,id_movie,comment,score,time) VALUES ('SAVE',1,'กาก',5,NOW());
INSERT INTO review (name,id_movie,comment,score,time) VALUES ('SAVE',1,'กาก',3,NOW());
INSERT INTO review (name,id_movie,comment,score,time) VALUES ('SAVE',1,'กาก',4,NOW());
INSERT INTO review (name,id_movie,comment,score,time) VALUES ('SAVE',1,'กาก',3,NOW());

CALL Upd_Movie_Score('1')

select * from movie