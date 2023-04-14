from typing import Iterable
from unittest import TestCase

from visual.src.graps.prepare import GraphDataPreparerImp
from visual.src.graps.models import PreparedDataForGrap
from visual.src.models import BenchMark, AlgorithmType, OperationType


class FakeBencMarkCollector:
    def __init__(self):
        self.tree_prepare1 = BenchMark(
            operation=OperationType.prepare,
            algorithm=AlgorithmType.tree,
            rec_and_points_number=100,
            exec_time=10,
        )
        self.map_prepare_1 = BenchMark(
            operation=OperationType.prepare,
            algorithm=AlgorithmType.map,
            rec_and_points_number=20,
            exec_time=20,
        )

        self.basic_query_1 = BenchMark(
            operation=OperationType.query,
            algorithm=AlgorithmType.basic,
            rec_and_points_number=10,
            exec_time=20,
        )

    def load(self) -> Iterable[BenchMark]:
        yield self.tree_prepare1
        yield self.map_prepare_1
        yield self.basic_query_1
        yield self.basic_query_1
        yield self.map_prepare_1

    def get_right_grouped_bench(self):
        dd = {
            OperationType.prepare: {
                AlgorithmType.map: [self.map_prepare_1, self.map_prepare_1],
                AlgorithmType.tree: [self.tree_prepare1],
            },
            OperationType.query: {
                AlgorithmType.basic: [self.basic_query_1, self.basic_query_1]
            },
        }
        return PreparedDataForGrap(data=dd)


class TestGraphDataPreparer(TestCase):
    def test_preparer(self):
        loader = FakeBencMarkCollector()
        preparer = GraphDataPreparerImp(loader)

        data = preparer.get_data_for_graphs()

        right_data = loader.get_right_grouped_bench()

        self.assertEqual(right_data, data)
