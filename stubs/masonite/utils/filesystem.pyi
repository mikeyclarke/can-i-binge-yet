def make_directory(directory: str) -> bool: ...
def file_exists(directory: str) -> bool: ...
def make_full_directory(directory: str) -> bool: ...
def creation_date(path_to_file: str) -> int | float: ...
def modified_date(path_to_file: str) -> int | float: ...
def render_stub_file(stub_file: str, name: str) -> str: ...
def get_module_dir(module_file: str) -> str: ...
def get_extension(filepath: str, without_dot: bool = ...) -> str: ...
