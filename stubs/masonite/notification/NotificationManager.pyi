from ..exceptions.exceptions import NotificationException as NotificationException
from .AnonymousNotifiable import AnonymousNotifiable as AnonymousNotifiable
from _typeshed import Incomplete

class NotificationManager:
    sent_notifications: Incomplete
    dry_notifications: Incomplete
    application: Incomplete
    drivers: Incomplete
    driver_config: Incomplete
    options: Incomplete
    def __init__(self, application, driver_config: Incomplete | None = ...) -> None: ...
    def add_driver(self, name, driver) -> None: ...
    def get_driver(self, name): ...
    def set_configuration(self, config): ...
    def get_config_options(self, driver): ...
    def send(self, notifiables, notification, drivers=..., dry: bool = ..., fail_silently: bool = ...): ...
    def route(self, driver, route): ...
