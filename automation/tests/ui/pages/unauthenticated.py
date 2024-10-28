from logging import Logger, getLogger

from automation.tests.ui.pages.base import BasePage
from automation.tests.ui.utils.name_extractor import get_fn_name
from selenium.webdriver.common.by import By

lg: Logger = getLogger(__name__)


class UnauthenticatedPage(BasePage):
    """Attributes common to pages where a user is not signed in."""

    locators = {
        "nav_log_in": (By.ID, "nav-login"),
    }

    def __init__(self, driver):
        super().__init__(driver)
        self.driver = driver

    # Nav Menu

    def click_nav_log_in(self) -> None:
        """Redirects to LoginPage"""
        self.driver.find_element(*self.locators["nav_log_in"]).click()
        lg.debug(f"TEST {self.__class__.__name__} --- Completed: {get_fn_name()}")
