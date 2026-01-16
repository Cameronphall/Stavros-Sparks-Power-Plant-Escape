package main


type Scene int

const (
    SceneMain Scene = iota
    SceneBreaker
    SceneLeftRoom
    SceneComputerRoom
    SceneFinalGame
    SceneWin
    ScenePause
    SceneMenu
)


type Game struct {
    Scene Scene
}

type GameProgress struct {
    BreakerCompleted bool
    LeftCompleted    bool
    RightCompleted   bool
    ComputerRoomCompleted bool
}

func (p GameProgress) Batteries() int {
    count := 0
    if p.BreakerCompleted {
        count++
    }
    if p.LeftCompleted {
        count++
    }
    if p.RightCompleted {
        count++
    }
    return count
}

func NewGame() Game {
    return Game{Scene: SceneMain}
}


func NewProgress() GameProgress {
    return GameProgress{}
}