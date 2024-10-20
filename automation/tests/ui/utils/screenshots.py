import re
from datetime import datetime
from logging import Logger, getLogger
from os import path, makedirs, getcwd

from _pytest.fixtures import FixtureRequest
from selenium import webdriver
from selenium.webdriver.common.by import By

lg: Logger = getLogger(__name__)


class ScreenshotUtils:
    timestamp: str | None = None

    def __init__(self):
        self.timestamp = datetime.now().strftime("%Y-%b-%d-%H:%M")

    def save_picture(self, request: FixtureRequest, driver: webdriver) -> None:
        """Take two screenshots, one of the full page and a partial screenshot and save them as a .png file

        :param request: Pytest request fixture
        :param driver: Current webdriver
        """

        class_str = re.search(r"\.(Test[A-Z]+)", str(request.cls)).group(
            1
        )  # A little stinky regex to find the class name from a full object str

        if driver is not None:
            screenshot_file_name: str = path.join(
                f"{self._make_screenshots_dir(class_str)}", f"{request.node.name}"
            )
            driver.find_element(By.TAG_NAME, "body").screenshot(
                f"{screenshot_file_name}-full.png"
            )
            driver.save_screenshot(f"{screenshot_file_name}.png")

    def _make_screenshots_dir(self, test_class: str) -> str:
        """Create a directory to store screenshots for a test (if it doesn't exist already)

        :param test_class: Pytest Test Class

        :returns: Full path to screenshot directory
        """

        screenshots_directory: str = path.join(getcwd(), "reports/screenshots")
        makedirs(
            screenshots_directory, exist_ok=True
        )  # Creates screenshots directory if it doesn't exist

        if self.timestamp is None:
            self.timestamp = datetime.now().strftime("%Y-%m-%d-%H:%M")
        new_dir_name: str = f"{self.timestamp}-{test_class}"
        new_path: str = path.join(screenshots_directory, new_dir_name)

        candidate: str = new_path
        i = 0
        dir_exists: bool = path.exists(candidate)
        if dir_exists:  # If the candidate doesn't exist, create it. Otherwise, append an int to the filename.
            while path.exists(candidate):
                i += 1
                candidate = f"{new_path}_{i}"
                makedirs(candidate)
        elif not dir_exists:
            makedirs(candidate)
        return candidate
