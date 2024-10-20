import os
from logging import Logger, getLogger

from dotenv import load_dotenv, find_dotenv

load_dotenv(find_dotenv(".env"))
lg: Logger = getLogger(__name__)


def execution_environment() -> str:
    """ Set the env where tests are ran in the .env file (default is local) """
    env: str = os.environ.get("ENV")

    if env == "test":
        return "TEST_URL"
    elif env == "dev":
        return "DEV_URL"
    elif env == "prod":
        return "PROD_URL"
    elif env == "local":
        return "LOCAL_URL"
    else:
        lg.error("Invalid environment to execute tests. Use either: local, test, dev or prod.")
    return "TEST_PORTAL_URL"


class Urls:
    # UI
    HOME_PAGE: str = os.environ.get(execution_environment())
    # SIGN_UP_PAGE: str = HOME_PAGE + "/signup/"
    SIGN_IN_PAGE: str = HOME_PAGE + "/login/"
    # REQUEST_PASSWORD_RESET_PAGE: str = HOME_PAGE + "/accounts/password/reset/"
    # PASSWORD_RESET_DONE_PAGE: str = HOME_PAGE + "/accounts/password/reset/done/"
    # SIGN_IN_AFTER_PW_RESET_PAGE: str = SIGN_IN_PAGE + "?next=/accounts/password/change/"
    # TERMS_AND_CONDITIONS_PAGE: str = HOME_PAGE + "/content/terms-and-conditions/"
    # PROFILE_PAGE: str = HOME_PAGE + "/users/profile/"

    # API (placeholder)
    API: str = os.environ.get("API_URL")


class Users:
    TEST_USER: tuple[str, str] = (
        f"{os.environ.get('USER_EMAIL')}",
        f"{os.environ.get('USER_PASSWORD')}",
    )
