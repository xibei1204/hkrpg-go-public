package player

import (
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

func (g *GamePlayer) NewActivity() *spb.Activity {
	return &spb.Activity{
		TrialActivity: make([]uint32, 0),
		ActivityLogin: make(map[uint32]uint32),
	}
}

func (g *GamePlayer) GetActivity() *spb.Activity {
	db := g.GetBasicBin()
	if db.Activity == nil {
		db.Activity = g.NewActivity()
	}
	return db.Activity
}

func (g *GamePlayer) GetTrialActivity() []uint32 {
	if g.GetActivity().TrialActivity == nil {
		g.GetActivity().TrialActivity = make([]uint32, 0)
	}
	return g.GetActivity().TrialActivity
}

func (g *GamePlayer) GetLoginActivity() map[uint32]uint32 {
	if g.GetActivity().ActivityLogin == nil {
		g.GetActivity().ActivityLogin = make(map[uint32]uint32)
	}
	return g.GetActivity().ActivityLogin
}
