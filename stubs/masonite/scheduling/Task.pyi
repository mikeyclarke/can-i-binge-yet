from datetime import datetime

class Task:
    run_every: bool
    run_at: bool
    run_every_hour: bool
    run_every_minute: bool
    twice_daily: bool
    run_weekly: bool
    name: str
    def __init__(self) -> None: ...
    def every(self, time: str) -> Task: ...
    def every_minute(self) -> Task: ...
    def every_15_minutes(self) -> Task: ...
    def every_30_minutes(self) -> Task: ...
    def every_45_minutes(self) -> Task: ...
    def hourly(self) -> Task: ...
    def daily(self) -> Task: ...
    def weekly(self) -> Task: ...
    def monthly(self) -> Task: ...
    def at(self, run_time: int) -> Task: ...
    def at_twice(self, run_time: list[int]) -> Task: ...
    def daily_at(self, run_time: int) -> Task: ...
    def handle(self) -> None: ...
    def should_run(self, date: datetime | None = ...) -> bool: ...
