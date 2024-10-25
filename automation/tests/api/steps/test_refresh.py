import json

import pytest
from pytest_bdd import given, when, then, scenarios, parsers
from requests import Response

from automation.tests.api.services import backend

scenarios("../features/refresh.feature")


@pytest.fixture
def response() -> list[Response]:
    return []


# General steps
@given("the backend api is healthy")
def step_given():
    res: Response = backend.get_health()

    assert res.status_code == 200


@then('the response status_code is 400')
def then_response_bad_request(response):
    res: Response = response[0]

    assert res.status_code == 400


# SCENARIO 1: Get new access_token
@given("I have authenticated through POST request to /auth")
def step_given_auth(response):
    res: Response = backend.post_authenticate("alice", "root")
    response.append(res)

    assert res.status_code == 200


@given("my refresh_token is available")
def and_token_available(response):
    res: Response = response[0]

    assert "refresh_token" in res.text


@when("I send a POST request to /refresh with my refresh_token")
def when_post_refresh(response):
    og_res: json = response[0].json()
    token: str = og_res.get("refresh_token")

    res: Response = backend.post_refresh(token)
    response[0] = res


@then('the response status_code is 200')
def then_token_created(response):
    res: Response = response[0]

    assert res.status_code == 200


@then('the response body contains a new access_token')
def and_valid(response):
    res: Response = response[0]

    assert "access_token" in res.text


# SCENARIO 2: Attempt to get new access_token with invalid refresh_token
@when("I send a POST request to /refresh with invalid refresh_token")
def when_no_refresh(response):
    res: Response = backend.post_refresh("super.invalid.token")
    response.append(res)


# (runs after then_response_bad_request)
@then('the response body contains Invalid refresh_token')
def and_no_refresh(response):
    res: Response = response[0]

    assert "Invalid refresh_token" in res.text


# SCENARIO 3: Attempt to get new access_token with no refresh token value
@when("I send a POST request to /refresh with no refresh_token value")
def when_no_refresh(response):
    res: Response = backend.post_refresh("empty")
    response.append(res)


# (runs after then_response_bad_request)
@then('the response body contains Value for refresh_token is empty')
def and_no_refresh(response):
    res: Response = response[0]

    assert "Value for refresh_token is empty" in res.text


# SCENARIO 4: Attempt to get new access_token with empty request
@when("I send a POST request to /refresh with empty request")
def when_no_refresh(response):
    res: Response = backend.post_refresh("none")
    response.append(res)


# (runs after then_response_bad_request)
@then('the response body contains Value for refresh_token is empty')
def and_no_refresh(response):
    res: Response = response[0]

    assert "Value for refresh_token is empty" in res.text


# SCENARIO 5: Attempt to call /refresh with invalid request method
@when(parsers.parse("I send a {method} request to /refresh"),
      converters={"method": str})
def when_invalid_method(response, method: str):
    res: Response = backend.wrong_method_api_call(backend.Urls.POST_REFRESH, method=method)
    response.append(res)


@then('the response status_code is 405')
def then_invalid_method(response):
    res: Response = response[0]

    assert res.status_code == 405


@then('the response body contains Only POST method allowed')
def and_invalid_method(response):
    res: Response = response[0]

    assert "Only POST method allowed" in res.text




