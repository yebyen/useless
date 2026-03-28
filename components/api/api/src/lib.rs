use anyhow::Result;
use spin_sdk::http::{IntoResponse, Request, Response};
use spin_sdk::http_component;
use serde::{Deserialize, Serialize};

#[derive(Deserialize, Serialize, Clone, Debug)]
pub struct UselessMachineStatus {
    #[serde(rename = "dailyCount")]
    pub daily_count: i32,
    #[serde(rename = "lastPushed")]
    pub last_pushed: String,
    #[serde(rename = "isNagging")]
    pub is_nagging: bool,
}

#[http_component]
async fn handle_api(req: Request) -> Result<impl IntoResponse> {
    let path = req.path();
    
    if path.contains("/status") {
        return get_status().await;
    } else if path.contains("/push") {
        return push_button().await;
    }

    Ok(Response::builder()
        .status(404)
        .body("Not Found")
        .build())
}

async fn get_status() -> Result<Response> {
    // In a real Spin app on K8s, we would read /var/run/secrets/kubernetes.io/serviceaccount/token
    // For now, we return a mock or expect the host to provide it via environment
    
    // This is a placeholder for the actual K8s API call logic
    // We will eventually use spin_sdk::http::send to talk to the K8s API
    
    let mock_status = UselessMachineStatus {
        daily_count: 42,
        last_pushed: "2026-03-28T12:00:00Z".to_string(),
        is_nagging: false,
    };
    
    let body = serde_json::to_string(&mock_status)?;

    Ok(Response::builder()
        .status(200)
        .header("content-type", "application/json")
        .body(body)
        .build())
}

async fn push_button() -> Result<Response> {
    // Placeholder for K8s API Patch
    Ok(Response::builder()
        .status(200)
        .body("Push requested (Mock)")
        .build())
}
