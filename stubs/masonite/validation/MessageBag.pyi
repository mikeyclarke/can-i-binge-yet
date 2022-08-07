ITEMS = dict[str, list[str]]

class MessageBag:
    items: ITEMS
    def __init__(self, items: ITEMS =...) -> None: ...
    def add(self, error: str, message: str) -> None: ...
    def all(self) -> ITEMS: ...
    def any(self) -> bool: ...
    def has(self, key: str) -> bool: ...
    def empty(self) -> bool: ...
    def first(self, key: str) -> str | None: ...
    def count(self) -> int: ...
    def json(self) -> str: ...
    def amount(self, key: str) -> int: ...
    def get(self, key: str) -> list[str]: ...
    def errors(self) -> list[str]: ...
    def messages(self) -> list[str]: ...
    def reset(self) -> None: ...
    def merge(self, dictionary: ITEMS) -> bool: ...
    def new(self, dictionary: ITEMS) -> MessageBag: ...
    def __len__(self) -> int: ...
    def get_response(self) -> str: ...
    @staticmethod
    def view_helper(errors: ITEMS =...) -> MessageBag: ...
