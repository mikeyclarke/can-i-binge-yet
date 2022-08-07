from ..foundation.Application import Application
from typing import Any, Callable, Protocol, TypeVar

CD = TypeVar('CD', bound=CacheDriver)

class CacheDriver(Protocol):
    def set_options(self, options: dict[str, Any]) -> None: ...
    def add(self, key: str, value: Any) -> None: ...
    def get(self, key: str, default: Any | None = ..., **options: Any) -> Any: ...
    def put(self, key: str, value: Any, seconds: int | None = ..., **options: Any) -> None: ...
    def has(self, key: str) -> bool: ...
    def increment(self, key: str, amount: int = ...) -> None: ...
    def decrement(self, key: str, amount: int = ...) -> None: ...
    def remember(self, key: str, callable: Callable[[CD], None]) -> None: ...
    def forget(self, key: str) -> None: ...
    def flush(self) -> None: ...
    def get_value(self, value: Any) -> Any: ...

C = TypeVar('C', bound=Cache)

class Cache:
    application: Application
    drivers: dict[str, CacheDriver]
    store_config: dict[str, Any]
    options: dict[str, Any]
    def __init__(self, application: Application, store_config: dict[str, Any] | None = ...) -> None: ...
    def add_driver(self, name: str, driver: CacheDriver) -> None: ...
    def set_configuration(self, config: dict[str, Any]) -> C: ...
    def get_driver(self, name: str | None = ...) -> CacheDriver: ...
    def get_config_options(self, name: str | None = ...) -> Any: ...
    def store(self, name: str = ...) -> CacheDriver: ...
    def add(self, *args: Any, store: str | None = ..., **kwargs: Any) -> None: ...
    def get(self, *args: Any, store: str | None = ..., **kwargs: Any) -> Any: ...
    def put(self, *args: Any, store: str | None = ..., **kwargs: Any) -> None: ...
    def has(self, *args: Any, store: str | None = ..., **kwargs: Any) -> bool: ...
    def forget(self, *args: Any, store: str | None = ..., **kwargs: Any) -> None: ...
    def increment(self, *args: Any, store: str | None = ..., **kwargs: Any) -> None: ...
    def decrement(self, *args: Any, store: str | None = ..., **kwargs: Any) -> None: ...
    def flush(self, *args: Any, store: str | None = ..., **kwargs: Any) -> None: ...
    def remember(self, *args: Any, store: str | None = ..., **kwargs: Any) -> None: ...
