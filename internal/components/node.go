package components

import "github.com/yohamta/donburi"

type NodeData struct {
	ID         string
	Label      string
	State      string
	CustomData map[string]interface{}
}

var Node = donburi.NewComponentType[NodeData]()
