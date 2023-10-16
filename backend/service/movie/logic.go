package movie

type MovieLogic struct{}

func (l *MovieLogic) InitMovieLogic() *MovieLogic {
	return &MovieLogic{}
}

func (l *MovieLogic) AddMovieLogic(request MovieRequest) (err error) {
	return nil
}
