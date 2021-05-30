package database

import (
	"context"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/go-redis/redismock/v8"
)

var ctx = context.TODO()
var client RedisDatabase

// Init initializes mocked Redis client and sets it as global client used in tests
func Init() redismock.ClientMock {
	var mock redismock.ClientMock
	mock, client = InitializeMockDB()

	return mock
}

// TestGet_Success tests successful finding and returning of entry
func TestGet_Success(t *testing.T) {
	mock := Init()

	key := "halksjdhasd"
	value := "https://facebook.com"

	mock.ExpectGet(key).SetVal(value)

	received, err := client.Get(ctx, key)
	assert.NoError(t, err)
	assert.Equal(t, received, value)
}

// TestGet_NoResult tests not finding a result for a sent key
func TestGet_NoResult(t *testing.T) {
	mock := Init()

	key := "halksjdhasd"

	// expect value to not be found
	mock.ExpectGet(key).RedisNil()

	received, err := client.Get(ctx, key)
	assert.EqualError(t, err, "get operation failed: redis: nil")
	assert.Equal(t, "", received)
}

// TestGet_DBFailure tests behavior when database fails
func TestGet_DBFailure(t *testing.T) {
	mock := Init()

	key := "halksjdhasd"

	// expect value to not be found
	mock.ExpectGet(key).SetErr(errors.New("forced error"))

	received, err := client.Get(ctx, key)
	assert.EqualError(t, err, "Fatal database failure: forced error")
	assert.Equal(t, "", received)
}

// TestSet_Success tests successful setting and returning of entry
func TestSet_Success(t *testing.T) {
	mock := Init()

	key := "halksjdhasd"
	value := "https://facebook.com"

	mock.ExpectSet(key, value, 0).SetVal(key)

	received, err := client.Set(ctx, key, value)
	assert.NoError(t, err)
	assert.Equal(t, received, key)
}

// TestSet_NoResult tests error returned by Redis that is not fatal (redis.Nil type of error)
func TestSet_NoResult(t *testing.T) {
	mock := Init()

	key := "halksjdhasd"
	value := "https://facebook.com"

	// expect value to not be found
	mock.ExpectSet(key, value, 0).RedisNil()

	received, err := client.Set(ctx, key, value)
	assert.EqualError(t, err, "set operation failed: redis: nil")
	assert.Equal(t, "", received)
}

// TestSet_DBFailure tests error returned by Redis that is fatal
func TestSet_DBFailure(t *testing.T) {
	mock := Init()

	key := "halksjdhasd"
	value := "https://facebook.com"

	// expect value to not be found
	mock.ExpectSet(key, value, 0).SetErr(errors.New("forced error"))

	received, err := client.Set(ctx, key, value)
	assert.EqualError(t, err, "Fatal database failure: forced error")
	assert.Equal(t, "", received)
}

// TestDelete_Success tests successful deleting and returning of entry
func TestDelete_Success(t *testing.T) {
	mock := Init()

	key := "halksjdhasd"

	mock.ExpectDel(key).SetVal(0)

	received, err := client.Delete(ctx, key)
	assert.NoError(t, err)
	assert.Equal(t, key, received)
}

// TestDelete_NoResult tests not finding a result for a sent key
func TestDelete_NoResult(t *testing.T) {
	mock := Init()

	key := "halksjdhasd"

	// expect value to not be found
	mock.ExpectDel(key).RedisNil()

	received, err := client.Delete(ctx, key)
	assert.EqualError(t, err, "delete operation failed: redis: nil")
	assert.Equal(t, "", received)
}

// TestDelete_DBFailure tests behavior when database fails
func TestDelete_DBFailure(t *testing.T) {
	mock := Init()

	key := "halksjdhasd"

	// expect value to not be found
	mock.ExpectDel(key).SetErr(errors.New("forced error"))

	received, err := client.Delete(ctx, key)
	assert.EqualError(t, err, "Fatal database failure: forced error")
	assert.Equal(t, "", received)
}

// TestExpire_Success tests successful setting expiry and returning of key
func TestExpire_Success(t *testing.T) {
	mock := Init()

	key := "halksjdhasd"
	expiry := time.Minute * 60

	mock.ExpectExpire(key, expiry).SetVal(true)

	received, err := client.Expire(ctx, key, expiry)
	assert.NoError(t, err)
	assert.Equal(t, received, key)
}

// TestExpire_NoResult tests error returned by Redis that is not fatal (redis.Nil type of error)
func TestExpire_NoResult(t *testing.T) {
	mock := Init()

	key := "halksjdhasd"
	expiry := time.Minute * 60

	mock.ExpectExpire(key, expiry).RedisNil()

	received, err := client.Expire(ctx, key, expiry)
	assert.EqualError(t, err, "expire operation failed: redis: nil")
	assert.Equal(t, "", received)
}

// TestExpire_DBFailure tests error returned by Redis that is fatal
func TestExpire_DBFailure(t *testing.T) {
	mock := Init()

	key := "halksjdhasd"
	expiry := time.Minute * 60

	mock.ExpectExpire(key, expiry).SetErr(errors.New("forced error"))

	received, err := client.Expire(ctx, key, expiry)
	assert.EqualError(t, err, "Fatal database failure: forced error")
	assert.Equal(t, "", received)
}
