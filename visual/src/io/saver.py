from matplotlib.figure import Figure
from matplotlib.backends.backend_pdf import PdfPages


def save_pdf(figures: list[Figure], path: str):
    doc = PdfPages(path)
    for fig in figures:
        fig.savefig(doc, format="pdf")
    doc.close()
