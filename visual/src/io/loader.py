import json
from typing import Iterable
from pathlib import Path

from visual.src.protocols import BenchDataParse
from visual.src.execptions import NotFoundBenchmarkDataFile
from visual.src.models import BencDataDict, BencDataAction, BenchMark


class BencDataLoaderImp:
    def __init__(self, data_path: str, parser: BenchDataParse):
        self._data_file_path = data_path
        self._parse = parser

    def load(self) -> Iterable[BenchMark]:
        file = self._get_file()

        with file.open(mode="r") as f:
            line = f.readline()
            while line:
                raw_bench: BencDataDict = json.loads(line)
                if is_valid_benchmark(raw_bench):
                    yield self._parse(raw_bench["Output"])
                line = f.readline()

    def _get_file(self) -> Path:
        file = Path(self._data_file_path)

        if not file.exists() or not file.is_file():
            raise NotFoundBenchmarkDataFile(self._data_file_path)

        return file


def is_valid_benchmark(data: BencDataDict) -> bool:
    return (
        data["Action"] == BencDataAction.output
        and data["Output"].startswith("BenchmarkAllAlgo")
        and "ns/op" in data["Output"]
    )
