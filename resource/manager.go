// resource/manager.go
package resource

import (
	"bytes"
	"image"
	_ "image/png" // Enable PNG support
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

// Manager handles loading and caching of game resources
type Manager struct {
	images     map[string]*ebiten.Image
	imageMutex sync.RWMutex
}

// NewManager creates a new resource manager
func NewManager() *Manager {
	return &Manager{
		images: make(map[string]*ebiten.Image),
	}
}

// LoadImage loads an image from the resources directory and caches it
func (m *Manager) LoadImage(name string) *ebiten.Image {
	m.imageMutex.RLock()
	if img, exists := m.images[name]; exists {
		m.imageMutex.RUnlock()
		return img
	}
	m.imageMutex.RUnlock()

	// Load the image
	m.imageMutex.Lock()
	defer m.imageMutex.Unlock()

	// Check again in case another goroutine loaded it
	if img, exists := m.images[name]; exists {
		return img
	}

	// Construct the full path
	path := filepath.Join("resources", "images", name)

	// Read the file
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to load image %s: %v", name, err)
	}

	// Decode the image
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		log.Fatalf("failed to decode image %s: %v", name, err)
	}

	// Convert to Ebiten image
	ebitenImg := ebiten.NewImageFromImage(img)
	m.images[name] = ebitenImg

	return ebitenImg
}

// MustLoadImage loads multiple images at once, using their filenames as keys
func (m *Manager) MustLoadImages(names ...string) {
	for _, name := range names {
		m.LoadImage(name)
	}
}

// GetImage retrieves a cached image
func (m *Manager) GetImage(name string) *ebiten.Image {
	m.imageMutex.RLock()
	defer m.imageMutex.RUnlock()

	img, exists := m.images[name]
	if !exists {
		log.Fatalf("image %s not loaded", name)
	}
	return img
}

// Clear removes all cached resources
func (m *Manager) Clear() {
	m.imageMutex.Lock()
	defer m.imageMutex.Unlock()

	m.images = make(map[string]*ebiten.Image)
}
