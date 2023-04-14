import enum
import dataclasses as d
from typing import TypedDict


class AlgorithmType(enum.StrEnum):
    basic = "basic"
    map = "map"
    tree = "tree"


class OperationType(enum.StrEnum):
    prepare = "prepare"
    query = "query"


class BencDataAction(enum.StrEnum):
    output = "output"
    _pass = "pass"


nanoseconds = float


@d.dataclass(frozen=True)
class BenchMark:
    algorithm: AlgorithmType
    operation: OperationType
    rec_and_points_number: int
    exec_time: nanoseconds


class BencDataDict(TypedDict):
    Action: BencDataAction
    Output: str
