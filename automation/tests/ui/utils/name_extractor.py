import inspect


def get_fn_name() -> str:
    return inspect.currentframe().f_back.f_code.co_name
