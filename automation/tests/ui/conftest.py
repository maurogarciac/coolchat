from logging import Logger, getLogger

import pytest
from _pytest.fixtures import FixtureRequest
from selenium import webdriver
from selenium.webdriver.firefox.service import Service as F_Service
from selenium.webdriver.chrome.service import Service as C_Service
from webdriver_manager.chrome import ChromeDriverManager
from webdriver_manager.firefox import GeckoDriverManager

from automation.tests.ui.utils.randomizer import get_random_string

lg: Logger = getLogger(__name__)
IMPLICIT_TIMEOUT: float = 3


@pytest.fixture(autouse=True)
def driver(request: FixtureRequest):
    global _driver
    browser = request.config.option.browser

    if browser == "firefox":
        _driver = webdriver.Firefox(
            service=F_Service(
                GeckoDriverManager().install(), log_path="./reports/geckodriver.log"
            )
        )
    elif (
        browser == "remote"
    ):  # Not implemented yet
        capabilities = {"browserName": "firefox", "javascriptEnabled": True}
        _driver = webdriver.Remote(
            command_executor="http://127.0.0.1:4444/wd/hub",
            desired_capabilities=capabilities,
        )
    elif browser == "chrome_headless":  # Mock the user agent to prevent being automatically blocked
        user_agent: str = (
            "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.50 "
            "Safari/537.36"
        )
        op = webdriver.ChromeOptions()
        op.add_argument("--headless")
        op.add_argument("--disable-dev-shm-usage")
        op.add_argument("--no-sandbox")
        op.add_argument("window-size=1920x1080")
        op.add_argument(f"user_agent={user_agent}")
        _driver = webdriver.Chrome(
            service=C_Service(ChromeDriverManager().install()), options=op
        )
    elif browser == "chrome":
        _driver = webdriver.Chrome(service=C_Service(ChromeDriverManager().install()))
    else:
        lg.error("Something's wrong initializing the web-driver! Fix it!")
    _driver.implicitly_wait(IMPLICIT_TIMEOUT)
    _driver.maximize_window()
    lg.info(f"Web Driver initialized as: {browser}")
    yield _driver

    # This bit of code takes a screenshot if a test fails

    # if request.node.rep_call.failed:  # If the test fails, take a screenshot
    #    su: ScreenshotUtils = ScreenshotUtils()
    #    su.save_picture(request, _driver)
    #    lg.info(f"Error screenshot saved for {request.node.name}")
    _driver.quit()

@pytest.fixture(autouse=False, scope="function")
def random_email() -> str:
    """Generate a random email address

    :returns: Random email string
    """
    return f"{get_random_string(7)}@mgc.sh"
