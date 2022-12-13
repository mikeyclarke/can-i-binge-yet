from cleo import Application as CommandApplication, Command # type: ignore

class CommandCapsule:
    command_application: CommandApplication
    commands: list[Command]
    command_name: list[str]
    def __init__(self, command_application: CommandApplication) -> None: ...
    def add(self, *commands: Command) -> CommandCapsule: ...
    def swap(self, command: Command) -> None: ...
    def run(self) -> int: ...
