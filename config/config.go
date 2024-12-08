// config/config.go
package config

// GameConfig contains the game's configuration settings
type GameConfig struct {
	// Base design resolution (what the game is designed for)
	DesignWidth  int
	DesignHeight int

	// Current window size
	WindowWidth  int
	WindowHeight int

	// Scale factors
	ScaleX float64
	ScaleY float64
}

var Current *GameConfig

func Init(designWidth, designHeight int) {
	Current = &GameConfig{
		DesignWidth:  designWidth,
		DesignHeight: designHeight,
		WindowWidth:  designWidth, // Initial window size matches design size
		WindowHeight: designHeight,
		ScaleX:       1.0,
		ScaleY:       1.0,
	}
}

// UpdateWindowSize updates the window size and recalculates scaling factors
func (c *GameConfig) UpdateWindowSize(width, height int) {
	c.WindowWidth = width
	c.WindowHeight = height

	// Calculate scale while maintaining aspect ratio
	designAspect := float64(c.DesignWidth) / float64(c.DesignHeight)
	windowAspect := float64(width) / float64(height)

	if windowAspect > designAspect {
		// Window is wider than design - scale based on height
		c.ScaleY = float64(height) / float64(c.DesignHeight)
		c.ScaleX = c.ScaleY
	} else {
		// Window is taller than design - scale based on width
		c.ScaleX = float64(width) / float64(c.DesignWidth)
		c.ScaleY = c.ScaleX
	}
}

// ScalePosition converts design coordinates to screen coordinates
func (c *GameConfig) ScalePosition(x, y float64) (float64, float64) {
	// Center the content in the window
	screenX := (x * c.ScaleX) + (float64(c.WindowWidth)-float64(c.DesignWidth)*c.ScaleX)/2
	screenY := (y * c.ScaleY) + (float64(c.WindowHeight)-float64(c.DesignHeight)*c.ScaleY)/2
	return screenX, screenY
}

// UnscalePosition converts screen coordinates to design coordinates
func (c *GameConfig) UnscalePosition(x, y float64) (float64, float64) {
	return x / c.ScaleX, y / c.ScaleY
}
