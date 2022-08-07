from ..exceptions.exceptions import AuthorizationException as AuthorizationException
from typing import TypeVar

T = TypeVar('T', bound=AuthorizationResponse)

class AuthorizationResponse:
    status: int | None
    def __init__(self, allowed: bool, message: str = ..., status: int | None = ...) -> None: ...
    @classmethod
    def allow(cls, message: str = ..., status: int | None = ...) -> T: ...
    @classmethod
    def deny(cls, message: str = ..., status: int | None = ...) -> T: ...
    def allowed(self) -> bool: ...
    def authorize(self) -> T: ...
    def get_response(self) -> tuple[str, int | None]: ...
    def message(self) -> str: ...
