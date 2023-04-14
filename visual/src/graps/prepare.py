from collections import defaultdict

from visual.src.protocols import BenchMarkCollector
from visual.src.io.protocols import BenchDataLoader
from visual.src.graps.models import PreparedDataForGrap, OpsBench, AlgoBenc
from visual.src.models import BenchMark


class GraphDataPreparerImp:
    def __init__(self, loader: BenchDataLoader):
        self._loader = loader

    def get_data_for_graphs(self) -> PreparedDataForGrap:
        per_ops: OpsBench = defaultdict(dict)

        for benc in self._loader.load():
            if not per_ops.get(benc.operation):
                per_ops[benc.operation]: AlgoBenc = defaultdict(list[BenchMark])

            per_ops[benc.operation][benc.algorithm].append(benc)

        return PreparedDataForGrap(data=per_ops)
