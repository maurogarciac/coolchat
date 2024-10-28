import pytest
from pytest_bdd import given, when, then, scenarios
from selenium import webdriver
from selenium.webdriver.remote.webelement import WebElement

from automation.config.settings import Users
from automation.tests.ui.pages.log_in import LogInPage
from automation.tests.ui.pages.home import HomePage
from automation.tests.ui.pages.chat import ChatPage


scenarios("../features/chat.feature")


@pytest.fixture()
def chat_page(driver: webdriver) -> ChatPage:
    """Perform a full login and navigate to the ChatPage as user test_alice"""
    login_page: LogInPage = LogInPage(driver)
    login_page.full_log_in(Users.Alice.get("username"), Users.Alice.get("password"))
    home_page: HomePage = HomePage(driver)
    home_page.click_chat_href()

    return ChatPage(driver)


# Scenario 1: Send and recieve a chat message
@given('I am authenticated as test_alice in the chat page')
def given_x(chat_page):
    assert "chat" in chat_page.get_url()
    assert "test_alice" in chat_page.get_username_from_input_placeholder()


@given('I send a message in the chat')
def and_x(chat_page):
    chat_page.write_message("Test message")
    chat_page.send_message()

    message: dict[str, str] = chat_page.get_last_message()
    assert message, "Last message not found in chatbox"
    assert "test_alice" in message.get('sender')
    assert "Test message" in message.get('text')


@when('I open the chat app as test_bob')
def when_x(chat_page, driver):
    chat_page.clear_cookies()

    login_page: LogInPage = LogInPage(driver)
    login_page.go_to_page()
    login_page.enter_username(Users.Bob.get("username"))
    login_page.enter_password(Users.Bob.get("password"))
    login_page.click_log_in_button()

    login_page.click_nav_chat()
    assert "chat" in login_page.get_url()


@then("I can see test_alice's message in the chat")
def then_x(driver):
    chat_page: ChatPage = ChatPage(driver)

    message: dict[str, str] = chat_page.get_last_message()
    assert message, "Last message not found in chatbox"
    assert "test_alice" in message.get('sender')
    assert "Test message" in message.get('text')
