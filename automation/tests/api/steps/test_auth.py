import pytest
from pytest_bdd import given, when, then, parsers, scenarios
from requests import Response

from automation.tests.api.services import backend

scenarios("../features/auth.feature")


@pytest.fixture
def response() -> list[Response]:
    return []


@given("the backend api is healthy")
def step_given():
    res: Response = backend.get_health()

    assert res.status_code == 200


# Get auth with valid credentials
@when(parsers.parse("I send a POST request to /auth with {username} and {password}"),
      converters={"username": str, "password": str})
def when_valid(response, username: str, password: str):
    res: Response = backend.post_authenticate(username, password)
    response.append(res)


@then('the response status_code is 200')
def then_valid(response):
    res: Response = response[0]

    assert res.status_code == 200


@then('the response body contains an access_token and refresh_token')
def and_valid(response):
    res: Response = response[0]

    body: str = res.text
    assert "access_token" in body
    assert "refresh_token" in body


# Attempt to get auth with invalid request method tests
@when(parsers.parse("I send a {method} request to /auth"),
      converters={"method": str})
def when_invalid_method(response, method: str):
    res: Response = backend.wrong_method_api_call(backend.Urls.POST_AUTHENTICATE, method=method)
    response.append(res)


@then('the response status_code is 405')
def then_invalid_method(response):
    res: Response = response[0]

    assert res.status_code == 405


@then('the response body contains Only POST method allowed')
def and_invalid_method(response):
    res: Response = response[0]

    assert "Only POST method allowed" in res.text


# Attempt to get auth with invalid credentials test
@when(parsers.parse("I send a POST request to /auth with invalid {username} or {password}"),
      converters={"username": str, "password": str})
def when_invalid_credentials(response, username: str, password: str):
    res: Response = backend.post_authenticate(username, password)
    response.append(res)


@then('the response status_code is 403')
def then_invalid_credentials(response):
    res: Response = response[0]

    assert res.status_code == 403


@then('the response body contains User does not exist')
def and_invalid_credentials(response):
    res: Response = response[0]

    assert "User does not exist" in res.text


# Attempt to get auth with missing credentials test
@when(parsers.parse("I send a POST request to /auth with missing {username} or {password}"),
      converters={"username": str, "password": str})
def when_missing_credentials(response, username: str, password: str):
    res: Response = backend.post_authenticate(username, password)
    response.append(res)


@then('the response status_code is 400')
def then_missing_credentials(response):
    res: Response = response[0]

    assert res.status_code == 400

