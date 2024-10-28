from selenium import webdriver
from selenium.webdriver.common.by import By

from automation.config.settings import Urls
from automation.tests.ui.pages.authenticated import AuthenticatedPage


class LogOutPage(AuthenticatedPage):
    """The Log-out page"""

    locators: dict = {
        "button_yes": (By.ID, "logout-yes"),
        "button_no": (By.ID, "logout-no"),
    }

    def __init__(self, driver: webdriver):
        super().__init__(driver)
        self.driver: webdriver = driver

    def go_to_page(self) -> None:
        self.driver.get(Urls.LOG_OUT_PAGE)

    def click_yes_button(self) -> None:
        self.driver.find_element(*self.locators["button_yes"]).click()

    def click_no_button(self) -> None:
        self.driver.find_element(*self.locators["button_no"]).click()
