use anyhow::{Result, Context};
use spin_sdk::http::{IntoResponse, Request, Response, Method};
use spin_sdk::http_component;
use spin_sdk::variables;
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

#[derive(Deserialize, Serialize, Clone, Debug)]
struct K8sResource {
    status: Option<UselessMachineStatus>,
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
    let token = variables::get("k8s_token").unwrap_or_else(|_| "mock-token".to_string());
    let api_url = "https://kubernetes.default.svc/apis/mecris.io/v1alpha1/namespaces/default/uselessmachines/global";

    let req = Request::builder()
        .method(Method::Get)
        .uri(api_url)
        .header("authorization", format!("Bearer {}", token))
        .build();

    // Use Spin's HTTP client to call the K8s API
    let resp: Response = spin_sdk::http::send(req).await?;

    if resp.status().is_success() {
        let body = resp.body();
        let resource: K8sResource = serde_json::from_slice(body)?;
        let status = resource.status.context("Machine has no status yet")?;
        let status_json = serde_json::to_string(&status)?;
        
        Ok(Response::builder()
            .status(200)
            .header("content-type", "application/json")
            .body(status_json)
            .build())
    } else {
        Ok(Response::builder()
            .status(resp.status())
            .body(format!("K8s API error: {:?}", resp.status()))
            .build())
    }
}

async fn push_button() -> Result<Response> {
    let token = variables::get("k8s_token").unwrap_or_else(|_| "mock-token".to_string());
    let api_url = "https://kubernetes.default.svc/apis/mecris.io/v1alpha1/namespaces/default/uselessmachines/global";

    // Patch spec.action to "push"
    let patch = serde_json::json!({
        "spec": {
            "action": "push"
        }
    });

    let req = Request::builder()
        .method(Method::Patch)
        .uri(api_url)
        .header("authorization", format!("Bearer {}", token))
        .header("content-type", "application/merge-patch+json")
        .body(serde_json::to_vec(&patch)?)
        .build();

    let resp: Response = spin_sdk::http::send(req).await?;

    if resp.status().is_success() {
        Ok(Response::builder()
            .status(200)
            .body("Push requested")
            .build())
    } else {
        Ok(Response::builder()
            .status(resp.status())
            .body(format!("K8s Patch error: {:?}", resp.status()))
            .build())
    }
}
