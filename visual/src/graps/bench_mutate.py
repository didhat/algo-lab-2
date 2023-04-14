from typing import Final

from visual.src.graps.models import (
    AlgoBenc,
    PlotDrawData,
    LinearGraph,
    OpsBench,
    PreparedDataForGrap,
)
from visual.src.models import AlgorithmType, OperationType

algo_per_label: Final = {
    AlgorithmType.tree: "Дерево отрезков",
    AlgorithmType.map: "Карта",
    AlgorithmType.basic: "Линейный",
}

operation_per_name: Final = {
    OperationType.query: "Время запроса (Query)",
    OperationType.prepare: "Время подготовки для запросов (Prepare)",
}


def algo2label(algo: AlgorithmType) -> str:
    return algo_per_label.get(algo)


def operation2plot_name(operation: OperationType) -> str:
    name = operation_per_name.get(operation)

    return name


def benches2graphs(per_algo_benches: AlgoBenc) -> list[LinearGraph]:
    graphs: list[LinearGraph] = []

    for algo, benches in per_algo_benches.items():
        gr = LinearGraph(
            x_values=[b.rec_and_points_number for b in benches],
            y_values=[b.exec_time for b in benches],
            label=algo2label(algo),
        )
        graphs.append(gr)

    return graphs


def benches2plot(
    bench: AlgoBenc,
    plot_name: str,
    x_name: str,
    y_name: str,
    is_log: bool = False,
) -> PlotDrawData:
    graphs = benches2graphs(bench)

    return PlotDrawData(
        graphs=graphs, plot_name=plot_name, x_name=x_name, y_name=y_name, is_log=is_log
    )


def graph_data_to_plots_for_drawer(
    graph_data: PreparedDataForGrap,
) -> list[PlotDrawData]:
    plots: list[PlotDrawData] = []
    for operation, per_alog in graph_data.data.items():
        plot_name = operation2plot_name(operation)
        plot = benches2plot(
            per_alog,
            plot_name,
            "Количество точек и прямоугольников",
            "Время(нс)",
        )
        plots.append(plot)

    return plots
