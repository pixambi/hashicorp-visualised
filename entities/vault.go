// entities/vault.go
package entities

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type VaultState string

const (
	Follower  VaultState = "FOLLOWER"
	Candidate VaultState = "CANDIDATE"
	Leader    VaultState = "LEADER"
)

type VaultEntity struct {
	BaseEntity
	size       float32
	voteCount  int
	quorum     int
	totalNodes int
	votes      []*VoteEntity // Track ongoing votes
}

func NewVaultEntity(size float32, totalNodes int) *VaultEntity {
	quorum := (totalNodes + 1) / 2
	return &VaultEntity{
		BaseEntity: BaseEntity{
			position: rl.Vector2{X: 0, Y: 0},
			visible:  true,
			state:    Follower,
		},
		size:       size,
		voteCount:  0,
		quorum:     quorum,
		totalNodes: totalNodes,
		votes:      make([]*VoteEntity, 0),
	}
}

func (v *VaultEntity) Draw() {
	if !v.visible {
		return
	}

	// Calculate offset to center the triangle around its position point
	halfSize := v.size / 2

	// Draw the triangle centered on position
	rl.DrawTriangle(
		rl.Vector2{X: v.position.X + halfSize, Y: v.position.Y - halfSize},
		rl.Vector2{X: v.position.X - halfSize, Y: v.position.Y - halfSize},
		rl.Vector2{X: v.position.X, Y: v.position.Y + halfSize},
		rl.Yellow,
	)

	// Draw outline
	rl.DrawTriangleLines(
		rl.Vector2{X: v.position.X - halfSize, Y: v.position.Y - halfSize},
		rl.Vector2{X: v.position.X + halfSize, Y: v.position.Y - halfSize},
		rl.Vector2{X: v.position.X, Y: v.position.Y + halfSize},
		rl.Black,
	)

	// Draw the label with vote count
	v.DrawLabelWithVotes()
}

func (v *VaultEntity) DrawLabelWithVotes() {
	if v.label != "" {
		fontSize := int32(20)
		stateText := v.label
		voteText := ""

		if v.state == Candidate {
			voteText = fmt.Sprintf(" (%d/%d)", v.voteCount, v.quorum)
		}

		text := stateText + voteText
		textWidth := rl.MeasureText(text, fontSize)
		textX := int32(v.position.X) - textWidth/2
		textY := int32(v.position.Y) - 30 // Offset above the entity
		rl.DrawText(text, textX, textY, fontSize, rl.Black)
	}
}

func (v *VaultEntity) Update() {
	// Update vote positions and check for completed votes
	for i := len(v.votes) - 1; i >= 0; i-- {
		vote := v.votes[i]
		vote.Update()

		// Check if vote has reached its target
		if vote.HasReachedTarget() {
			v.voteCount++
			// Remove the vote from the array
			v.votes = append(v.votes[:i], v.votes[i+1:]...)

			// Check if we've reached quorum
			if v.voteCount >= v.quorum {
				v.state = Leader
				v.label = "Leader"
			}
		}
	}
}

func (v *VaultEntity) SendVoteTo(target *VaultEntity) {
	vote := NewVoteEntity(rl.Gold, 10)
	vote.MoveTo(v.position.X, v.position.Y)
	vote.SetTarget(target.GetPosition())
	v.votes = append(v.votes, vote)
}

// Draw all ongoing votes
func (v *VaultEntity) DrawVotes() {
	for _, vote := range v.votes {
		vote.Draw()
	}
}
