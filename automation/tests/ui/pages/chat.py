from logging import Logger, getLogger
from typing import Optional
import re

from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.remote.webelement import WebElement

from automation.config.settings import Urls
from automation.tests.ui.pages.authenticated import AuthenticatedPage
from automation.tests.ui.utils.name_extractor import get_fn_name

lg: Logger = getLogger(__name__)


class ChatPage(AuthenticatedPage):
    """The chat page"""

    locators: dict = {
        "last_message": (By.CSS_SELECTOR, "#message-wrapper:last-of-type"),
        "input_chat": (By.CSS_SELECTOR, "form input"),
        "input_submit": (By.CSS_SELECTOR, "input[type='submit']"),
    }

    def __init__(self, driver: webdriver):
        super().__init__(driver)
        self.driver: webdriver = driver

    def go_to_page(self) -> None:
        self.driver.get(Urls.HOME_PAGE)
        lg.debug(f"TEST {self.__class__.__name__} --- Completed: {get_fn_name()}")

    # Page content

    def click_chat_href(self) -> None:
        self.wait_for(self.locators["href_for_chat"]).click()
        lg.debug(f"TEST {self.__class__.__name__} --- Completed: {get_fn_name()}")

    def write_message(self, message: str) -> None:
        element: WebElement = self.driver.find_element(*self.locators["input_chat"])
        element.clear()
        element.send_keys(message)

    def send_message(self) -> None:
        self.driver.find_element(*self.locators["input_submit"]).click()

    def get_username_from_input_placeholder(self) -> str:
        element: WebElement = self.driver.find_element(*self.locators["input_chat"])
        preview: str = element.get_attribute("placeholder")
        match = re.search(r",\s*(\w+)\?", preview)
        if match:
            return match.group(1)
        return ""

    def get_last_message(self) -> dict[str, str]:
        """Get the last message in chat
            :returns: Message dict with 'sender', 'text' and 'ts' (time-stamp)
        """
        msg_wrp: WebElement = self.wait_for(self.locators["last_message"])

        message: dict[str, str] = {
            "sender": msg_wrp.find_element(By.ID, "sent-by").text,
            "ts": msg_wrp.find_element(By.ID, "time-sent").text,
            "text": msg_wrp.find_element(By.ID, "msg-content").text,
        }

        return message
