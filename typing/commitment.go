package typing

import "time"

type Commitment struct {
	ID          string    `json:"id" dynamobav:"id"`
	Owner       string    `json:"owner" dynamobav:"owner"`
	IsActive    bool      `json:"is_active" dynamobav:"is_active"`
	CIDs        []string  `json:"cids" dynamobav:"cids"`
	Earnings    float64   `json:"earnings" dynamobav:"earnings"`
	Duration    int       `json:"duration" dynamobav:"duration"`
	Created     time.Time `json:"created" dynamobav:"created"`
	Updated     time.Time `json:"updated" dynamobav:"updated"`
	SkippedCIDs []string  `json:"skipped_cids"`
	Size        int64     `json:"size" dynamobav:"size"`
	IsSelected  bool      `json:"is_selected" dynamobav:"is_selected"`
	Status      string    `json:"status" dynamobav:"status"`
	History     []History `json:"history" dynamobav:"history"`
}

type History struct {
	Action      string    `json:"action" dynamobav:"action"`
	Description string    `json:"description" dynamobav:"description"`
	Created     time.Time `json:"created" dynamobav:"created"`
	CIDs        []string  `json:"cids" dynamobav:"cids"`
	Author      string    `json:"author" dynamobav:"author"`
}

type TotalCommitment struct {
	Commitments []Commitment `json:"commitments"`
	TotalSize   int64        `json:"total_size"`
	Count       int          `json:"count"`
}
