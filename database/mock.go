package database

import "github.com/go-redis/redismock/v8"

// InitializeMockDB initializes mocked Redis client and sets it as global client used in tests
func InitializeMockDB() (redismock.ClientMock, RedisDatabase) {
	db, mock := redismock.NewClientMock()

	return mock, RedisDatabase{db}
}
