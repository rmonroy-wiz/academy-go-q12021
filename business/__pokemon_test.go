package business

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/wizelineacademy/academy-go-q12021/model"
	"github.com/wizelineacademy/academy-go-q12021/repository"
	"github.com/wizelineacademy/academy-go-q12021/service"

	mocksrepo "github.com/wizelineacademy/academy-go-q12021/repository/mocks"
	mocksservice "github.com/wizelineacademy/academy-go-q12021/service/mocks"
)

func TestNewPokemonBusiness(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocksrepo.NewMockIPokemonRepository(mockCtrl)
	mockService := mocksservice.NewMockIExternalPokemonAPI(mockCtrl)

	type args struct {
		repository repository.IPokemonRepository
		service    service.IExternalPokemonAPI
	}
	tests := []struct {
		name    string
		args    args
		want    IPokemonBusiness
		wantErr bool
	}{
		{
			name: "Happy path",
			args: args{
				repository: mockRepo,
				service:    mockService,
			},
			want: &PokemonBusiness{
				pokemonRepository: mockRepo,
				serviceAPI:        mockService,
			},
			wantErr: false,
		},
		{
			name: "service is nil",
			args: args{
				repository: mockRepo,
				service:    nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPokemonBusiness(tt.args.repository, tt.args.service)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPokemonBusiness() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPokemonBusiness() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPokemonBusiness_GetAll(t *testing.T) {
	type fields struct {
		pokemonRepository repository.IPokemonRepository
		serviceAPI        service.IExternalPokemonAPI
	}
	tests := []struct {
		name    string
		fields  fields
		want    []model.Pokemon
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &PokemonBusiness{
				pokemonRepository: tt.fields.pokemonRepository,
				serviceAPI:        tt.fields.serviceAPI,
			}
			got, err := s.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("PokemonBusiness.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PokemonBusiness.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPokemonBusiness_GetByID(t *testing.T) {
	type fields struct {
		pokemonRepository repository.IPokemonRepository
		serviceAPI        service.IExternalPokemonAPI
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Pokemon
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &PokemonBusiness{
				pokemonRepository: tt.fields.pokemonRepository,
				serviceAPI:        tt.fields.serviceAPI,
			}
			got, err := s.GetByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("PokemonBusiness.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PokemonBusiness.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPokemonBusiness_StoreByID(t *testing.T) {
	type fields struct {
		pokemonRepository repository.IPokemonRepository
		serviceAPI        service.IExternalPokemonAPI
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Pokemon
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &PokemonBusiness{
				pokemonRepository: tt.fields.pokemonRepository,
				serviceAPI:        tt.fields.serviceAPI,
			}
			got, err := s.StoreByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("PokemonBusiness.StoreByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PokemonBusiness.StoreByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
