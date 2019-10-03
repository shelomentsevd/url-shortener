package db

import (
	"os"
	"testing"

	"github.com/go-redis/redis"
	"github.com/stretchr/testify/require"
)

var db Database

func TestMain(m *testing.M) {
	db = NewRedisDatabase(&redis.Options{
		Addr: ":6739",
	})

	os.Exit(m.Run())
}

func TestRedisDatabase_Find(t *testing.T) {
	// tsurl is already in redis database
	// Usual url
	url, err := db.Find("tsurl")
	require.NoError(t, err)
	require.Equal(t, "http://example.com", url)

	// Non-existed url
	url, err = db.Find("turl2")
	require.Equal(t, err, ErrNotFound)
}

func TestRedisDatabase_Insert(t *testing.T) {

}

func TestRedisDatabaseConcurrency(t *testing.T) {

}
