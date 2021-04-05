package business

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wizelineacademy/academy-go-q12021/model"
	repository "github.com/wizelineacademy/academy-go-q12021/repository/mocks"
	service "github.com/wizelineacademy/academy-go-q12021/service/mocks"
)

func TestGetAllPokemons(t *testing.T) {
	mockRepo := repository.MockPokemonRepository{}

	testService, err := NewPokemonBusiness(&mockRepo, nil)
	if err != nil {
		panic(err)
	}

	mockRepo.On("GetAll").Return([]model.Pokemon{
		{
			Id:             3,
			Name:           "Bulbasaur",
			Height:         10,
			Weight:         20,
			BaseExperience: 100,
			PrimaryType:    "Plant",
			SecondaryType:  "Poison",
		},
		{
			Id:             10,
			Name:           "Charmander",
			Height:         10,
			Weight:         20,
			BaseExperience: 100,
			PrimaryType:    "Plant",
			SecondaryType:  "Poison",
		},
	}, nil).Once()

	pokemons, _ := testService.GetAll()

	assert.NotNil(t, pokemons)

	mockRepo.On("GetAll").Return(nil, errors.New("There was an error")).Once()

	listPokemons, err := testService.GetAll()

	assert.NotNil(t, err)
	assert.Nil(t, listPokemons)

}

func TestGetByID(t *testing.T) {
	mockRepo := repository.MockPokemonRepository{}

	testBusiness, err := NewPokemonBusiness(&mockRepo, nil)
	if err != nil {
		panic(err)
	}

	mockRepo.On("GetByID", mock.Anything).Return(
		&model.Pokemon{
			Id:             1,
			Name:           "Bulbasaur",
			Height:         10,
			Weight:         20,
			BaseExperience: 100,
			PrimaryType:    "Plant",
			SecondaryType:  "Poison",
		}, nil).Once()

	pokemon, err := testBusiness.GetByID(1)

	assert.NotNil(t, &pokemon)
	assert.Nil(t, err)

	mockRepo.On("GetByID", mock.Anything).Return(nil, nil).Once()

	pokemon2, err := testBusiness.GetByID(1)

	assert.Nil(t, err)
	assert.Nil(t, pokemon2)
}

func TestGetByIDErrorWithoutInfo(t *testing.T) {
	mockRepo := repository.MockPokemonRepository{}

	mockRepo.On("GetByID", mock.Anything).Return(nil, nil).Once()

	testBusiness, err := NewPokemonBusiness(&mockRepo, nil)
	if err != nil {
		panic(err)
	}

	pokemon, err := testBusiness.GetByID(1)

	assert.Nil(t, err)
	assert.Nil(t, pokemon)

}

func TestStoreByID(t *testing.T) {
	mockService := service.MockPokemonService{}

	testBusiness, err := NewPokemonBusiness(nil, &mockService)
	if err != nil {
		panic(err)
	}

	mockService.On("GetPokemonFromAPI", mock.Anything).Return(nil, errors.New("Error")).Once()

	pokemon, err := testBusiness.StoreByID(1)

	assert.Nil(t, pokemon)
	assert.NotNil(t, err)
}

func TestStoreByIDErrorStoreToCSV(t *testing.T) {
	mockService := service.MockPokemonService{}
	mockRepo := repository.MockPokemonRepository{}

	testBusiness, err := NewPokemonBusiness(&mockRepo, &mockService)
	if err != nil {
		panic(err)
	}

	pokemonAPI := model.PokemonAPI{
		Id:             1,
		Name:           "Bulbasaur",
		Height:         10,
		Weight:         20,
		BaseExperience: 100,
		Types: []model.TypeSlot{
			{Type: model.Type{Name: "Plant"}},
			{Type: model.Type{Name: "Poison"}},
		},
	}

	mockService.On("GetPokemonFromAPI", mock.Anything).Return(&pokemonAPI, nil).Once()
	mockRepo.On("StoreToCSV", mock.Anything).Return(nil, errors.New("Error storing data")).Once()

	pokemon2, err := testBusiness.StoreByID(1)

	assert.Nil(t, pokemon2)
	assert.NotNil(t, err)
}
