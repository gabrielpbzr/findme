package core_test

import (
	"testing"

	"github.com/gabrielpbzr/findme/core"
	"github.com/stretchr/testify/assert"
)

func TestPositionInitializing(t *testing.T) {
	assert := assert.New(t)
	t.Run("Should initialize a position with given longitude and latitude", func(t *testing.T) {
		longitude := -42.875326
		latitude := -20.757294
		position := core.CreatePosition(longitude, latitude)

		assert.Equal(longitude, position.Longitude, "Position longitude should be the same provided")
		assert.Equal(latitude, position.Latitude, "Position latitude should be the same provided")
		assert.NotNil(position.Timestamp, "Timestamp should correspond to the current UTC Time")
	})
}
