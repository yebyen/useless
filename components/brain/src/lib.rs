use extism_pdk::*;
use serde::{Deserialize, Serialize};

mod logic;
use logic::{UselessMachineState, push_button, NagStatus};

#[derive(Serialize, Deserialize)]
struct PushRequest {
    state: UselessMachineState,
    now: String,
}

#[derive(Serialize, Deserialize)]
struct StatusRequest {
    state: UselessMachineState,
    now: String,
}

#[plugin_fn]
pub fn push_button_wasm(input: String) -> FnResult<String> {
    let req: PushRequest = serde_json::from_str(&input)
        .map_err(|e| anyhow::anyhow!("failed to parse PushRequest: {}", e))?;
    
    let next_state = push_button(req.state, req.now);
    
    let output = serde_json::to_string(&next_state)
        .map_err(|e| anyhow::anyhow!("failed to serialize next state: {}", e))?;
    
    Ok(output)
}

#[plugin_fn]
pub fn get_status_wasm(input: String) -> FnResult<String> {
    let req: StatusRequest = serde_json::from_str(&input)
        .map_err(|e| anyhow::anyhow!("failed to parse StatusRequest: {}", e))?;
    
    let is_lazy = logic::is_nagging(&req.state, req.now);
    
    let status = NagStatus {
        is_nagging: is_lazy,
        message: if is_lazy { Some("Time to push the button!".to_string()) } else { None },
    };
    
    let output = serde_json::to_string(&status)
        .map_err(|e| anyhow::anyhow!("failed to serialize status: {}", e))?;
    
    Ok(output)
}
