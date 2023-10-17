-- PROCEDURE --
DELIMITER //

CREATE PROCEDURE Upd_Movie_Score(v_id_movie INT)
BEGIN
    UPDATE movie
    SET movie_score = (
        SELECT AVG(review.score) FROM review
        WHERE review.id_movie = movie.id_movie
    )
    WHERE id_movie = v_id_movie;
END //

DELIMITER ;

-- HOW TO CALL PROCEDURE Upd_Movie_Score(v_id_movie INT) USE
CALL Upd_Movie_Score('id_movie')