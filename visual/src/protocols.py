import abc
from typing import Protocol, Iterable

from visual.src.models import BenchMark


class BenchDataParse(Protocol):
    @staticmethod
    @abc.abstractmethod
    def __call__(data: str) -> BenchMark:
        pass


class BenchMarkCollector(Protocol):
    @abc.abstractmethod
    def collect_benchmarks(self) -> Iterable[BenchMark]:
        pass
