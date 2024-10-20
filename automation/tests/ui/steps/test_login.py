import random
import string
from logging import Logger, getLogger

import pytest
from selenium import webdriver
from selenium.webdriver.support import expected_conditions as ec
from selenium.webdriver.support.ui import WebDriverWait

from automation.config.settings import Urls
from automation.tests.ui.pages.home import HomePage
from automation.tests.ui.pages.unauthenticated import UnauthenticatedPage

from pytest_bdd import scenario, given, when, then  # Implement pytestbdd

# @scenario('publish_article.feature', 'Publishing the article')
# def test_publish():
#     pass


# @given("I'm an author user")
# def author_user(auth, author):
#     auth['user'] = author.user


# @given("I have an article", target_fixture="article")
# def article(author):
#     return create_test_article(author=author)


# @when("I go to the article page")
# def go_to_article(article, browser):
#     browser.visit(urljoin(browser.url, '/manage/articles/{0}/'.format(article.id)))


# @when("I press the publish button")
# def publish_article(browser):
#     browser.find_by_css('button[name=publish]').first.click()


# @then("I should not see the error message")
# def no_error_message(browser):
#     with pytest.raises(ElementDoesNotExist):
#         browser.find_by_css('.message.error').first


# @then("the article should be published")
# def article_is_published(article):
#     article.refresh()  # Refresh the object in the SQLAlchemy session
#     assert article.is_published

lg: Logger = getLogger(__name__)


class TestLogin:
    def test_home_page(self, driver: webdriver):
        """Open the home page"""
        home_page: HomePage = HomePage(driver)
        home_page.go_to_page()

        url: bool = home_page.wait_for_url(Urls.HOME_PAGE)
        assert url, "Not the home page url"
        assert "Welcome" in home_page.get_title()

    def test_open_signup(self, driver: webdriver):
        """Navigate to Sign Up"""
        home_page: HomePage = HomePage(driver)
        home_page.go_to_page()
        sign_up_page: SignUpPage = home_page.click_create_account_button()

        url: bool = sign_up_page.wait_for_url(Urls.SIGN_UP_PAGE)
        assert url, "Not the sign-up page"
        assert "Sign Up" in sign_up_page.get_title()

    def test_email_required(self, driver: webdriver):
        """Email is required"""
        signup_page: SignUpPage = SignUpPage(driver)
        signup_page.go_to_page()
        signup_page.click_sign_up_button()

        assert "Please fill out this field." in signup_page.get_email_error()

    def test_password_required(self, driver: webdriver):
        """Password is required"""
        signup_page: SignUpPage = SignUpPage(driver)
        signup_page.go_to_page()
        signup_page.enter_email("passwd_test@domain.com")
        signup_page.click_sign_up_button()

        assert "Please fill out this field." in signup_page.get_passwd_error()

    def test_password_length(self, driver: webdriver):
        """Password must contain at least 8 characters."""
        page = SignUpPage(driver)
        page.go_to_page()
        page.enter_email("short_password@domain.com")
        page.enter_password("short")
        page.accept_terms_and_conditions()
        page.click_sign_up_button()

        expected: str = (
            "This password is too short. It must contain at least 8 characters."
        )
        assert expected in page.get_passwd_validation_error()

    def test_password_common(self, driver: webdriver):
        """Password can't be a commonly used password."""
        page = SignUpPage(driver)
        page.go_to_page()
        page.enter_email("common_password@domain.com")
        page.enter_password("password")
        page.accept_terms_and_conditions()
        page.click_sign_up_button()

        assert "This password is too common." in page.get_passwd_validation_error()

    def test_password_numeric(self, driver: webdriver):
        """Password can't be entirely numeric."""
        password = "".join(
            random.SystemRandom().choice(string.digits) for _ in range(16)
        )

        page = SignUpPage(driver)
        page.go_to_page()
        page.enter_email("numeric_password@domain.com")
        page.enter_password(password)
        page.accept_terms_and_conditions()
        page.click_sign_up_button()

        assert (
            "This password is entirely numeric." in page.get_passwd_validation_error()
        )

    def test_signup_success(self, driver: webdriver, random_email):
        """Successful sign up (not confirmed)"""
        email: str = random_email

        page: SignUpPage = SignUpPage(driver)
        page.go_to_page()
        page.enter_email(email)
        page.enter_team_name("Test team!")
        page.enter_password("S3cur3P4ssw0rd!")
        page.accept_terms_and_conditions()
        page.click_sign_up_button()

        assert page.wait_for_url(Urls.HOME_PAGE)

