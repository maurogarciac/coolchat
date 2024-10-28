from logging import Logger, getLogger

from selenium import webdriver
from selenium.webdriver.common.by import By

from automation.tests.ui.pages.common import CommonOperations
from automation.tests.ui.utils.name_extractor import get_fn_name

lg: Logger = getLogger(__name__)


class BasePage(CommonOperations):
    """Attributes common to all the app's pages."""

    base_locators: dict = {
        # Navigation bar
        "nav_home": (By.ID, "nav-home"),
        "nav_chat": (By.ID, "nav-chat"),
        "page_title": (By.CSS_SELECTOR, "#main h1"),
        "page_subtitle": (By.CSS_SELECTOR, "#main h2"),
    }

    def __init__(self, driver: webdriver):
        super().__init__(driver)
        self.driver: webdriver = driver

    # Nav
    def click_nav_home(self) -> None:
        """Redirects to HomePage"""
        self.driver.find_element(*self.base_locators["nav_home"]).click()
        lg.debug(f"TEST {self.__class__.__name__} --- Completed: {get_fn_name()}")

    def click_nav_chat(self) -> None:
        """Redirects to ChatPage"""
        self.driver.find_element(*self.base_locators["nav_chat"]).click()
        lg.debug(f"TEST {self.__class__.__name__} --- Completed: {get_fn_name()}")

    # Page content 
    def get_page_title(self) -> str:
        """ Get the current page's title"""
        e = self.driver.find_element(*self.base_locators["page_title"])
        return e.text

    def get_page_subtitle(self) -> str:
        """ Get the current page's subtitle"""
        e = self.driver.find_element(*self.base_locators["page_subtitle"])
        return e.text