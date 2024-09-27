package mocks

import (
	"github.com/axillusion/go-assigment/database"
	"github.com/stretchr/testify/mock"
)

type MockDataService struct {
	mock.Mock
}

func (m *MockDataService) FetchData(language string, customerID string, page int, pageSize int, destination interface{}, errorChannel chan<- error) {
	m.Called(language, customerID, page, pageSize, destination, errorChannel)
	errorChannel <- nil
}

func (m *MockDataService) Delete(dialogID string, errorChannel chan<- error) {
	m.Called(dialogID, errorChannel)
	errorChannel <- nil
}

func (m *MockDataService) SaveToDB(customerID string, dialogID string, text string, language string, errorChannel chan<- error) {
	m.Called(customerID, dialogID, text, language, errorChannel)
	errorChannel <- nil
}

func (m *MockDataService) ModifyDB(dialogID string, errorChannel chan<- error) {
	m.Called(dialogID, errorChannel)
	errorChannel <- nil
}

// MockDatabase is a mock implementation for the Database interface
type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) Connect() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockDatabase) GetDB() database.Database {
	args := m.Called()
	result := args.Get(0)
	if result == nil {
		return nil
	}
	return result.(database.Database)
}

func (m *MockDatabase) Create(value interface{}) error {
	args := m.Called(value)
	return args.Error(0)
}

func (m *MockDatabase) Table(name string) database.Database {
	m.Called(name)
	return m
}

func (m *MockDatabase) Select(query interface{}, args ...interface{}) database.Database {
	methodArgs := make([]interface{}, len(args)+1)
	methodArgs[0] = query
	copy(methodArgs[1:], args)
	args = m.Called(methodArgs...)
	return m
}

func (m *MockDatabase) Where(query interface{}, args ...interface{}) database.Database {
	methodArgs := make([]interface{}, len(args)+1)
	methodArgs[0] = query
	copy(methodArgs[1:], args)
	args = m.Called(methodArgs...)
	return m
}

func (m *MockDatabase) Order(value interface{}) database.Database {
	m.Called(value)
	return m
}

func (m *MockDatabase) Modify(value interface{}) error {
	args := m.Called(value)
	return args.Error(0)
}

func (m *MockDatabase) Update(column string, value interface{}) error {
	args := m.Called(column, value)
	return args.Error(0)
}

func (m *MockDatabase) Delete(value interface{}, conds ...interface{}) error {
	methodArgs := make([]interface{}, len(conds)+1)
	methodArgs[0] = value
	copy(methodArgs[1:], conds)
	args := m.Called(methodArgs...)
	return args.Error(0)
}

func (m *MockDatabase) Limit(limit int) database.Database {
	m.Called(limit)
	return m
}

func (m *MockDatabase) Offset(offset int) database.Database {
	m.Called(offset)
	return m
}

func (m *MockDatabase) Find(dest interface{}, conds ...interface{}) error {
	methodArgs := make([]interface{}, len(conds)+1)
	methodArgs[0] = dest
	copy(methodArgs[1:], conds)
	args := m.Called(methodArgs...)
	return args.Error(0)
}

func (m *MockDatabase) Close() error {
	args := m.Called()
	return args.Error(0)
}
