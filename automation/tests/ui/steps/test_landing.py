import pytest
from pytest_bdd import given, when, then, parsers, scenario, scenarios
from selenium import webdriver

from automation.config.settings import Users
from automation.tests.ui.pages.landing import LandingPage
from automation.tests.ui.pages.log_in import LogInPage
from automation.tests.ui.pages.home import HomePage

scenarios("../features/landing.feature")


# Common steps
@when('I open the coolchat page')
def when_landing_opens(driver: webdriver):
    page = LandingPage(driver)
    page.go_to_page()


# Scenario 1: Go to landing unauthenticated
@given('I am not authenticated')
def given_not_authenticated():
    pass


@then('I am redirected to the login page')
def when_redirect_login(driver: webdriver):
    page = LogInPage(driver)

    assert "login" in page.get_url()


# Scenario 2: Go to landing authenticated
@given('I am an authenticated user')
def given_authenticated(driver: webdriver):
    page = LogInPage(driver)
    page.full_log_in(Users.Alice.get("username"), Users.Alice.get("password"))


@then('I am redirected to the home page')
def when_redirect_home(driver: webdriver):
    page = HomePage(driver)

    assert "home" in page.get_url()
