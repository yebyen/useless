import extism
import json
import os
from .models import UselessMachineState, NagStatus, PushRequest, StatusRequest

class WasmBridge:
    def __init__(self, wasm_path: str):
        if not os.path.exists(wasm_path):
            raise FileNotFoundError(f"WASM brain not found at {wasm_path}")
        
        with open(wasm_path, "rb") as f:
            wasm_data = f.read()
            
        self.plugin = extism.Plugin(wasm_data, wasi=True)

    def push_button(self, state: UselessMachineState, now: str) -> UselessMachineState:
        req = PushRequest(state=state, now=now)
        input_data = req.model_dump_json()
        
        output = self.plugin.call("push_button_wasm", input_data)
        result = json.loads(output)
        
        return UselessMachineState(**result)

    def get_status(self, state: UselessMachineState, now: str) -> NagStatus:
        req = StatusRequest(state=state, now=now)
        input_data = req.model_dump_json()
        
        output = self.plugin.call("get_status_wasm", input_data)
        result = json.loads(output)
        
        return NagStatus(**result)
