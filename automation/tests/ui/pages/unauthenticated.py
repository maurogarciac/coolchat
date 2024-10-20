from logging import Logger, getLogger

from automation.tests.ui.pages.base import BasePage
from automation.tests.ui.utils.name_extractor import get_fn_name


lg: Logger = getLogger(__name__)


class UnauthenticatedPage(BasePage):
    """Attributes common to pages where a user has not signed in."""

    locators = {
        "sign_up": ("id", "sign_up"),
        "sign_in": ("id", "sign_in"),
    }

    def __init__(self, driver):
        super().__init__(driver)
        self.driver = driver

    # Nav Menu
    def click_nav_sign_up(self) -> None:
        """Redirects to SignUpPage"""
        self.driver.find_element(*self.locators["sign_up"]).click()
        lg.debug(f"TEST {self.__class__.__name__} --- Completed: {get_fn_name()}")

    def click_nav_sign_in(self) -> None:
        """Redirects to SignInPage"""
        self.driver.find_element(*self.locators["sign_in"]).click()
        lg.debug(f"TEST {self.__class__.__name__} --- Completed: {get_fn_name()}")
