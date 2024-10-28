from logging import Logger, getLogger
from typing import Optional

from selenium import webdriver
from selenium.webdriver.common.by import By

from automation.config.settings import Urls
from automation.tests.ui.pages.authenticated import AuthenticatedPage
from automation.tests.ui.utils.name_extractor import get_fn_name

lg: Logger = getLogger(__name__)


class HomePage(AuthenticatedPage):
    """The home page"""

    locators: dict = {
        "href_for_chat": (By.CSS_SELECTOR, "#main a"),
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
