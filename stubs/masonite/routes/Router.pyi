from .HTTPRoute import HTTPRoute
from typing import Any, Literal, Optional, TypeVar

T = TypeVar('T', bound=Router)

class Router:
    http_methods: list[str]
    routes: list[HTTPRoute]
    def __init__(self, *routes: list[HTTPRoute], module_location: Literal[None] = ...) -> None: ...
    def find(self, path: str, request_method: str, subdomain: Optional[str] = ...) -> Optional[HTTPRoute]: ...
    def matches(self, path: str) -> Optional[HTTPRoute]: ...
    def find_by_name(self, name: str) -> Optional[HTTPRoute]: ...
    def route(self, name: str, parameters: dict[str, Any] = ..., query_params: dict[str, Any] = ...) -> str: ...
    controller_locations: list[str]
    def set_controller_locations(self, location: list[str]) -> T: ...
    def add(self, *routes: list[HTTPRoute]) -> None: ...
    def set(self, *routes: list[HTTPRoute]) -> None: ...
    @classmethod
    def compile_to_url(cls, uncompiled_route: str, params: dict[str, Any] = ..., query_params: dict[str, Any] =...) -> str: ...
