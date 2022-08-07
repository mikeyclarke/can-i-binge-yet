from ..foundation.Application import Application
from .Provider import Provider

class FrameworkProvider(Provider):
    application: Application
    def __init__(self, application: Application) -> None: ...
    def register(self) -> None: ...
    def boot(self) -> None: ...
