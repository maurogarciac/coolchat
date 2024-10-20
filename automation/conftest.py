import asyncio
import os
from datetime import datetime

import pytest


# Global test configuration file
def pytest_addoption(parser):  # Reads flags and options
    parser.addoption(
        "--browser",
        action="store",
        default="chrome_headless",
        help="browser to execute the automation",
    )
    parser.addoption(
        "--env",
        action="store",
        default="test",
        help="environment to run the test suite on",
    )


@pytest.hookimpl(tryfirst=True)
def pytest_configure(config):
    if not os.path.exists("reports"):
        os.makedirs("reports")
    config.option.htmlpath = (
        f"reports/test-report-{datetime.now().strftime('%Y-%b-%d-%H:%M')}.html"
    )


@pytest.hookimpl(tryfirst=True, hookwrapper=True)
def pytest_runtest_makereport(item, call):  # Reporting configuration
    outcome = yield
    rep = outcome.get_result()

    setattr(item, "rep_" + rep.when, rep)


@pytest.fixture(scope="session")
def event_loop():
    try:
        loop = asyncio.get_running_loop()
    except RuntimeError:
        loop = asyncio.new_event_loop()
    yield loop
    loop.close()
