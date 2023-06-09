package customer

import (
	"crmservice/entities"
	mocks "crmservice/repositories/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
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
		customerRepoMock.EXPECT().CreateCustomer(gomock.Any()).Return(&expectedCustomer, nil)

		// Call method CreateCustomer
		actualCustomer, err := au.CreateCustomer(customerParam)
		actualCustomer.Id = 1
		actualCustomer.Created_at = time.Now()
		actualCustomer.Updated_at = time.Now()
		expectedCustomer.Created_at = time.Now()
		expectedCustomer.Updated_at = time.Now()

		// Check result
		assert.NoError(t, err)
		assert.Equal(t, expectedCustomer, actualCustomer)
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
		var expectedResult any
		expectedResult = nil

		// Set up mock behavior
		customerRepoMock.EXPECT().DeleteCustomer(customerID).Return(nil, nil)

		// Call method DeleteCustomer
		actualResult, err := au.DeleteCustomer(customerID)

		// Check result
		assert.NoError(t, err)
		assert.Equal(t, expectedResult, actualResult)
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
			Created_at: time.Now(),
			Updated_at: time.Now(),
		}

		// Set up mock behavior
		customerRepoMock.EXPECT().GetCustomerById(customerID).Return(expectedResult, nil)
		actualResult, err := au.GetCustomerById(customerID)
		assert.NoError(t, err)
		assert.Equal(t, expectedResult, actualResult)
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
			Created_at: time.Now(),
			Updated_at: time.Now(),
		}
		customerID := uint(1)
		var expectedResult entities.Customer
		expectedResult.Id = customerID
		expectedResult.First_name = customerParam.First_name
		expectedResult.Last_name = customerParam.Last_name
		expectedResult.Email = customerParam.Email
		expectedResult.Avatar = customerParam.Avatar
		expectedResult.Created_at = customerParam.Created_at
		expectedResult.Updated_at = customerParam.Updated_at

		// Set up mock behavior
		customerRepoMock.EXPECT().UpdateCustomer(gomock.Any()).Return(&expectedResult, nil)
		actualResult, err := au.UpdateCustomer(customerParam, customerID)
		assert.NoError(t, err)
		assert.Equal(t, expectedResult, actualResult)
	})
}
