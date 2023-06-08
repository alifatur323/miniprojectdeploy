package account

import (
	"crmservice/entities"
	mock_repositories "crmservice/repositories/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
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
			Password:   "f78243d4936738a47d809c7532f58e6d28f09fdb357cc63bd4f65459bf3ac109",
			IsVerified: "true",
			IsActive:   "true",
		}
		expectedActor := entities.Actor{
			ID:         1,
			Username:   "alifatur",
			Password:   "f78243d4936738a47d809c7532f58e6d28f09fdb357cc63bd4f65459bf3ac109",
			RoleId:     1,
			IsVerified: "true",
			IsActive:   "true",
		}

		// Set up mock behavior
		actorRepoMock.EXPECT().CreateActor(gomock.Eq(&expectedActor)).Return(&expectedActor, nil)

		// Panggil metode CreateActor
		actualActor, err := au.CreateActor(actorParam)

		// Periksa hasil
		assert.NoError(t, err)
		assert.Equal(t, expectedActor, actualActor)

		// Periksa pemanggilan mock repository
		actorRepoMock.EXPECT().CreateActor(gomock.Eq(&expectedActor)).Times(1)

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
		expectedResult := struct{}{}

		// Set up mock behavior
		actorRepoMock.EXPECT().DeleteActor(actorID).Return(expectedResult, nil)

		// Panggil metode DeleteActor
		actualResult, err := au.DeleteActor(actorID)

		// Periksa hasil
		assert.NoError(t, err)
		assert.Equal(t, expectedResult, actualResult)

		// Periksa pemanggilan mock repository
		actorRepoMock.EXPECT().DeleteActor(actorID).Times(1)
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

		// Periksa pemanggilan mock repository
		actorRepoMock.EXPECT().GetActorById(actorID).Times(1)
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
		username := "john.doe"
		expectedResult := entities.Actor{
			ID:       1,
			Username: "alifatur",
			Password: "f78243d4936738a47d809c7532f58e6d28f09fdb357cc63bd4f65459bf3ac109",
			RoleId:   1,
			IsActive: "true",
		}

		// Set up mock behavior
		actorRepoMock.EXPECT().GetActorByUsername(username).Return(expectedResult, nil)

		// Panggil metode GetActorByUsername
		actualResult, err := au.GetActorByUsername(username)

		// Periksa hasil
		assert.NoError(t, err)
		assert.Equal(t, expectedResult, actualResult)

		// Periksa pemanggilan mock repository
		actorRepoMock.EXPECT().GetActorByUsername(username).Times(1)
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
			Username: "alifatur",
			Password: "f78243d4936738a47d809c7532f58e6d28f09fdb357cc63bd4f65459bf3ac109",
			RoleId:   1,
		}
		actorID := uint(1)
		var expectedResult entities.Actor

		// Set up mock behavior
		actorRepoMock.EXPECT().UpdateActor(gomock.Eq(&actorParam)).Return(&expectedResult, nil)

		// Panggil metode UpdateActor
		actualResult, err := au.UpdateActor(actorParam, actorID)

		// Periksa hasil
		assert.NoError(t, err)
		assert.Equal(t, expectedResult, actualResult)

		// Periksa pemanggilan mock repository
		actorRepoMock.EXPECT().UpdateActor(gomock.Eq(&actorParam)).Times(1)
	})
}
