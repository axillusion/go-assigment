package commons

import "github.com/stretchr/testify/mock"

type MockDataService struct {
	mock.Mock
}

func (m *MockDataService) FetchData(language string, customerID string, page int, pageSize int, destination interface{}, errorChannel chan<- error) {
	m.Called(language, customerID, page, pageSize, destination, errorChannel)
}

func (m *MockDataService) Delete(dialogID string, errorChannel chan<- error) {
	m.Called(dialogID, errorChannel)
}

func (m *MockDataService) SaveToDB(customerID string, dialogID string, text string, language string, errorChannel chan<- error) {
	m.Called(customerID, dialogID, text, language, errorChannel)
}

func (m *MockDataService) ModifyDB(dialogID string, errorChannel chan<- error) {
	m.Called(dialogID, errorChannel)
}
