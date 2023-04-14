from unittest import TestCase

from visual.src.parse import parse
from visual.src.models import BencDataDict, BencDataAction, OperationType, AlgorithmType


class TestParse(TestCase):
    def test_parse(self):
        bench_data: BencDataDict = {
            "Output": "BenchmarkAllAlgoPreparing/treeAlgo:Prepare:100-12         \t    3662\t    292773 ns/op\n",
            "Action": BencDataAction.output,
        }

        bench = parse(bench_data["Output"])

        self.assertEqual(bench.operation, OperationType.prepare)
        self.assertEqual(bench.algorithm, AlgorithmType.tree)
        self.assertEqual(bench.exec_time, 292773)
        self.assertEqual(
            bench.rec_and_points_number,
            100,
        )

    def test_parse_exec_with_float(self):
        bench_data: BencDataDict = {
            "Output": "BenchmarkAllAlgoQueryPoint/basicAlgo:Query:10-12         \t 7619084\t       166.2 ns/op\n",
            "Action": BencDataAction.output,
        }

        bench = parse(bench_data["Output"])

        self.assertEqual(bench.exec_time, 166.2)
