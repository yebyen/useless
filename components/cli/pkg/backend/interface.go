package backend

import (
	"context"
)

type MachineStatus struct {
	DailyCount int32  `json:"dailyCount"`
	LastPushed string `json:"lastPushed"`
	IsNagging  bool   `json:"isNagging"`
}

type Backend interface {
	GetStatus(ctx context.Context) (*MachineStatus, error)
	PushButton(ctx context.Context) (string, error)
}
