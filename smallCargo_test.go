package ogame

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallCargoConstructionTime(t *testing.T) {
	sc := newSmallCargo()
	assert.Equal(t, 164, sc.ConstructionTime(1, 7, Facilities{Shipyard: 4}))
	assert.Equal(t, 328, sc.ConstructionTime(2, 7, Facilities{Shipyard: 4}))
}

func TestSmallCargoSpeed(t *testing.T) {
	sc := newSmallCargo()
	assert.Equal(t, 22000, sc.GetSpeed(Researches{CombustionDrive: 10, ImpulseDrive: 6}))
}
