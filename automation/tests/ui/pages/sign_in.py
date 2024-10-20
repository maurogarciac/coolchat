from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.remote.webelement import WebElement

from automation.config.settings import Urls

from automation.tests.ui.pages.unauthenticated import UnauthenticatedPage


class SignInPage(UnauthenticatedPage):
    """The sign-up page"""

    locators: dict = {
        "input_email": (By.ID, "email"),
        "input_password": (By.ID, "password"),
        "button_sign_in": (By.ID, "sign_in"),
    }

    def __init__(self, driver: webdriver):
        super().__init__(driver)
        self.driver: webdriver = driver

    def go_to_page(self) -> None:
        self.driver.get(Urls.SIGN_IN_PAGE)

    def enter_email(self, email: str) -> None:
        element: WebElement = self.driver.find_element(*self.locators["email"])
        element.clear()
        element.send_keys(email)

    def enter_password(self, password: str) -> None:
        element: WebElement = self.driver.find_element(*self.locators["password"])
        element.clear()
        element.send_keys(password)

    def click_sign_in_button(self) -> None:
        self.driver.find_element(*self.locators["sign_in"]).click()