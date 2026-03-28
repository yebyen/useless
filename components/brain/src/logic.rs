use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Debug, Clone, PartialEq)]
pub struct UselessMachineState {
    pub daily_count: i32,
    pub last_pushed: String, // ISO8601
}

#[derive(Serialize, Deserialize, Debug, Clone, PartialEq)]
pub struct NagStatus {
    pub is_nagging: bool,
    pub message: Option<String>,
}

pub fn push_button(mut state: UselessMachineState, now: String) -> UselessMachineState {
    state.daily_count += 1;
    state.last_pushed = now;
    state
}

pub fn is_nagging(state: &UselessMachineState, now: String) -> bool {
    // threshold: 24h
    let last = match chrono::DateTime::parse_from_rfc3339(&state.last_pushed) {
        Ok(t) => t,
        Err(_) => return true, // Fail to nagging if invalid
    };
    let current = match chrono::DateTime::parse_from_rfc3339(&now) {
        Ok(t) => t,
        Err(_) => return false,
    };

    let duration = current.signed_duration_since(last);
    duration.num_hours() >= 24
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_push_button() {
        let state = UselessMachineState {
            daily_count: 5,
            last_pushed: "2026-03-28T10:00:00Z".to_string(),
        };
        let now = "2026-03-28T12:00:00Z".to_string();
        let next_state = push_button(state, now.clone());
        assert_eq!(next_state.daily_count, 6);
        assert_eq!(next_state.last_pushed, now);
    }
}
