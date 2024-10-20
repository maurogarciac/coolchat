from logging import Logger, getLogger

from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.remote.webelement import WebElement

from automation.tests.ui.pages.common import CommonOperations
from automation.tests.ui.utils.name_extractor import get_fn_name

lg: Logger = getLogger(__name__)


class BasePage(CommonOperations):
    """Attributes common to all the app's pages."""

    base_locators: dict = {
        # Navigation bar
        "nav_home": (By.ID, "home"),
        # Page general
        "pg_title": (By.CLASS_NAME, "pg-title"),
        "pg_subtitle": (By.CLASS_NAME, "pg-subtitle"),
        "notification": (By.CLASS_NAME, "notification"),
    }

    def __init__(self, driver: webdriver):
        super().__init__(driver)
        self.driver: webdriver = driver

    def get_browser_title(self) -> str:
        """Get title from browser tab
        :returns: Title text from the browser tab
        """
        title: str = self.driver.title
        lg.debug(f"TEST {self.__class__.__name__} --- Completed: {get_fn_name()}")
        return title

    def get_title(self) -> str:
        """Get title from the html body
        :returns: Title text in the html body
        """
        title: WebElement = self.wait_for(self.base_locators["pg_title"])
        lg.debug(f"TEST {self.__class__.__name__} --- Completed: {get_fn_name()}")
        return title.text

    def get_subtitle(self) -> str:
        """Get subtitle from the html body
        :returns: Subtitle text in the html body
        """
        title: WebElement = self.wait_for(self.base_locators["pg_subtitle"])
        lg.debug(f"TEST {self.__class__.__name__} --- Completed: {get_fn_name()}")
        return title.text

    def get_subtitles(self) -> list[str]:
        """Get all subtitles from the html body
        :returns: List of subtitles text in the html body
        """
        elements: list[WebElement] = self.driver.find_elements(
            *self.base_locators["pg_subtitle"]
        )
        titles: list[str] = []
        for e in elements:
            titles.append(e.text)
        lg.debug(f"TEST {self.__class__.__name__} --- Completed: {get_fn_name()}")
        return titles

    def get_url(self) -> str:
        """Get url from the browser
        :returns: Url string
        """
        lg.debug(f"TEST {self.__class__.__name__} --- Completed: {get_fn_name()}")
        return self.driver.current_url

    def get_notification_text(self) -> str:
        """Get the text from notification at the top of the page.
        :returns: Notification text in the html body
        """
        notification: WebElement = self.driver.find_element(
            *self.base_locators["notification"]
        )
        lg.debug(f"TEST {self.__class__.__name__} --- Completed: {get_fn_name()}")
        return notification.text

    def get_notifications_text(self) -> list[str]:
        """Get a list of strings of text contained by all present notifications
        :returns: List of notifications text
        """
        notifications: list[WebElement] = self.driver.find_elements(
            *self.base_locators["notification"]
        )
        return [n.text for n in notifications]

    # Nav
    def click_nav_home(self) -> None:
        """Redirects to HomePage"""
        self.driver.find_element(*self.base_locators["home"]).click()
        lg.debug(f"TEST {self.__class__.__name__} --- Completed: {get_fn_name()}")

    def clear_cookies(self) -> None:
        """Clear the browser cookies.

        Could be useful to quickly log-out of a user without going through the whole manual process.
        """
        self.driver.delete_all_cookies()
        lg.debug(f"TEST {self.__class__.__name__} --- Completed: {get_fn_name()}")
