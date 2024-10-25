import pytest
from pytest_bdd import given, when, then, parsers, scenario, scenarios
from requests import Response

from automation.tests.api.services import backend

scenarios("../features/get_messages.feature")


@pytest.fixture
def response() -> list[Response]:
    return []


@given("the backend api is healthy")
def step_given():
    res: Response = backend.get_health()

    assert res.status_code == 200


@when("I send a GET request to /messages")
def when_valid(response):
    res: Response = backend.get_messages()
    response.append(res)


@when(parsers.parse("I send a {method} request to /messages"), converters={"method": str})
def when_invalid(response, method: str):
    res: Response = backend.wrong_method_api_call(backend.Urls.GET_MESSAGES, method=method)
    response.append(res)


@then('the response status_code is 200')
def then_valid(response):
    res: Response = response[0]

    assert res.status_code == 200


@then('the response status_code is 405')
def then_invalid(response):
    res: Response = response[0]

    assert res.status_code == 405


@then('the response body contains a json list of messages')
def and_valid(response):
    res: Response = response[0]

    assert res.text


@then('the response body contains Only GET method allowed')
def and_invalid(response):
    res: Response = response[0]

    assert "Only GET method allowed" in res.text
