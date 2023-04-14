import abc
from typing import Protocol, Iterable

from matplotlib.figure import Figure

from visual.src.models import BenchMark


class BenchDataLoader(Protocol):
    @abc.abstractmethod
    def load(self) -> Iterable[BenchMark]:
        pass


class FiguresSaver(Protocol):
    @abc.abstractmethod
    def __call__(self, figures: list[Figure], path: str):
        pass
