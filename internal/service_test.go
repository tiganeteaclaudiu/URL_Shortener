package service

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/go-redis/redismock/v8"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"jobtome.com/urlshortener/database"
	proto "jobtome.com/urlshortener/proto"
)

var ctx = context.Background()
var client database.RedisDatabase
var mockService UrlShortener

// Init initializes mocked Redis client and sets it as global client used in tests
func Init() redismock.ClientMock {
	mockService = UrlShortener{}

	var mock redismock.ClientMock
	mock, client = database.InitializeMockDB()
	InitializeService(&client)

	return mock
}

// TestGetShortenedUrl_Success tests successful finding and returning of entry
func TestGetShortenedUrl_Success(t *testing.T) {
	mock := Init()

	key := "halksjdhasd"
	value := "https://facebook.com"

	mock.ExpectGet(key).SetVal(value)

	received, err := mockService.GetShortenedUrl(ctx, &proto.Key{Key: key})
	assert.NoError(t, err)
	assert.Equal(t, received, &proto.Url{Url: value})
}

// TestGetShortenedUrl_NoResult tests not finding a result for a sent key
func TestGetShortenedUrl_NoResult(t *testing.T) {
	mock := Init()

	key := "halksjdhasd"

	// expect value to not be found
	mock.ExpectGet(key).RedisNil()
	var expected *proto.Url

	received, err := mockService.GetShortenedUrl(ctx, &proto.Key{Key: key})
	assert.EqualError(t, err, "Error getting entry: get operation failed: redis: nil")
	assert.Equal(t, expected, received)
}

// TestGetShortenedUrl_DBFailure tests behavior when database fails
func TestGetShortenedUrl_DBFailure(t *testing.T) {
	mock := Init()

	key := "halksjdhasd"

	// expect value to not be found
	mock.ExpectGet(key).SetErr(errors.New("forced error"))
	var expected *proto.Url

	received, err := mockService.GetShortenedUrl(ctx, &proto.Key{Key: key})
	assert.EqualError(t, err, "Error getting entry: Fatal database failure: forced error")
	assert.Equal(t, expected, received)
}

// TestSetShortenedUrl_SuccessNoExpiry tests successful creating and returning a new entry that doesn't expire
func TestSetShortenedUrl_SuccessNoExpiry(t *testing.T) {
	mock := Init()
	// seed randomizer with fixed seed
	rand.Seed(1)

	key := "XVlBzgbaiCMRAj"
	value := "https://facebook.com"

	mock.ExpectSet(key, value, 0).SetVal(value)

	received, err := mockService.SetShortenedUrl(ctx, &proto.SetShortenedUrlInput{
		Url:           value,
		ExpiryMinutes: 0,
	})
	assert.NoError(t, err)
	assert.Equal(t, received, &proto.Url{Url: "localhost:8080/XVlBzgbaiCMRAj"})
}

// TestSetShortenedUrl_SuccessExpiry tests successful creating entry and setting an expiry for it
func TestSetShortenedUrl_SuccessExpiry(t *testing.T) {
	mock := Init()
	// seed randomizer with fixed seed
	rand.Seed(1)

	key := "XVlBzgbaiCMRAj"
	value := "https://facebook.com"

	mock.ExpectSet(key, value, 0).SetVal(value)
	// also expect expire redis call as value for ExpiryMinutes is not 0
	mock.ExpectExpire(key, time.Minute*60).SetVal(true)

	received, err := mockService.SetShortenedUrl(ctx, &proto.SetShortenedUrlInput{
		Url:           value,
		ExpiryMinutes: 60,
	})
	assert.NoError(t, err)
	assert.Equal(t, received, &proto.Url{Url: "localhost:8080/XVlBzgbaiCMRAj"})
}

// TestSetShortenedUrl_FailedExpiry tests successful creating an entry, but failing to set an expiry
// EXPIRE call fail should trigger deletion of entry and return of error
func TestSetShortenedUrl_FailedExpiry(t *testing.T) {
	mock := Init()
	// seed randomizer with fixed seed
	rand.Seed(1)

	key := "XVlBzgbaiCMRAj"
	value := "https://facebook.com"

	mock.ExpectSet(key, value, 0).SetVal("localhost:8080/XVlBzgbaiCMRAj")
	// Force EXPIRE call failure
	mock.ExpectExpire(key, time.Minute*60).SetErr(errors.New("forced expiry error"))
	// Expire call failure should trigger DELETE call on key
	mock.ExpectDel(key).SetVal(0)

	var expected *proto.Url

	received, err := mockService.SetShortenedUrl(ctx, &proto.SetShortenedUrlInput{
		Url:           value,
		ExpiryMinutes: 60,
	})
	assert.EqualError(
		t,
		err,
		"Failed to set expiry on entry: Entry has been deleted: Fatal database failure: forced expiry error",
	)
	assert.Equal(t, received, expected)
}

// TestDeleteShortenedUrl_Success tests successful deleting an entry
func TestDeleteShortenedUrl_Success(t *testing.T) {
	mock := Init()
	// seed randomizer with fixed seed
	rand.Seed(1)

	key := "XVlBzgbaiCMRAj"

	mock.ExpectDel(key).SetVal(0)

	received, err := mockService.DeleteShortenedUrl(ctx, &proto.Key{
		Key: key,
	})
	assert.NoError(t, err)
	assert.Equal(t, received, &proto.Void{})
}

// TestDeleteShortenedUrl_NoResult tests deleting an entry when it cannot be found in the db
func TestDeleteShortenedUrl_NoResult(t *testing.T) {
	mock := Init()
	// seed randomizer with fixed seed
	rand.Seed(1)

	key := "XVlBzgbaiCMRAj"

	mock.ExpectDel(key).RedisNil()
	var expected *proto.Void

	received, err := mockService.DeleteShortenedUrl(ctx, &proto.Key{
		Key: key,
	})
	assert.EqualError(t, err, "Error deleting entry: delete operation failed: redis: nil")
	assert.Equal(t, received, expected)
}

// TestDeleteShortenedUrl_DBFailure tests deleting an entry, forcing a DB failure
func TestDeleteShortenedUrl_DBFailure(t *testing.T) {
	mock := Init()
	// seed randomizer with fixed seed
	rand.Seed(1)

	key := "XVlBzgbaiCMRAj"

	mock.ExpectDel(key).SetErr(errors.New("forced error"))
	var expected *proto.Void

	received, err := mockService.DeleteShortenedUrl(ctx, &proto.Key{
		Key: key,
	})
	assert.EqualError(t, err, "Error deleting entry: Fatal database failure: forced error")
	assert.Equal(t, received, expected)
}
