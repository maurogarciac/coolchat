from logging import Logger, getLogger

from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.remote.webelement import WebElement
from selenium.webdriver.support.wait import WebDriverWait
from selenium.webdriver.support import expected_conditions as ec

from automation.tests.ui.utils.name_extractor import get_fn_name

lg: Logger = getLogger(__name__)

class CommonOperations(object):
    """
    A wrapper for some of selenium's commonly used expressions and driver calls.
    """

    def __init__(self, driver: webdriver):
        self.driver: webdriver = driver

    def wait_for(self, locator: tuple[By, str], wait_time: int = 10) -> WebElement:
        """Wait for the presence of an element in the DOM

        :param locator: A tuple composed by a selenium 'By' type and a selector string
        :param wait_time: Max wait time for the WebDriverWait

        :returns: WebElement for the requested locator.
        """

        wait = WebDriverWait(self.driver, wait_time)
        return wait.until(ec.presence_of_element_located(locator))

    def wait_for_visibility(
        self, locator: tuple[By, str], wait_time: int = 10
    ) -> WebElement:
        """Wait for the presence and visibility of an element in the DOM

        :param locator: A tuple composed by a selenium 'By' type and a selector string
        :param wait_time: Max wait time for the WebDriverWait

        :returns: WebElement for the requested locator.
        """

        wait = WebDriverWait(self.driver, wait_time)
        return wait.until(ec.visibility_of_element_located(locator))

    def wait_for_url(self, url: str, wait_time: int = 10) -> bool:
        """Wait for a change in the browser's url

        :param url: Expected value to be contained in url
        :param wait_time: Max wait time for the WebDriverWait

        """

        wait = WebDriverWait(self.driver, wait_time)
        return wait.until(ec.url_contains(url))

    def get_browser_title(self) -> str:
        """Get title from browser tab
        :returns: Title text from the browser tab
        """
        title: str = self.driver.title
        lg.debug(f"TEST {self.__class__.__name__} --- Completed: {get_fn_name()}")
        return title

    def get_url(self) -> str:
        """Get url from the browser
        :returns: Url string
        """
        lg.debug(f"TEST {self.__class__.__name__} --- Completed: {get_fn_name()}")
        return self.driver.current_url

    def clear_cookies(self) -> None:
        """Clear the browser cookies.

        Could be useful to quickly log-out of a user without going through the whole manual process.
        """
        self.driver.delete_all_cookies()
        lg.debug(f"TEST {self.__class__.__name__} --- Completed: {get_fn_name()}")

    def get_cookie(self, cookie_name: str) -> str:
        """Get a browser cookie by name
            :param cookie_name: Cookie name

            :returns: Cookie
        """
        return self.driver.get_cookie(cookie_name)
