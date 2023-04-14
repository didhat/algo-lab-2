import matplotlib.pyplot as plt
from matplotlib.figure import Figure
from matplotlib.axes import Axes

from visual.src.graps.models import PlotDrawData


class GraphDrawerImp:
    @staticmethod
    def draw_common(plot_data: PlotDrawData) -> Figure:
        figure = plt.figure()
        title_ax: Axes = plt.axes()

        title_ax.set_ylabel(plot_data.y_name)
        title_ax.set_xlabel(plot_data.x_name)
        title_ax.set_title(plot_data.plot_name)
        figure.add_axes(title_ax)

        for index, graph in enumerate(plot_data.graphs):
            plt.plot(graph.x_values, graph.y_values, label=graph.label, figure=figure)
        plt.legend()
        return figure

    def draw_with_y_log(self, data: PlotDrawData):
        figure = self.draw_common(data)
        plt.semilogy(figure=figure)
        return figure

    def draw_with_x_log(self, data: PlotDrawData) -> Figure:
        figure = self.draw_common(data)
        plt.semilogx(figure=figure)
        return figure
