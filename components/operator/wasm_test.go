package operator

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	extism "github.com/extism/go-sdk"
	"github.com/stretchr/testify/assert"
)

type UselessMachineState struct {
	DailyCount int32  `json:"daily_count"`
	LastPushed string `json:"last_pushed"`
}

type PushRequest struct {
	State UselessMachineState `json:"state"`
	Now   string              `json:"now"`
}

func TestBrainWasm(t *testing.T) {
	wasmPath := "../brain/target/wasm32-wasip1/release/brain.wasm"
	if _, err := os.Stat(wasmPath); os.IsNotExist(err) {
		t.Skip("WASM file not found, skipping integration test")
	}

	manifest := extism.Manifest{
		Wasm: []extism.Wasm{
			extism.WasmFile{
				Path: wasmPath,
			},
		},
	}

	config := extism.PluginConfig{
		EnableWasi: true,
	}

	plugin, err := extism.NewPlugin(context.Background(), manifest, config, []extism.HostFunction{})
	if err != nil {
		t.Fatalf("Failed to create plugin: %v", err)
	}
	defer plugin.Close(context.Background())

	state := UselessMachineState{
		DailyCount: 10,
		LastPushed: "2026-03-28T10:00:00Z",
	}
	now := "2026-03-28T14:00:00Z"

	req := PushRequest{
		State: state,
		Now:   now,
	}

	input, _ := json.Marshal(req)

	exitCode, output, err := plugin.Call("push_button_wasm", input)
	if err != nil {
		t.Fatalf("Failed to call push_button_wasm: %v", err)
	}

	if exitCode != 0 {
		t.Fatalf("Plugin exited with non-zero code: %d", exitCode)
	}

	var nextState UselessMachineState
	err = json.Unmarshal(output, &nextState)
	if err != nil {
		t.Fatalf("Failed to unmarshal output: %v", err)
	}

	assert.Equal(t, int32(11), nextState.DailyCount)
	assert.Equal(t, now, nextState.LastPushed)
	fmt.Printf("Success! Count incremented to %d\n", nextState.DailyCount)
}

type NagStatus struct {
	IsNagging bool    `json:"is_nagging"`
	Message   *string `json:"message"`
}

func TestNagWasm(t *testing.T) {
	wasmPath := "../brain/target/wasm32-wasip1/release/brain.wasm"
	if _, err := os.Stat(wasmPath); os.IsNotExist(err) {
		t.Skip("WASM file not found, skipping integration test")
	}

	manifest := extism.Manifest{
		Wasm: []extism.Wasm{
			extism.WasmFile{
				Path: wasmPath,
			},
		},
	}

	config := extism.PluginConfig{
		EnableWasi: true,
	}

	plugin, err := extism.NewPlugin(context.Background(), manifest, config, []extism.HostFunction{})
	if err != nil {
		t.Fatalf("Failed to create plugin: %v", err)
	}
	defer plugin.Close(context.Background())

	// Test Case 1: Nagging after 25 hours
	state := UselessMachineState{
		DailyCount: 10,
		LastPushed: "2026-03-27T10:00:00Z",
	}
	now := "2026-03-28T11:00:00Z" // 25 hours later

	req := PushRequest{
		State: state,
		Now:   now,
	}

	input, _ := json.Marshal(req)

	exitCode, output, err := plugin.Call("get_status_wasm", input)
	if err != nil {
		t.Fatalf("Failed to call get_status_wasm: %v", err)
	}

	if exitCode != 0 {
		t.Fatalf("Plugin exited with non-zero code: %d", exitCode)
	}

	var status NagStatus
	err = json.Unmarshal(output, &status)
	if err != nil {
		t.Fatalf("Failed to unmarshal output: %v", err)
	}

	assert.True(t, status.IsNagging)
	assert.NotNil(t, status.Message)
	fmt.Printf("Nag Check 1 (25h): %v, Message: %s\n", status.IsNagging, *status.Message)

	// Test Case 2: Not nagging after 1 hour
	state.LastPushed = "2026-03-28T10:00:00Z"
	now = "2026-03-28T11:00:00Z" // 1 hour later
	req.State = state
	req.Now = now
	input, _ = json.Marshal(req)

	exitCode, output, err = plugin.Call("get_status_wasm", input)
	assert.NoError(t, err)
	json.Unmarshal(output, &status)
	assert.False(t, status.IsNagging)
	fmt.Printf("Nag Check 2 (1h): %v\n", status.IsNagging)
}
