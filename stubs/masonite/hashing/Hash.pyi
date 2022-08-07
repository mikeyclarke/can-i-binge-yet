from ..foundation.Application import Application
from typing import Any, Optional, Protocol, TypeVar


HD_T = TypeVar('HD_T', bound=HashDriver)

class HashDriver(Protocol):
    options: dict[str, Any]
    def __init__(self, options: dict[str, Any] =...) -> None: ...
    def set_options(self, options: dict[str, Any]) -> HD_T: ...
    def make(self, string: str) -> str: ...
    def make_bytes(self, string: str) -> bytes: ...
    def check(self, plain_string: str, hashed_string: bytes | str) -> bool: ...
    def needs_rehash(self, hashed_string: str) -> bool: ...


H_T = TypeVar('H_T', bound=Hash)

class Hash:
    application: Application
    drivers: dict[str, HashDriver]
    driver_config: dict[str, Any]
    options: dict[str, Any]
    def __init__(self, application: Application, driver_config: Optional[dict[str, Any]] = ...) -> None: ...
    def add_driver(self, name: str, driver: HashDriver) -> None: ...
    def set_configuration(self, config: dict[str, Any]) -> H_T: ...
    def get_driver(self, name: Optional[str] = ...) -> HashDriver: ...
    def get_config_options(self, driver: Optional[str] = ...) -> dict[str, Any]: ...
    def make(self, string: str, options: dict[str, Any] = ..., driver: Optional[str] = ...) -> str: ...
    def make_bytes(self, string: str, options: dict[str, Any] = ..., driver: Optional[str] = ...) -> bytes: ...
    def check(self, plain_string: str, hashed_string: bytes | str, options: dict[str, Any] = ..., driver: Optional[str] = ...) -> bool: ...
    def needs_rehash(self, hashed_string: bytes | str, options: dict[str, Any] = ..., driver: Optional[str] = ...) -> bool: ...
