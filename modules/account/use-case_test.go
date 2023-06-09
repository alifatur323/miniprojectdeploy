package account

import (
	"crmservice/entities"
	"crmservice/middleware"
	mock_repositories "crmservice/repositories/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestActorStructUseCase_CreateActor(t *testing.T) {
	// Inisialisasi mock repository
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	actorRepoMock := mock_repositories.NewMockActorInterfaceRepo(ctrl)

	// Inisialisasi use case
	au := ActorStructUseCase{
		actorRepo: actorRepoMock,
	}

	// Kasus 1: Testing create actor berhasil
	t.Run("Create actor success", func(t *testing.T) {
		// Inisialisasi input dan output yang diharapkan
		actorParam := ActorParam{
			Username:   "alifatur",
			Password:   "alif123",
			IsVerified: "false",
			IsActive:   "false",
		}
		expectedActor := entities.Actor{
			ID:         1,
			Username:   "alifatur",
			Password:   middleware.CreateHashPassword("alif123"),
			RoleId:     2,
			IsVerified: "false",
			IsActive:   "false",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		// Set up mock behavior
		actorRepoMock.EXPECT().CreateActor(gomock.Any()).Return(&expectedActor, nil)

		// Panggil metode CreateActor
		actualActor, err := au.CreateActor(actorParam)
		actualActor.ID = 1

		// Periksa hasil
		assert.NoError(t, err)
		assert.Equal(t, expectedActor, actualActor)
	})
}

func TestActorStructUseCase_DeleteActor(t *testing.T) {
	// Inisialisasi mock repository
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	actorRepoMock := mock_repositories.NewMockActorInterfaceRepo(ctrl)

	// Inisialisasi use case
	au := ActorStructUseCase{
		actorRepo: actorRepoMock,
	}

	// Kasus 1: Testing delete actor berhasil
	t.Run("Delete actor success", func(t *testing.T) {
		// Inisialisasi input dan output yang diharapkan
		actorID := uint(1)
		var expectedResult any
		expectedResult = nil

		// Set up mock behavior
		actorRepoMock.EXPECT().DeleteActor(actorID).Return(expectedResult, nil)

		// Panggil metode DeleteActor
		actualResult, err := au.DeleteActor(actorID)

		// Periksa hasil
		assert.NoError(t, err)
		assert.Equal(t, expectedResult, actualResult)
	})
}

func TestActorStructUseCase_GetActorById(t *testing.T) {
	// Inisialisasi mock repository
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	actorRepoMock := mock_repositories.NewMockActorInterfaceRepo(ctrl)

	// Inisialisasi use case
	au := ActorStructUseCase{
		actorRepo: actorRepoMock,
	}

	// Kasus 1: Testing get actor by ID berhasil
	t.Run("Get actor by ID success", func(t *testing.T) {
		// Inisialisasi input dan output yang diharapkan
		actorID := uint(1)
		expectedResult := entities.Actor{
			ID:       1,
			Username: "alifatur",
			Password: "f78243d4936738a47d809c7532f58e6d28f09fdb357cc63bd4f65459bf3ac109",
			RoleId:   1,
			IsActive: "true",
		}

		// Set up mock behavior
		actorRepoMock.EXPECT().GetActorById(actorID).Return(expectedResult, nil)

		// Panggil metode GetActorById
		actualResult, err := au.GetActorById(actorID)

		// Periksa hasil
		assert.NoError(t, err)
		assert.Equal(t, expectedResult, actualResult)
	})
}

func TestActorStructUseCase_GetActorByUsername(t *testing.T) {
	// Inisialisasi mock repository
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	actorRepoMock := mock_repositories.NewMockActorInterfaceRepo(ctrl)

	// Inisialisasi use case
	au := ActorStructUseCase{
		actorRepo: actorRepoMock,
	}

	// Kasus 1: Testing get actor by username berhasil
	t.Run("Get actor by username success", func(t *testing.T) {
		// Inisialisasi input dan output yang diharapkan
		username := "alifatur"
		expectedResult := entities.Actor{
			ID:         1,
			Username:   "alifatur",
			Password:   "f78243d4936738a47d809c7532f58e6d28f09fdb357cc63bd4f65459bf3ac109",
			RoleId:     1,
			IsActive:   "true",
			IsVerified: "true",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		// Set up mock behavior
		actorRepoMock.EXPECT().GetActorByUsername(username).Return(expectedResult, nil)

		// Panggil metode GetActorByUsername
		actualResult, err := au.GetActorByUsername(username)
		actualResult.CreatedAt = time.Now()
		actualResult.UpdatedAt = time.Now()

		// Periksa hasil
		assert.NoError(t, err)
		assert.Equal(t, expectedResult, actualResult)
	})
}

func TestActorStructUseCase_UpdateActor(t *testing.T) {
	// Inisialisasi mock repository
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	actorRepoMock := mock_repositories.NewMockActorInterfaceRepo(ctrl)

	// Inisialisasi use case
	au := ActorStructUseCase{
		actorRepo: actorRepoMock,
	}

	// Kasus 1: Testing update actor berhasil
	t.Run("Update actor success", func(t *testing.T) {
		// Inisialisasi input dan output yang diharapkan
		actorParam := ActorParam{
			Username:   "alifatur",
			Password:   "alif123",
			RoleId:     2,
			IsActive:   "true",
			IsVerified: "true",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}
		actorID := uint(1)
		var expectedResult entities.Actor
		expectedResult.ID = actorID
		expectedResult.Username = actorParam.Username
		expectedResult.Password = middleware.CreateHashPassword(actorParam.Password)
		expectedResult.RoleId = actorParam.RoleId
		expectedResult.IsActive = actorParam.IsActive
		expectedResult.IsVerified = actorParam.IsVerified
		expectedResult.CreatedAt = actorParam.CreatedAt
		expectedResult.UpdatedAt = actorParam.UpdatedAt
		// Set up mock behavior
		actorRepoMock.EXPECT().UpdateActor(gomock.Any()).Return(&expectedResult, nil)

		// Panggil metode UpdateActor
		actualResult, err := au.UpdateActor(actorParam, actorID)
		// Periksa hasil
		assert.NoError(t, err)
		assert.Equal(t, expectedResult, actualResult)
	})
}
