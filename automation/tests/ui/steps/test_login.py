import pytest
from pytest_bdd import given, when, then, parsers, scenarios
from selenium import webdriver

from automation.config.settings import Users
from automation.tests.ui.pages.log_in import LogInPage
from automation.tests.ui.pages.home import HomePage

scenarios("../features/login.feature")


@pytest.fixture()
def login_page(driver: webdriver) -> LogInPage:
    return LogInPage(driver)


# Common steps
@given('I can access the login page')
def given_login_accessible(login_page):
    page: LogInPage = login_page
    page.go_to_page()

    assert "login" in page.get_url()


# Scenario 1: Successfully log in
@when('I try to log in with my valid credentials')
def when_login_valid(login_page):
    page: LogInPage = login_page
    page.enter_username(Users.Alice.get("username"))
    page.enter_password(Users.Alice.get("password"))
    page.click_log_in_button()

    assert not page.error_present()


@then('I am redirected to the home page')
def then_redirect_home(driver: webdriver):
    page = HomePage(driver)

    assert "home" in page.get_url()


# Scenario 2: Attempt to log in with incorrect credentials
@when(parsers.parse('I try to log in with my invalid {username} or {password}'), converters={"username": str, "password": str})
def when_login_invalid(login_page, username: str, password: str):
    page: LogInPage = login_page
    page.enter_username(username)
    page.enter_password(password)
    page.click_log_in_button()


@then('I receive an error for Incorrect username or password')
def then_login_error(login_page):
    page: LogInPage = login_page

    assert page.error_present()
    assert "login" in page.get_url()


# Scenario 3: Attempt to access page sections without authentication
@when(parsers.parse('I click the {button} nav button'), converters={"button": str})
def when_click_nav_button(login_page: LogInPage, button: str):
    page: LogInPage = login_page

    match button:
        case "home":
            page.click_nav_home()
        case "chat":
            page.click_nav_chat()


@then('I am redirected to the login page')
def then_redirect_login(login_page):
    page: LogInPage = login_page

    assert "Sign in" in page.get_page_subtitle()
