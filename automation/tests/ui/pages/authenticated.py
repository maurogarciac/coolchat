from logging import Logger, getLogger

from selenium import webdriver
from selenium.webdriver.common.by import By

from automation.tests.ui.pages.base import BasePage
from automation.tests.ui.utils.name_extractor import get_fn_name

lg: Logger = getLogger(__name__)


class AuthenticatedPage(BasePage):
    """Attributes common to pages where a user has signed in."""

    auth_locators: dict = {
        "nav_sign_out": ( By.XPATH, "sign_out" ),
    }

    def __init__(self, driver: webdriver):
        super().__init__(driver)
        self.driver: webdriver = driver

    # Nav Menu
    def click_nav_tools(self) -> None:
        """Redirects to SubscriptionsPage"""
        self.driver.find_element(*self.auth_locators["nav_tools"]).click()
        lg.debug(f"TEST {self.__class__.__name__} --- Completed: {get_fn_name()}")

    def click_nav_my_team(self) -> None:
        """Redirects to TeamManagementPage"""
        self.driver.find_element(*self.auth_locators["nav_my_team"]).click()
        lg.debug(f"TEST {self.__class__.__name__} --- Completed: {get_fn_name()}")

    def click_nav_profile(self) -> None:
        """Redirects to ProfilePage"""
        self.driver.find_element(*self.auth_locators["nav_profile"]).click()
        lg.debug(f"TEST {self.__class__.__name__} --- Completed: {get_fn_name()}")

    def click_nav_sign_out(self) -> None:
        """Redirects to HomePage"""
        self.driver.find_element(*self.auth_locators["nav_sign_out"]).click()
        lg.debug(f"TEST {self.__class__.__name__} --- Completed: {get_fn_name()}")
