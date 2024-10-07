package redisClient

import (
	"github.com/stretchr/testify/assert"
	"github.com/xorima/slogger"
	"testing"
)

func TestNewClient(t *testing.T) {
	t.Run("it can create a new client", func(t *testing.T) {
		assert.NotNil(t, NewClient(newMockRedisCfg("127.0.0.1:123", 0), slogger.NewDevNullLogger()))
	})
}

func TestClient_Close(t *testing.T) {
	t.Run("it can close a client", func(t *testing.T) {
		c := NewClient(newMockRedisCfg("127.0.0.1:123", 0), slogger.NewDevNullLogger())
		assert.NoError(t, c.Close())
	})
	t.Run("it should error if closing a closed connection", func(t *testing.T) {
		c := NewClient(newMockRedisCfg("127.0.0.1:123", -1), slogger.NewDevNullLogger())
		assert.NoError(t, c.Close())
		assert.Error(t, c.Close())
	})

}
