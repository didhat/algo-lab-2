import dataclasses as dt
from typing import TypedDict
from numbers import Number

from visual.src.models import BenchMark, OperationType, AlgorithmType


AlgoBenc = dict[AlgorithmType, list[BenchMark]]


OpsBench = dict[OperationType, AlgoBenc]


@dt.dataclass
class PreparedDataForGrap:
    data: OpsBench


@dt.dataclass
class LinearGraph:
    x_values: list[Number]
    y_values: list[Number]
    label: str


@dt.dataclass
class PlotDrawData:
    graphs: list[LinearGraph]
    plot_name: str
    x_name: str
    y_name: str
    is_log: bool
