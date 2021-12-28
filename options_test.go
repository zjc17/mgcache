package mgcache

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestWithCodec(t *testing.T) {
	// Given
	storageOption := StorageOption{}

	// When
	// Then
	assert.Nil(t, storageOption.codec)

	// When
	WithCodec(NewDefaultCodec())(&storageOption)
	// Then
	assert.NotNil(t, storageOption.codec)
}

func TestWithTimeToLive(t *testing.T) {
	// Given
	storageOption := StorageOption{}
	timeToLive := time.Hour

	// When
	// Then
	assert.Equal(t, time.Duration(0), storageOption.timeToLive)

	// When
	WithTimeToLive(timeToLive)(&storageOption)
	// Then
	assert.NotNil(t, storageOption.timeToLive)
	assert.Equal(t, timeToLive, storageOption.timeToLive)
}

func TestWithContextTimeout(t *testing.T) {
	// Given
	storageOption := StorageOption{}
	contextTimeout := time.Hour

	// When
	// Then
	assert.Equal(t, time.Duration(0), storageOption.contextTimeout)

	// When
	WithContextTimeout(contextTimeout)(&storageOption)
	// Then
	assert.NotNil(t, storageOption.contextTimeout)
	assert.Equal(t, contextTimeout, storageOption.contextTimeout)
}
