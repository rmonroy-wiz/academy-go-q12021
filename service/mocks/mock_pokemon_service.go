// Code generated by mockery 2.7.4. DO NOT EDIT.

package service

import (
	mock "github.com/stretchr/testify/mock"
	model "github.com/wizelineacademy/academy-go-q12021/model"
)

// MockPokemonService is an autogenerated mock type for the IExternalPokemonAPI type
type MockPokemonService struct {
	mock.Mock
}

// GetPokemonFromAPI provides a mock function with given fields: id
func (_m *MockPokemonService) GetPokemonFromAPI(id int) (*model.PokemonAPI, error) {
	ret := _m.Called(id)

	var r0 *model.PokemonAPI
	if rf, ok := ret.Get(0).(func(int) *model.PokemonAPI); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PokemonAPI)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
