import abc
from numbers import Number
from typing import Protocol

from matplotlib.figure import Figure

from visual.src.graps.models import PreparedDataForGrap, AlgoBenc, PlotDrawData


class GraphDataPreparer(Protocol):
    def get_data_for_graphs(self) -> PreparedDataForGrap:
        pass


class GraphDrawer(Protocol):
    @abc.abstractmethod
    def draw_common(self, data: PlotDrawData) -> Figure:
        pass

    @abc.abstractmethod
    def draw_with_y_log(self, data: PlotDrawData) -> Figure:
        pass

    @abc.abstractmethod
    def draw_with_x_log(self, data: PlotDrawData) -> Figure:
        pass
