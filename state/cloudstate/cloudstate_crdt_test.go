// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package cloudstate

import (
	"testing"

	"github.com/dapr/components-contrib/state"
	"github.com/stretchr/testify/assert"
)

func TestMetadata(t *testing.T) {
	t.Run("With all required fields", func(t *testing.T) {
		properties := map[string]string{
			"host":       "localhost:8080",
			"serverPort": "9002",
		}
		c := NewCRDT()
		err := c.Init(state.Metadata{
			Properties: properties,
		})
		assert.Nil(t, err)
	})

	t.Run("Missing host", func(t *testing.T) {
		properties := map[string]string{
			"serverPort": "9002",
		}
		c := NewCRDT()
		err := c.Init(state.Metadata{
			Properties: properties,
		})
		assert.NotNil(t, err)
	})

	t.Run("Missing serverPort", func(t *testing.T) {
		properties := map[string]string{
			"host": "localhost:8080",
		}
		c := NewCRDT()
		err := c.Init(state.Metadata{
			Properties: properties,
		})
		assert.NotNil(t, err)
	})
}
