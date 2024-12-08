// config/config.go
package config

// GameConfig contains the game's configuration settings
type GameConfig struct {
	Width  int
	Height int
}

var Current *GameConfig

func Init(width, height int) {
	Current = &GameConfig{
		Width:  width,
		Height: height,
	}
}
