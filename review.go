package main

import "math/rand"

func runReview(actors Actors) {
	round := 1
	numAlive := actors.Len()
	action := 0
	for {
		Output("green", "\nCode Review patchset ", round, " begins...")
		outputActors("green", actors)
		for x := 0; x < actors.Len(); x++ {
			if actors[x].Morale <= 0 {
				continue
			}
			if !actors[x].Npc {
				Output("blue", "What Do you want to do?")
				for option := 0; option < len(actors[x].Tactics); option++ {
					Output("blue", "\t", option+1, " - ", Tactics[actors[x].Tactics[option]].Name)
				}
				UserInput(&action)
				action--
			} else {
				action = rand.Intn(len(actors[x].Tactics))
			}
			tgt := selectTarget(actors, x)
			if tgt != -1 {
				var effect, tactic = actors[x].Act(action)
				actors[tgt].Morale = actors[tgt].Morale + effect
				if actors[tgt].Morale <= 0 {
					numAlive--
				}
				Output("green", actors[x].Name+" uses ", tactic.Name, " to affect Morale by ", effect, ".")
			}
		}
		if isReviewEnded(actors) {
			break
		} else {
			round++
		}
	}

	Output("green", "Code Review is over.")
	/*
		for x := 0; x < actors.Len(); x++ {
			if actors[x].Morale > 0 {
				Output("blue", actors[x].Name+" is still working on the change!!!")
			}
		}
	*/
}

func outputActors(color string, actors Actors) {
	for x := 0; x < actors.Len(); x++ {
		actors[x].Output(color)
	}
}

// This is a little silly, because npcs can only ever target the player,
// and there are currently only ever two participants in a code review.
// But it might get more useful if we expand.
func selectTarget(actors []Actor, selectorIndex int) int {
	y := selectorIndex
	for {
		y = y + 1
		if y >= len(actors) {
			y = 0
		}
		if (actors[y].Npc != actors[selectorIndex].Npc) && actors[y].Morale > 0 {
			return y
		}
		if y == selectorIndex {
			return -1
		}
	}
	return -1
}

func isReviewEnded(actors []Actor) bool {
	count := make([]int, 2)
	count[0] = 0
	count[1] = 0
	for _, pla := range actors {
		if pla.Morale > 0 {
			if pla.Npc == false {
				count[0]++
			} else {
				count[1]++
			}
		}
	}
	if count[0] == 0 || count[1] == 0 {
		return true
	} else {
		return false
	}
}
