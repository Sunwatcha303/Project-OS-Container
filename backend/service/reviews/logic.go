package reviews

type ReviewsLogic struct {
	repo *ReviewsRepository
}

func (l *ReviewsLogic) InitReviewLogic() *ReviewsLogic {
	return &ReviewsLogic{
		repo: l.repo.InitReviewRepository(),
	}
}
