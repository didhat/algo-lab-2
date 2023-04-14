import pathlib
from unittest import TestCase

from visual.src.io.loader import BencDataLoaderImp
from visual.src.parse import parse
from visual.src.models import BenchMark, OperationType


class TestLoader(TestCase):
    def setUp(self) -> None:
        self.filename = "test_bench"
        file = pathlib.Path(self.filename)

        with file.open(mode="w") as f:
            f.write(
                '{"Time":"2023-04-09T21:19:43.64938117+03:00","Action":"output","Package":"lab2/src/algo",'
                '"Output":"BenchmarkAllAlgoPreparing/treeAlgo:Prepare:10000-12 22 51704417 ns/op"}'
            )

    def tearDown(self) -> None:
        file = pathlib.Path(self.filename)
        file.unlink()

    def test_loader(self):
        loader = BencDataLoaderImp(self.filename, parse)

        for b in loader.load():
            self.assertEqual(b.operation, OperationType.prepare)
