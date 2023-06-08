package customer

import (
	"crmservice/entities"
	mocks "crmservice/repositories/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_useCaseCustomer_CreateCustomer(t *testing.T) {
	// Initiate mock repository
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	customerRepoMock := mocks.NewMockCustomerInterfaceRepo(ctrl)

	// Initiate use case
	au := useCaseCustomer{
		customerRepo: customerRepoMock,
	}

	// Case 1: Testing create customer success
	t.Run("Create customer success", func(t *testing.T) {
		// Initiate input dan output yang
		customerParam := CustomerParam{
			First_name: "Michael Jackson",
			Last_name:  "Lawson",
			Email:      "michael.lawson@reqres.in",
			Avatar:     "https://reqres.in/img/faces/7-image.jpg",
		}
		expectedCustomer := entities.Customer{
			Id:         1,
			First_name: "Michael Jackson",
			Last_name:  "Lawson",
			Email:      "michael.lawson@reqres.in",
			Avatar:     "https://reqres.in/img/faces/7-image.jpg",
		}

		// Set up mock behavior
		customerRepoMock.EXPECT().CreateCustomer(gomock.Eq(&expectedCustomer)).Return(&expectedCustomer, nil)

		// Call method CreateCustomer
		actualCustomer, err := au.CreateCustomer(customerParam)

		// Check result
		assert.NoError(t, err)
		assert.Equal(t, expectedCustomer, actualCustomer)

		// Check call mock repository
		customerRepoMock.EXPECT().CreateCustomer(gomock.Eq(&expectedCustomer)).Times(1)
	})

	// Case 2: Testing create customer failure
	t.Run("Create customer failure", func(t *testing.T) {
		// Initiate input yang menghasilkan kondisi salah
		customerParam := CustomerParam{
			First_name: "", // Contoh: First_name kosong
			Last_name:  "Lawson",
			Email:      "michael.lawson@reqres.in",
			Avatar:     "https://reqres.in/img/faces/7-image.jpg",
		}

		// Call method CreateCustomer
		actualCustomer, err := au.CreateCustomer(customerParam)

		// Check result
		assert.Error(t, err)
		assert.Equal(t, entities.Customer{}, actualCustomer)

		// Check call mock repository (tidak boleh dipanggil)
		customerRepoMock.EXPECT().CreateCustomer(gomock.Any()).Times(0)
	})
}

func Test_useCaseCustomer_DeleteCustomer(t *testing.T) {
	// Initiate mock repository
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	customerRepoMock := mocks.NewMockCustomerInterfaceRepo(ctrl)

	// Initiate use case
	au := useCaseCustomer{
		customerRepo: customerRepoMock,
	}

	// Case 1: Testing delete customer success
	t.Run("Delete customer success", func(t *testing.T) {
		// Initiate input dan output yang expected
		customerID := uint(1)
		expectedResult := struct{}{}

		// Set up mock behavior
		customerRepoMock.EXPECT().DeleteCustomer(customerID).Return(expectedResult, nil)

		// Call method DeleteCustomer
		actualResult, err := au.DeleteCustomer(customerID)

		// Check result
		assert.NoError(t, err)
		assert.Equal(t, expectedResult, actualResult)

		// Check call mock repository
		customerRepoMock.EXPECT().DeleteCustomer(customerID).Times(1)
	})
}

func Test_useCaseCustomer_GetCustomerById(t *testing.T) {
	// Initiate mock repository
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	customerRepoMock := mocks.NewMockCustomerInterfaceRepo(ctrl)

	// Initiate use case
	au := useCaseCustomer{
		customerRepo: customerRepoMock,
	}

	// Case 1: Testing get customer by ID success
	t.Run("Get customer by ID success", func(t *testing.T) {
		// Initiate input dan output yang di expected
		customerID := uint(1)
		expectedResult := entities.Customer{
			Id:         1,
			First_name: "Michael Jackson",
			Last_name:  "Lawson",
			Email:      "michael.lawson@reqres.in",
			Avatar:     "https://reqres.in/img/faces/7-image.jpg",
		}

		// Set up mock behavior
		customerRepoMock.EXPECT().GetCustomerById(customerID).Return(expectedResult, nil)
		actualResult, err := au.GetCustomerById(customerID)
		assert.NoError(t, err)
		assert.Equal(t, expectedResult, actualResult)
		customerRepoMock.EXPECT().GetCustomerById(customerID).Times(1)
	})
}

func Test_useCaseCustomer_UpdateCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	customerRepoMock := mocks.NewMockCustomerInterfaceRepo(ctrl)
	au := useCaseCustomer{
		customerRepo: customerRepoMock,
	}

	// Case 1: Testing update customer success
	t.Run("Update customer success", func(t *testing.T) {
		customerParam := CustomerParam{
			First_name: "Michael Jackson",
			Last_name:  "Lawson",
			Email:      "michael.lawson@reqres.in",
			Avatar:     "https://reqres.in/img/faces/7-image.jpg",
		}
		customerID := uint(1)
		var expectedResult entities.Customer

		// Set up mock behavior
		customerRepoMock.EXPECT().UpdateCustomer(gomock.Eq(&customerParam)).Return(&expectedResult, nil)
		actualResult, err := au.UpdateCustomer(customerParam, customerID)
		assert.NoError(t, err)
		assert.Equal(t, expectedResult, actualResult)
		customerRepoMock.EXPECT().UpdateCustomer(gomock.Eq(&customerParam)).Times(1)
	})
}
