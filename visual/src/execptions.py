class NotFoundBenchmarkDataFile(Exception):
    def __int__(self, error_path: str):
        self.error_path = error_path
