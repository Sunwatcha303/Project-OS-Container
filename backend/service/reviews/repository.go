package reviews

type ReviewsRepository struct{}

func (r *ReviewsRepository) InitReviewRepository() *ReviewsRepository {
	return &ReviewsRepository{}
}
