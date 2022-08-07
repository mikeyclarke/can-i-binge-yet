class TableNotFound:
    def title(self) -> str: ...
    def description(self) -> str: ...
    def regex(self) -> str: ...

class MissingCSRFToken:
    def title(self) -> str: ...
    def description(self) -> str: ...
    def regex(self) -> str: ...

class InvalidCSRFToken:
    def title(self) -> str: ...
    def description(self) -> str: ...
    def regex(self) -> str: ...

class TemplateNotFound:
    def title(self) -> str: ...
    def description(self) -> str: ...
    def regex(self) -> str: ...

class NoneResponse:
    def title(self) -> str: ...
    def description(self) -> str: ...
    def regex(self) -> str: ...

class RouteMiddlewareNotFound:
    def title(self) -> str: ...
    def description(self) -> str: ...
    def regex(self) -> str: ...
