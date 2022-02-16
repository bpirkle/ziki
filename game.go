package main

import "math/rand"

// Game top-level structure for the game
// Not doing much with this yet, but I feel like I eventually may
type Game struct {
	Player Actor
}

func (g *Game) Play() {
	g.Player = *new(Actor)
	g.Player.Name = "Jimmy Wales"
	g.Player.Initiative = 1 + rand.Intn(100)
	g.Player.Morale = 100
	g.Player.Tactic = 1
	g.Player.CurrentLocation = "Phab"

	Output("blue", Messages["welcome"])
	for {
		Output("blue", LocationMap[g.Player.CurrentLocation].Description)
		g.ProcessEvents(LocationMap[g.Player.CurrentLocation].Events)
		if g.Player.Morale <= 0 {
			Output("white", "You have given up hope on your change. Game over!!!")
			return
		}
		Output("blue", "Morale:", g.Player.Morale)
		if len(LocationMap[g.Player.CurrentLocation].Items) > 0 {
			Output("yellow", "You can see:")
			for _, itm := range LocationMap[g.Player.CurrentLocation].Items {
				Outputf("yellow", "\t%s", Items[itm].Name)
			}
		}
		Output("green", "You can go to these places:")
		for _, loc := range LocationMap[g.Player.CurrentLocation].Transitions {
			Outputf("green", "\t%s", loc)
		}
		cmd := UserInputln()
		ProcessCommands(&g.Player, cmd)
	}
}

func (g *Game) ProcessEvents(events []string) {
	for _, evtName := range events {
		g.Player.Morale += Events[evtName].ProcessEvent(&g.Player)
	}
}
