package repositories

import "rabbitmq/datamodels"

type MovieRepository interface {
	GetMovieName() string
}

type MovieManager struct {
}

func NewMovieManager() MovieRepository {
	return &MovieManager{}
}

func (m *MovieManager) MovieRepository() string {
	movie := &datamodels.Movie{Name: "慕课网视频"}
	return movie.Name
}
