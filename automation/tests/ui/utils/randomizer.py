import random
import string


def get_random_string(size: int = 5) -> str:
    """Get a random string of numbers / lower-case letters / upper-case letters.

    :param size: Desired size for the string.

    :returns: String of random characters with length {size}.
    """
    return "".join(
        random.SystemRandom().choice(string.ascii_lowercase + string.digits)
        for _ in range(size)
    )
