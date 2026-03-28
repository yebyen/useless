package wasm

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	extism "github.com/extism/go-sdk"
)

type UselessMachineState struct {
	DailyCount int32  `json:"daily_count"`
	LastPushed string `json:"last_pushed"`
}

type NagStatus struct {
	IsNagging bool    `json:"is_nagging"`
	Message   *string `json:"message"`
}

type PushRequest struct {
	State UselessMachineState `json:"state"`
	Now   string              `json:"now"`
}

type Runner struct {
	WasmPath string
}

func (r *Runner) call(ctx context.Context, funcName string, req interface{}) ([]byte, error) {
	if _, err := os.Stat(r.WasmPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("WASM file not found at %s", r.WasmPath)
	}

	manifest := extism.Manifest{
		Wasm: []extism.Wasm{
			extism.WasmFile{
				Path: r.WasmPath,
			},
		},
	}

	config := extism.PluginConfig{
		EnableWasi: true,
	}

	plugin, err := extism.NewPlugin(ctx, manifest, config, []extism.HostFunction{})
	if err != nil {
		return nil, fmt.Errorf("failed to create plugin: %w", err)
	}
	defer plugin.Close(ctx)

	input, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	exitCode, output, err := plugin.Call(funcName, input)
	if err != nil {
		return nil, fmt.Errorf("failed to call %s: %w", funcName, err)
	}

	if exitCode != 0 {
		return nil, fmt.Errorf("plugin exited with non-zero code: %d", exitCode)
	}

	return output, nil
}

func (r *Runner) PushButton(ctx context.Context, state UselessMachineState, now string) (UselessMachineState, error) {
	req := PushRequest{State: state, Now: now}
	output, err := r.call(ctx, "push_button_wasm", req)
	if err != nil {
		return state, err
	}

	var nextState UselessMachineState
	if err := json.Unmarshal(output, &nextState); err != nil {
		return state, err
	}
	return nextState, nil
}

func (r *Runner) GetStatus(ctx context.Context, state UselessMachineState, now string) (NagStatus, error) {
	req := PushRequest{State: state, Now: now} // Uses same shape for now
	output, err := r.call(ctx, "get_status_wasm", req)
	if err != nil {
		return NagStatus{}, err
	}

	var status NagStatus
	if err := json.Unmarshal(output, &status); err != nil {
		return NagStatus{}, err
	}
	return status, nil
}
