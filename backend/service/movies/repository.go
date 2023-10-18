package movies

import (
	"errors"

	"github.com/Sunwatcha303/Project-OS-Container/config"
	templateError "github.com/Sunwatcha303/Project-OS-Container/error"
	"gorm.io/gorm"
)

type MovieRepository struct{}

func (r *MovieRepository) InitMovieRepository() *MovieRepository {
	return &MovieRepository{}
}

func (r *MovieRepository) GetAllMovie() (response *[]MovieResponse, err error) {
	if config.Database.DB == nil {
		return nil, templateError.DatabaseConnectedError
	}
	db := config.Database.DB
	if err = db.Table("movie").Select("*").Find(&response).Error; err != nil {
		return nil, err
	}
	return
}

func (r *MovieRepository) GetMovieById(id int) (response *MovieResponse, err error) {
	if config.Database.DB == nil {
		return nil, templateError.DatabaseConnectedError
	}
	db := config.Database.DB
	if err = db.Table("movie").Select("*").Where("id_movie = ?", id).First(&response).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, templateError.MovieNotFoundError
		} else {
			return nil, err
		}
	}
	return
}

func (r *MovieRepository) GetScoreBymovieid(id int) (response *ScoreResponse, err error) {
	if config.Database.DB == nil {
		return nil, templateError.DatabaseConnectedError
	}
	db := config.Database.DB
	if err = db.Table("movie").Select("movie_score").Where("id_movie = ?", id).First(&response).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, templateError.MovieNotFoundError
		} else {
			return nil, err
		}
	}
	return
}

func (r *MovieRepository) AddMovie(request MovieRequest) (err error) {
	if config.Database.DB == nil {
		return templateError.DatabaseConnectedError
	}
	db := config.Database.DB
	err = db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Table("movie").Create(&request).Error; err != nil {
			return err
		} else {
			return nil
		}
	})
	return
}

func (r *MovieRepository) DeleteMoviebyId(id int) (err error) {
	if config.Database.DB == nil {
		return templateError.DatabaseConnectedError
	}
	db := config.Database.DB
	err = db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Table("movie").Where("id_movie = ?", id).Delete(&MovieResponse{}).Error; err != nil {
			return err
		} else {
			return nil
		}
	})
	return
}
