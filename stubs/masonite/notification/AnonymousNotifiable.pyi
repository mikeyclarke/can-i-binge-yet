from .Notifiable import Notifiable as Notifiable
from _typeshed import Incomplete

class AnonymousNotifiable(Notifiable):
    application: Incomplete
    def __init__(self, application: Incomplete | None = ...) -> None: ...
    def route(self, driver, recipient): ...
    def route_notification_for(self, driver): ...
    def send(self, notification, dry: bool = ..., fail_silently: bool = ...): ...
