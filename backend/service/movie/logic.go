package movie

type MovieLogic struct {
	MovieRepository *MovieRepository
}

func (l *MovieLogic) InitMovieLogic() *MovieLogic {
	repo := MovieRepository{}
	return &MovieLogic{
		MovieRepository: repo.InitMovieRepository(),
	}
}

func (l *MovieLogic) AddMovieLogic(request MovieRequest) (err error) {
	return nil
}
