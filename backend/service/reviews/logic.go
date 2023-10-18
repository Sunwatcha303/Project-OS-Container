package reviews

type ReviewsLogic struct {
	ReviewsRepository *ReviewsRepository
}

func (l *ReviewsLogic) InitReviewLogic() *ReviewsLogic {
	repo := ReviewsRepository{}
	return &ReviewsLogic{
		ReviewsRepository: repo.InitReviewRepository(),
	}
}

func (l *ReviewsLogic) GetAllReviewsLogic() (response *[]ReviewResponse, err error){
	if response, err = l.ReviewsRepository.GetAllReviews(); err != nil {
		return nil, err
	}
	return
}
