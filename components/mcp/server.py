import asyncio
import os
from datetime import datetime
from mcp.server.fastmcp import FastMCP
from .wasm_bridge import WasmBridge
from .models import UselessMachineState

# Path to the compiled brain.wasm
BRAIN_WASM_PATH = os.environ.get(
    "BRAIN_WASM_PATH", 
    os.path.join(os.path.dirname(__file__), "../brain/target/wasm32-wasip1/release/brain.wasm")
)

mcp = FastMCP("Vind-Box")
bridge = None

def get_bridge():
    global bridge
    if bridge is None:
        bridge = WasmBridge(BRAIN_WASM_PATH)
    return bridge

@mcp.tool()
def push_button(daily_count: int, last_pushed: str) -> str:
    """
    Pushes the button on the Useless Machine.
    Increments count and resets the timestamp.
    """
    state = UselessMachineState(daily_count=daily_count, last_pushed=last_pushed)
    now = datetime.now().isoformat() + "Z"
    
    next_state = get_bridge().push_button(state, now)
    return next_state.model_dump_json()

@mcp.tool()
def get_status(daily_count: int, last_pushed: str) -> str:
    """
    Checks if the user needs to be nagged based on the machine state.
    """
    state = UselessMachineState(daily_count=daily_count, last_pushed=last_pushed)
    now = datetime.now().isoformat() + "Z"
    
    status = get_bridge().get_status(state, now)
    return status.model_dump_json()

if __name__ == "__main__":
    mcp.run()
