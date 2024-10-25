import requests
import json
from automation.config.settings import Urls
from requests.models import Response


# Get api health (GET /health)
def get_health() -> Response:
    try:
        response: Response = requests.get(Urls.API_HEALTH)
        return response
    except Exception as e:
        print(f"An error ocurred: {e}")


# Get message history (GET /messages)
def get_messages() -> Response:
    try:
        response: Response = requests.get(Urls.GET_MESSAGES)
        return response
    except Exception as e:
        print(f"An error occurred: {e}")


# Authenticate a user (POST /authenticate)
def post_authenticate(username: str, password: str) -> Response:
    payload: dict[str, str]
    if username == "none":
        payload = {
            "password": password
        }
    elif password == "none":
        payload = {
            "username": username
        }
    else:
        payload = {
            "username": username,
            "password": password
        }

    headers: dict[str, str] = {"Content-Type": "application/json"}
    try:
        response: Response = requests.post(Urls.POST_AUTHENTICATE, data=json.dumps(payload), headers=headers)
        return response
    except Exception as e:
        print(f"An error occurred: {e}")


# Refresh a user's access token (POST /refresh)
def post_refresh(refresh_token: str) -> Response:
    payload: dict[str, str]
    if refresh_token == "none":
        payload = {}
    elif refresh_token == "empty":
        payload = {
            "refresh_token": ""
        }
    else:
        payload = {
            "refresh_token": refresh_token
        }
    headers: dict[str, str] = {"Content-Type": "application/json"}
    try:
        response: Response = requests.post(Urls.POST_REFRESH, data=json.dumps(payload), headers=headers)
        return response
    except Exception as e:
        print(f"An error occurred: {e}")


# Call an api endpoint with the wrong request method
def wrong_method_api_call(endpoint: str, method: str) -> Response:
    try:
        response = Response()
        match method:
            case "GET":
                response: Response = requests.get(endpoint)
            case "POST":
                response: Response = requests.post(endpoint, data=json.dumps({}),
                                                   headers={"Content-Type": "application/json"})
            case "PUT":
                response: Response = requests.put(endpoint, data=json.dumps({}),
                                                  headers={"Content-Type": "application/json"})
            case "PATCH":
                response: Response = requests.patch(endpoint, data=json.dumps({}),
                                                    headers={"Content-Type": "application/json"})
            case "DELETE":
                response: Response = requests.delete(endpoint, data=json.dumps({}),
                                                     headers={"Content-Type": "application/json"})
            case "HEAD":
                response: Response = requests.head(endpoint, data=json.dumps({}),
                                                   headers={"Content-Type": "application/json"})
        return response
    except Exception as e:
        print(f"An error occurred: {e}")
