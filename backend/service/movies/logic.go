package movies

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
	var id int
	id, err = strconv.Atoi(v_id)
	if err != nil {
		return nil, templateError.BadrequestError
	}
	if response, err = l.MovieRepository.GetMovieById(id); err != nil {
		return nil, err
	}
	return
}

func (l *MovieLogic) GetScoreBymovieidLogic(v_id string) (response *ScoreResponse, err error) {
	var id int
	id, err = strconv.Atoi(v_id)
	if err != nil {
		return nil, templateError.BadrequestError
	}
	if response, err = l.MovieRepository.GetScoreBymovieid(id); err != nil {
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

func (l *MovieLogic) UpdateMoviebyIdLogic(v_id string, requestUpdate MovieUpdateRequest) (err error) {
	var id int
	id, err = strconv.Atoi(v_id)
	if err != nil {
		return templateError.BadrequestError
	}
	if _, err = l.MovieRepository.GetMovieById(id); err != nil {
		return templateError.MovieNotFoundError
	}
	if requestUpdate.MovieName == nil && requestUpdate.Category == nil && requestUpdate.ImageMovie == nil {
		return templateError.BadrequestError
	} else if *requestUpdate.MovieName == "" || *requestUpdate.Category == "" || *requestUpdate.ImageMovie == "" {
		return templateError.BadrequestError
	}
	if err = l.MovieRepository.UpdateMoviebyId(id, requestUpdate); err != nil {
		return err
	}
	return
}

func (l *MovieLogic) DeleteMoviebyIdLogic(v_id string) (err error) {
	var id int
	id, err = strconv.Atoi(v_id)
	if err != nil {
		return templateError.BadrequestError
	}
	if _, err = l.MovieRepository.GetMovieById(id); err != nil {
		return err
	}
	if err = l.MovieRepository.DeleteMoviebyId(id); err != nil {
		return err
	}
	return
}
