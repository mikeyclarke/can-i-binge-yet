class HasColoredOutput:
    def success(self, message: str) -> None: ...
    def warning(self, message: str) -> None: ...
    def danger(self, message: str) -> None: ...
    def info(self, message: str) -> None: ...

class AddCommandColors:
    def error(self, text: str) -> None: ...
    def warning(self, text: str) -> None: ...
