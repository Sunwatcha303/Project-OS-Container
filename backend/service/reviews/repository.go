package reviews

import (
	"errors"

	"github.com/Sunwatcha303/Project-OS-Container/config"
	templateError "github.com/Sunwatcha303/Project-OS-Container/error"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ReviewsRepository struct{}

func (r *ReviewsRepository) InitReviewRepository() *ReviewsRepository {
	return &ReviewsRepository{}
}

func (r *ReviewsRepository) GetAllReviews() (response *[]ReviewResponse, err error) {
	if config.Database.DB == nil {
		return nil, templateError.DatabaseConnectedError
	}
	db := config.Database.DB
	if err = db.Table("review").Select("*").Find(&response).Order(clause.OrderByColumn{Column: clause.Column{Name: "time"}, Desc: true}).Error; err != nil {
		return nil, err
	}
	return
}

func (r *ReviewsRepository) GetAllReviewsbyMovieId(id int) (response *[]ReviewResponse, err error) {
	if config.Database.DB == nil {
		return nil, templateError.DatabaseConnectedError
	}
	db := config.Database.DB
	if err = db.Table("review").Select("*").Where("id_movie = ?", id).Find(&response).Order(clause.OrderByColumn{Column: clause.Column{Name: "time"}, Desc: true}).Error; err != nil {
		return nil, err
	}
	return
}

func (r *ReviewsRepository) GetReviewsbyId(id int) (response *ReviewResponse, err error) {
	if config.Database.DB == nil {
		return nil, templateError.DatabaseConnectedError
	}
	db := config.Database.DB
	if err = db.Table("review").Select("*").Where("id = ?", id).First(&response).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, templateError.ReviewNotFoundError
		} else {
			return nil, err
		}
	}
	return
}

func (r *ReviewsRepository) AddReview(request ReviewRequest) (err error) {
	if config.Database.DB == nil {
		return templateError.DatabaseConnectedError
	}
	db := config.Database.DB
	err = db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Table("review").Create(&request).Error; err != nil {
			return err
		} else {
			return nil
		}
	})
	err = db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Exec("CALL Upd_Movie_Score(?)", request.IdMovie).Error; err != nil {
			return err
		} else {
			return nil
		}
	})
	return
}

func (r *ReviewsRepository) UpdateReviewbyId(id, id_movie int, request ReviewUpdateRequest) (err error) {
	if config.Database.DB == nil {
		return templateError.DatabaseConnectedError
	}
	db := config.Database.DB
	err = db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Table("review").Where("id = ?", id).Updates(&request).Error; err != nil {
			return err
		} else {
			return nil
		}
	})
	err = db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Exec("CALL Upd_Movie_Score(?)", id_movie).Error; err != nil {
			return err
		} else {
			return nil
		}
	})
	return
}

func (r *ReviewsRepository) DeleteReviewById(id, id_movie int) (err error) {
	if config.Database.DB == nil {
		return templateError.DatabaseConnectedError
	}
	db := config.Database.DB
	err = db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Table("review").Where("id = ?", id).Delete(&ReviewResponse{}).Error; err != nil {
			return err
		} else {
			return nil
		}
	})
	err = db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Exec("CALL Upd_Movie_Score(?)", id_movie).Error; err != nil {
			return err
		} else {
			return nil
		}
	})
	return
}
