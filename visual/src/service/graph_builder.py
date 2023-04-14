from typing import Iterable

from matplotlib.figure import Figure

from visual.src.graps.protocols import GraphDataPreparer, GraphDrawer
from visual.src.graps.bench_mutate import graph_data_to_plots_for_drawer


class GraphBuilderService:
    def __init__(
        self,
        preparer: GraphDataPreparer,
        drawer: GraphDrawer,
    ):
        self._preparer = preparer
        self._drawer = drawer

    def create_graphs(self) -> Iterable[Figure]:
        data_for_graphs = self._preparer.get_data_for_graphs()
        plots = graph_data_to_plots_for_drawer(data_for_graphs)

        for pl in plots:
            yield self._drawer.draw_common(pl)
            yield self._drawer.draw_with_x_log(pl)
            yield self._drawer.draw_with_y_log(pl)
