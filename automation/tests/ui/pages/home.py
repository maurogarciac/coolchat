from logging import Logger, getLogger
from typing import Optional

from selenium import webdriver
from selenium.webdriver.common.by import By

from automation.config.settings import Urls
from automation.tests.ui.pages.unauthenticated import UnauthenticatedPage
from automation.tests.ui.pages.sign_in import SignInPage
from automation.tests.ui.utils.name_extractor import get_fn_name

lg: Logger = getLogger(__name__)


class HomePage(UnauthenticatedPage):
    """The home page"""

    locators: dict = {
        "button_sign_up": (By.ID, "sign_up"),
        "button_sign_in": (By.ID, "sign_in"),
    }

    def __init__(self, driver: webdriver):
        super().__init__(driver)
        self.driver: webdriver = driver

    def go_to_page(self) -> None:
        self.driver.get(Urls.HOME_PAGE)
        lg.debug(f"TEST {self.__class__.__name__} --- Completed: {get_fn_name()}")

    # Page content

    # def click_create_account_button(self) -> SignUpPage:
    #     button = self.wait_for(self.locators["button_sign_up"])
    #     button.click()
    #     lg.debug(f"TEST {self.__class__.__name__} --- Completed: {get_fn_name()}")
    #     return SignUpPage(self.driver)

    def click_sign_in_button(self) -> SignInPage:
        button = self.wait_for(self.locators["button_sign_in"])
        button.click()
        lg.debug(f"TEST {self.__class__.__name__} --- Completed: {get_fn_name()}")
        return SignInPage(self.driver)