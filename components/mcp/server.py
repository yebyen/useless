import asyncio
import os
from datetime import datetime
from mcp.server.fastmcp import FastMCP
from kubernetes import client, config
from kubernetes.client.rest import ApiException
from wasm_bridge import WasmBridge
from models import UselessMachineState

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

def get_k8s_client():
    try:
        config.load_incluster_config()
    except config.ConfigException:
        try:
            config.load_kube_config()
        except config.ConfigException:
            return None
    return client.CustomObjectsApi()

def fetch_state() -> UselessMachineState:
    api = get_k8s_client()
    if not api:
        # Fallback for local dev without K8s
        return UselessMachineState(daily_count=0, last_pushed="1970-01-01T00:00:00Z")
    
    try:
        obj = api.get_namespaced_custom_object(
            group="mecris.io",
            version="v1alpha1",
            namespace="default",
            plural="uselessmachines",
            name="global"
        )
        status = obj.get("status", {})
        return UselessMachineState(
            daily_count=status.get("dailyCount", 0),
            last_pushed=status.get("lastPushed", "1970-01-01T00:00:00Z")
        )
    except ApiException as e:
        if e.status == 404:
            return UselessMachineState(daily_count=0, last_pushed="1970-01-01T00:00:00Z")
        raise

@mcp.tool()
def push_button() -> str:
    """
    Pushes the button on the Useless Machine.
    This tool fetches the current state from Kubernetes, processes it 
    through the WASM Brain, and patches the CRD to trigger the action.
    """
    api = get_k8s_client()
    if not api:
        return "Error: Kubernetes API not accessible."

    # Fetch latest state
    state = fetch_state()
    
    # We patch the spec.action to "push" to let the Operator handle the increment
    # via the same WASM Brain. This ensures perfect synchronization.
    patch = {
        "spec": {
            "action": "push"
        }
    }
    
    try:
        api.patch_namespaced_custom_object(
            group="mecris.io",
            version="v1alpha1",
            namespace="default",
            plural="uselessmachines",
            name="global",
            body=patch
        )
        return "Push action recorded in Kubernetes. The Operator will process it shortly."
    except ApiException as e:
        return f"Error patching Kubernetes: {e}"

@mcp.tool()
def get_status() -> str:
    """
    Checks the current status of the Useless Machine from Kubernetes.
    Returns the daily count and whether the user needs to be nagged.
    """
    state = fetch_state()
    now = datetime.now().isoformat() + "Z"
    
    # The WASM Brain determines the nag status
    status = get_bridge().get_status(state, now)
    
    return f"Count: {state.daily_count}, Last Pushed: {state.last_pushed}, Nagging: {status.is_nagging}"

if __name__ == "__main__":
    mcp.run()
