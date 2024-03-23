package main

import "fmt"

type SafetyPlacer interface {
	placeSafeties()
}

type RockClimber struct {
	rockedClimbed int
	sp            SafetyPlacer
}

func newRockClimber(sp SafetyPlacer) *RockClimber {
	return &RockClimber{sp: sp}
}

type IceSafetyPlacer struct{}

type NOPSafetyPlacer struct{}

func (sp IceSafetyPlacer) placeSafeties() {
	fmt.Println("placing my ICE safeties...")
}

func (sp NOPSafetyPlacer) placeSafeties() {
	fmt.Println("placing NO safeties...")
}

func (rc *RockClimber) climbRock() {
	rc.rockedClimbed++
	if rc.rockedClimbed == 10 {
		rc.sp.placeSafeties()
	}
}

func main() {
	rc := newRockClimber(NOPSafetyPlacer{})
	for i := 0; i < 15; i++ {
		rc.climbRock()
	}
	rc.sp = IceSafetyPlacer{}
	rc.rockedClimbed = 0
	for i := 0; i < 15; i++ {
		rc.climbRock()
	}
}
