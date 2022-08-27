package object

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	maxDucks = 3
)

type ShootingTargets struct {
	tick  int
	ducks []Object
}

func NewDefaultShootingTargets() Object {
	return &ShootingTargets{
		ducks: []Object{
			NewDefaultDuck(),
		},
	}
}

func (st *ShootingTargets) Update() error {
	st.tick++

	if st.canAddDuck() {
		st.ducks = append(st.ducks, NewDefaultDuck())
	}

	for _, duck := range st.ducks {
		duck.Update()
	}
	return nil
}

func (st *ShootingTargets) Draw(screen *ebiten.Image) error {
	for _, duck := range st.ducks {
		duck.Draw(screen)
	}
	return nil
}

func (st *ShootingTargets) canAddDuck() bool {
	if len(st.ducks) >= maxDucks {
		return false
	}

	rand.Seed(time.Now().Unix())

	return st.tick%60 == 0 && rand.Float64() < 0.3
}
