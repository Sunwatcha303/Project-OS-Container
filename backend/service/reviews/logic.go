package reviews

import (
	"reflect"
	"strconv"

	templateError "github.com/Sunwatcha303/Project-OS-Container/error"
)

type ReviewsLogic struct {
	ReviewsRepository *ReviewsRepository
}

func (l *ReviewsLogic) InitReviewLogic() *ReviewsLogic {
	repo := ReviewsRepository{}
	return &ReviewsLogic{
		ReviewsRepository: repo.InitReviewRepository(),
	}
}

func (l *ReviewsLogic) GetAllReviewsLogic() (response *[]ReviewResponse, err error) {
	if response, err = l.ReviewsRepository.GetAllReviews(); err != nil {
		return nil, err
	}
	return
}

func (l *ReviewsLogic) GetAllReviewsbyMovieIdLogic(v_id string) (response *[]ReviewResponse, err error) {
	var id int
	id, err = strconv.Atoi(v_id)
	if err != nil {
		return nil, templateError.BadrequestError
	}
	if response, err = l.ReviewsRepository.GetAllReviewsbyMovieId(id); err != nil {
		return nil, err
	}
	return
}

func (l *ReviewsLogic) AddReviewLogic(request *ReviewRequest) (err error) {
	if request.Name == "" {
		return templateError.BadrequestError
	} else if reflect.TypeOf(request.IdMovie) != reflect.TypeOf(int(0)) {
		return templateError.BadrequestError
	} else if reflect.TypeOf(request.Score) != reflect.TypeOf(float64(0)) || (request.Score < 0 || request.Score > 5) {
		return templateError.BadrequestError
	}
	if err = l.ReviewsRepository.AddReview(*request); err != nil {
		return err
	}
	return
}

func (l *ReviewsLogic) DeleteMoviebyIdLogic(v_id string) (err error) {
	var id int
	id, err = strconv.Atoi(v_id)
	if err != nil {
		return templateError.BadrequestError
	}
	if _, err = l.ReviewsRepository.GetReviewsbyId(id); err != nil {
		return err
	}
	if err = l.ReviewsRepository.DeleteReviewById(id); err != nil {
		return err
	}
	return
}
