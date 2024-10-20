from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.remote.webelement import WebElement
from selenium.webdriver.support.wait import WebDriverWait
from selenium.webdriver.support import expected_conditions as ec


class CommonOperations(object):
    """
    A wrapper for some of selenium's commonly used expressions.
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

        :param url: Value contained in url
        :param wait_time: Max wait time for the WebDriverWait

        """

        wait = WebDriverWait(self.driver, wait_time)
        return wait.until(ec.url_contains(url))
