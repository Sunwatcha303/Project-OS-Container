package movie

import (
	"strconv"

	templateError "github.com/Sunwatcha303/Project-OS-Container/error"
)

type MovieLogic struct {
	MovieRepository *MovieRepository
}

func (l *MovieLogic) InitMovieLogic() *MovieLogic {
	repo := MovieRepository{}
	return &MovieLogic{
		MovieRepository: repo.InitMovieRepository(),
	}
}

func (l *MovieLogic) GetAllMovieLogic() (response *[]MovieResponse, err error) {
	if response, err = l.MovieRepository.GetAllMovie(); err != nil {
		return nil, err
	}
	return
}

func (l *MovieLogic) GetMovieByIdLogic(v_id string) (response *MovieResponse, err error) {
	id, err := strconv.Atoi(v_id)
	if err != nil {
		return nil, templateError.BadrequestError
	}
	if response, err = l.MovieRepository.GetMovieById(id); err != nil {
		return nil, err
	}
	return
}

func (l *MovieLogic) AddMovieLogic(request MovieRequest) (err error) {
	if request.IdMovie <= 0 || request.MovieName == "" || request.Category == "" || request.ImageMovie == "" {
		return templateError.BadrequestError
	}
	if err = l.MovieRepository.AddMovie(request); err != nil {
		return err
	}
	return
}

func (l *MovieLogic) DeleteMoviebyIdLogic(v_id string) (err error) {
	id, err := strconv.Atoi(v_id)
	var exitsMovie *MovieResponse
	if exitsMovie, err = l.MovieRepository.GetMovieById(id); err != nil {
		return err
	} else if exitsMovie == nil {
		return templateError.MovieNotFoundError
	}
	if err = l.MovieRepository.DeleteMoviebyId(id); err != nil {
		return err
	}
	return
}
