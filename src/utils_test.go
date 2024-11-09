package cfgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJoinMaps__HappyPath(t *testing.T) {
	m1 := map[string]interface{}{
		"k1": "v1",
		"k2": "v2",
	}

	m2 := map[string]interface{}{
		"k3": "v3",
		"k4": "v4",
	}

	m3 := joinMaps(m1, m2)

	assert.Equal(t, m3["k1"], "v1")
	assert.Equal(t, m3["k2"], "v2")
	assert.Equal(t, m3["k3"], "v3")
	assert.Equal(t, m3["k4"], "v4")
}

func TestJoinMaps__EarlierValuesTakePrecedenceOnConflict(t *testing.T) {
	m1 := map[string]interface{}{
		"k1": "v1",
	}

	m2 := map[string]interface{}{
		"k2": "v2",
	}

	m3 := joinMaps(m1, m2)

	assert.Equal(t, m3["k1"], "v1")
}
