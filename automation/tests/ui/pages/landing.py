from logging import Logger, getLogger

from selenium import webdriver

from automation.config.settings import Urls
from automation.tests.ui.pages.common import CommonOperations
from automation.tests.ui.utils.name_extractor import get_fn_name

lg: Logger = getLogger(__name__)


class LandingPage(CommonOperations):
    """The landing page (if we can even call it a page)"""

    def __init__(self, driver: webdriver):
        super().__init__(driver)
        self.driver: webdriver = driver

    def go_to_page(self) -> None:
        """ Open the landing page """
        self.driver.get(Urls.LANDING_PAGE)
        lg.debug(f"TEST {self.__class__.__name__} --- Completed: {get_fn_name()}")
