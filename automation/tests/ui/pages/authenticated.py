from logging import Logger, getLogger

from selenium import webdriver
from selenium.webdriver.common.by import By

from automation.tests.ui.pages.base import BasePage
from automation.tests.ui.utils.name_extractor import get_fn_name

lg: Logger = getLogger(__name__)


class AuthenticatedPage(BasePage):
    """Attributes common to pages where a user is signed in."""

    auth_locators: dict = {
        "nav_log_out": (By.ID, "nav-logout"),
    }

    def __init__(self, driver: webdriver):
        super().__init__(driver)
        self.driver: webdriver = driver

    # Nav Menu
    def click_nav_log_out(self) -> None:
        """Redirects to LoginPage"""
        self.driver.find_element(*self.auth_locators["nav_log_out"]).click()
        lg.debug(f"TEST {self.__class__.__name__} --- Completed: {get_fn_name()}")
