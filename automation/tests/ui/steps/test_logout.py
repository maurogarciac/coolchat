import pytest
from pytest_bdd import given, when, then, scenarios
from selenium import webdriver

from automation.config.settings import Users
from automation.tests.ui.pages.log_out import LogOutPage
from automation.tests.ui.pages.log_in import LogInPage
from automation.tests.ui.pages.home import HomePage


scenarios("../features/logout.feature")


@pytest.fixture()
def logout_page(driver: webdriver) -> LogOutPage:
    """Perform a full login and navigate to the LogOutPage"""
    login_page: LogInPage = LogInPage(driver)
    login_page.full_log_in(Users.Alice.get("username"), Users.Alice.get("password"))
    home_page: HomePage = HomePage(driver)
    home_page.click_nav_log_out()

    return LogOutPage(driver)


# Common steps
@given('I am authenticated and can access the log out page')
def given_logout_accessible(logout_page):
    assert "logout" in logout_page.get_url()


# Scenario 1: Successfully log out
@when('I click the "Yes" button')
def when_click_yes(logout_page):
    logout_page.click_yes_button()


@then('my token credentials are removed')
def then_credentials_removed(logout_page):
    cookie: str = logout_page.get_cookie("access_token")

    assert not cookie


@then('I am redirected to the log in page')
def and_redirect_login(logout_page):

    assert logout_page.wait_for_url("login")


# Scenario: Cancel log out
@when('I click the "No, I want to chat" button')
def when_click_no(logout_page):
    logout_page.click_no_button()


@then('I am redirected to the chat page')
def and_redirect_chat(logout_page):

    assert logout_page.wait_for_url("chat")
