import os
from logging import Logger, getLogger

from dotenv import load_dotenv, find_dotenv

load_dotenv(find_dotenv(".env"))
lg: Logger = getLogger(__name__)


class Urls:
    # UI
    LANDING_PAGE: str = os.environ.get("WEB_URL")
    LOG_IN_PAGE: str = LANDING_PAGE + "/login"
    HOME_PAGE: str = LANDING_PAGE + "/home"
    CHAT_PAGE: str = LANDING_PAGE + "/chat"
    LOG_OUT_PAGE: str = LANDING_PAGE + "/logout"

    # API
    API: str = os.environ.get("API_URL")
    API_HEALTH: str = API + "/health"
    WEBSOCKET: str = API + "/ws"
    GET_MESSAGES: str = API + "/messages"
    POST_AUTHENTICATE: str = API + "/auth"
    POST_REFRESH: str = API + "/refresh"


class Users:
    Alice: dict[str, str] = {
        "username": f"{os.environ.get('USER_ALICE')}",
        "password": f"{os.environ.get('USER_PASSWORD')}",
    }
    Bob: dict[str, str] = {
        "username": f"{os.environ.get('USER_BOB')}",
        "password": f"{os.environ.get('USER_PASSWORD')}",
    }
