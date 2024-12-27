package nodes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pixambi/hashicorp-visualised.git/internal/components"
	"github.com/yohamta/donburi"
	dmath "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
)

type NodeManager struct {
	world  donburi.World
	images map[string]*ebiten.Image
}

func NewNodeManager(world donburi.World) *NodeManager {
	return &NodeManager{
		world:  world,
		images: make(map[string]*ebiten.Image),
	}
}

func (nm *NodeManager) CreateNode(id, label string, pos dmath.Vec2, imagePath string) donburi.Entity {
	entity := nm.world.Create(
		components.Node,
		components.Sprite,
		transform.Transform,
	)

	entry := nm.world.Entry(entity)

	components.Node.SetValue(entry, components.NodeData{
		ID:         id,
		Label:      label,
		State:      "default",
		CustomData: make(map[string]interface{}),
	})

	transform.SetWorldPosition(entry, pos)
	transform.SetWorldScale(entry, dmath.Vec2{X: 1, Y: 1})
	transform.SetWorldRotation(entry, 0)

	if img, ok := nm.images[imagePath]; ok {
		components.Sprite.SetValue(entry, components.SpriteData{
			Image: img,
			Scale: 1.0,
		})
	}
	return entity
}

func (nm *NodeManager) CreateChildNode(parentEntity donburi.Entity, id, label string, localPos dmath.Vec2, imagePath string) donburi.Entity {
	childEntity := nm.CreateNode(id, label, dmath.Vec2{}, imagePath)
	childEntry := nm.world.Entry(childEntity)
	parentEntry := nm.world.Entry(parentEntity)

	transform.AppendChild(parentEntry, childEntry, false)

	transformData := transform.GetTransform(childEntry)
	transformData.LocalPosition = localPos

	return childEntity

}

func (nm *NodeManager) UpdateNodeState(entity donburi.Entity, newState string) {
	entry := nm.world.Entry(entity)
	if !entry.HasComponent(components.Node) {
		return
	}
	nodeData := components.Node.Get(entry)
	nodeData.State = newState
	components.Node.Set(entry, nodeData)
}

func (nm *NodeManager) RemoveNode(entity donburi.Entity) {
	entry := nm.world.Entry(entity)
	transform.RemoveRecursive(entry)
}

func (nm *NodeManager) GetNodesByState(state string) []*donburi.Entry {
	var nodes []*donburi.Entry
	query := donburi.NewQuery(filter.Contains(components.Node))
	for entry := range query.Iter(nm.world) {
		nodeData := components.Node.Get(entry)
		if nodeData.State == state {
			nodes = append(nodes, entry)
		}
	}
	return nodes
}

func (nm *NodeManager) UpdateNodeWorldPosition(entity donburi.Entity, newPos dmath.Vec2) {
	entry := nm.world.Entry(entity)
	transform.SetWorldPosition(entry, newPos)
}

func (nm *NodeManager) UpdateNodeLocalPosition(entity donburi.Entity, newLocalPos dmath.Vec2) {
	entry := nm.world.Entry(entity)
	transformData := transform.GetTransform(entry)
	transformData.LocalPosition = newLocalPos
}

func (nm *NodeManager) RotateNode(entity donburi.Entity, angle float64) {
	entry := nm.world.Entry(entity)
	transform.SetWorldRotation(entry, angle)
}

func (nm *NodeManager) LookAt(entity donburi.Entity, target dmath.Vec2) {
	entry := nm.world.Entry(entity)
	transform.LookAt(entry, target)
}

func (nm *NodeManager) GetChildren(entity donburi.Entity) []*donburi.Entry {
	entry := nm.world.Entry(entity)
	children, _ := transform.GetChildren(entry)
	return children
}

func (nm *NodeManager) GetParent(entity donburi.Entity) (*donburi.Entry, bool) {
	entry := nm.world.Entry(entity)
	return transform.GetParent(entry)
}
