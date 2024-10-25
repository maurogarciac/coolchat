import os
from logging import Logger, getLogger

from dotenv import load_dotenv, find_dotenv

load_dotenv(find_dotenv(".env"))
lg: Logger = getLogger(__name__)


class Urls:
    # UI
    LANDING_PAGE: str = os.environ.get("WEB_URL")
    SIGN_IN_PAGE: str = LANDING_PAGE + "/login"
    HOME_PAGE: str = LANDING_PAGE + "/home"
    CHAT_PAGE: str = LANDING_PAGE + "/chat"
    SIGN_OUT: str = LANDING_PAGE + "/logout"

    # API
    API: str = os.environ.get("API_URL")
    API_HEALTH: str = API + "/health"
    WEBSOCKET: str = API + "/ws"
    GET_MESSAGES: str = API + "/messages"
    POST_AUTHENTICATE: str = API + "/auth"
    POST_REFRESH: str = API + "/refresh"


class Users:
    USER_ALICE: tuple[str, str] = (
        f"{os.environ.get('USER_ALICE')}",
        f"{os.environ.get('USER_PASSWORD')}",
    )
    USER_BOB: tuple[str, str] = (
        f"{os.environ.get('USER_BOB')}",
        f"{os.environ.get('USER_PASSWORD')}",
    )
