from visual.src.models import (
    BencDataDict,
    BencDataAction,
    BenchMark,
    OperationType,
    AlgorithmType,
)

nanoseconds = float


def parse(out: str) -> BenchMark:
    operation = _get_benc_operation(out)
    algo = _get_algorithm_type(out)
    recs_and_points_number = _get_recs_and_points_number(out)
    exec_time = _get_time_execution(out)

    return BenchMark(
        algorithm=algo,
        operation=operation,
        rec_and_points_number=recs_and_points_number,
        exec_time=exec_time,
    )


def _get_time_execution(out: str) -> nanoseconds:
    return float(out.split()[-2])


def _get_recs_and_points_number(out: str) -> int:
    return int(out.split()[0].split(":")[-1].split("-")[0])


def _get_benc_operation(out: str) -> OperationType:
    out_parts = out.split(":")
    operation = out_parts[1]

    if operation == "Prepare":
        return OperationType.prepare

    return OperationType.query


def _get_algorithm_type(out: str) -> AlgorithmType:
    out_slash = out.split("/")
    out_algo = out_slash[1].split(":")[0]

    if out_algo == "treeAlgo":
        return AlgorithmType.tree
    elif out_algo == "basicAlgo":
        return AlgorithmType.basic
    else:
        return AlgorithmType.map


def is_valid_benchmark(data: BencDataDict) -> bool:
    return (
        data["Action"] == BencDataAction.output
        and data["Output"].startswith("BenchmarkAllAlgo")
        and "ns/op" in data["Output"]
    )
