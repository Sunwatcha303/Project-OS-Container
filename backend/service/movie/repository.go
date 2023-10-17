package movie

type MovieRepository struct{}

func (r *MovieRepository) InitMovieRepository() *MovieRepository {
	return &MovieRepository{}
}

func (r *MovieRepository) AddMovie() (err error) {
	return nil
}
