from typing import Any

def random_string(length: int = ...) -> str: ...
def modularize(file_path: str, suffix: str = ...) -> str: ...
def as_filepath(dotted_path: str) -> str: ...
def removeprefix(string: str, prefix: str) -> str: ...
def removesuffix(string: str, suffix: str) -> str: ...
def add_query_params(url: str, query_params: dict[str | int, str | int | bool]) -> str: ...
def get_controller_name(controller: str | Any) -> str: ...
