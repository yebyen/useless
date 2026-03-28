from pydantic import BaseModel, Field
from typing import Optional

class UselessMachineState(BaseModel):
    daily_count: int = Field(..., description="Total number of button pushes today")
    last_pushed: str = Field(..., description="ISO8601 timestamp of last push")

class NagStatus(BaseModel):
    is_nagging: bool = Field(..., description="Whether nagging is required")
    message: Optional[str] = None

class PushRequest(BaseModel):
    state: UselessMachineState
    now: str

class StatusRequest(BaseModel):
    state: UselessMachineState
    now: str
