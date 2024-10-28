from selenium import webdriver
from selenium.common import NoSuchElementException
from selenium.webdriver.common.by import By
from selenium.webdriver.remote.webelement import WebElement

from automation.config.settings import Urls
from automation.tests.ui.pages.unauthenticated import UnauthenticatedPage


class LogInPage(UnauthenticatedPage):
    """The Log-in page"""

    locators: dict = {
        "input_email": (By.ID, "username"),
        "input_password": (By.ID, "password"),
        "form_log_in": (By.CSS_SELECTOR, "#login button"),
        "form_error": (By.CLASS_NAME, "error"),
    }

    def __init__(self, driver: webdriver):
        super().__init__(driver)
        self.driver: webdriver = driver

    def go_to_page(self) -> None:
        self.driver.get(Urls.LOG_IN_PAGE)

    def enter_username(self, username: str) -> None:
        element: WebElement = self.driver.find_element(*self.locators["input_email"])
        element.clear()
        element.send_keys(username)

    def enter_password(self, password: str) -> None:
        element: WebElement = self.driver.find_element(*self.locators["input_password"])
        element.clear()
        element.send_keys(password)

    def click_log_in_button(self) -> None:
        self.driver.find_element(*self.locators["form_log_in"]).click()

    def get_error_login(self) -> "WebElement":
        return self.driver.find_element(*self.locators["form_error"])

    def error_present(self) -> bool:
        try:
            self.driver.find_element(*self.locators["form_error"])
            return True
        except NoSuchElementException:
            return False

    def full_log_in(self, username: str, password: str) -> None:
        """ Perform a full log-in

            :param username: Username string
            :param password: Password string
        """

        self.go_to_page()
        self.enter_username(username)
        self.enter_password(password)
        self.click_log_in_button()
