from visual.src.io.loader import BencDataLoaderImp
from visual.src.io.saver import save_pdf
from visual.src.graps.prepare import GraphDataPreparerImp
from visual.src.graps.drawer import GraphDrawerImp
from visual.src.service.graph_builder import GraphBuilderService
from visual.src.parse import parse


def main():
    loader = BencDataLoaderImp("data/benchdata", parse)
    preparer = GraphDataPreparerImp(loader)
    drawer = GraphDrawerImp()

    graphs_builder = GraphBuilderService(preparer, drawer)

    graphs = list(graphs_builder.create_graphs())

    save_pdf(graphs, "data/graphs.pdf")


if __name__ == "__main__":
    main()
