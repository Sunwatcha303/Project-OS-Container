package reviews

import (
	"github.com/Sunwatcha303/Project-OS-Container/config"
	templateError "github.com/Sunwatcha303/Project-OS-Container/error"
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
	if err = db.Table("movie").Select("*").Find(&response).Error; err != nil {
		return nil, err
	}
	return
}
