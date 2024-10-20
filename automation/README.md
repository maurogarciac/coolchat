# Introduction 

Runs a suite of tests, ranging from web/ui to api, to confirm that site navigation and features work as
expected. 

## Requirements:

* Python 3.11.2

## Optional *recommended* requirements:

* Make 4.4.1
* Docker 26.1.1

## Local setup steps:

1. Create a Virtual Environment and name it `.venv`:
    ```shell
    python -m venv .venv
    ```
2. Activate the Environment:
    - Linux
    ```shell
    chmod +x .venv/bin/activate
    source .venv/bin/activate
    ```
    - MacOs
    ```shell
    source .venv/Scripts/activate
    ```
3. Install the required Packages:
    ```shell
    python -m pip install -r requirements/requirements.txt
    ``` 
4. Create an *.env* file with the contents of *.env.example*

## Docker setup steps:
Either run `docker-compose up` or `make d_test`

## Run tests:  
To run all the tests, just run `pytest` in local env or `make d_test` for docker container.  

### Flags for execution:
An optional `--browser` flag can be included after `pytest` to specify the browser option (defaults to *chrome*):  

> * chrome
> * firefox
> * chrome_headless
> * remote (not currently setup since it requires remote-wd server)
    
#### These are the commands to run the test groups, __api__ and __ui__: 

- Locally with pytest:
    ```shell
    pytest tests/api
    pytest tests/ui
    ```
- Locally with Make (to run in visual browsers):
    ```shell
    make test_ui_f
    make test_ui_c
    ```
- Run a single test locally:
    ```shell
    pytest --browser chrome tests/ui/steps/test_login.py::TestLogin::login
    ```
    (First the file path, then the Class name, and last goes the Test method)

### Test execution environment:
The execution environment (production, test, development or local) can be changed in the '.env' file on the line:
```.env
   ENV='test'
```

# To do:

1. Implement parallelism (pytest-xdist or pytest-parallel)


## Structure: Page Object Model

These tests are implemented with the Page Object Model to make them DRY and clean. That means there are modules that
provide classes to represent each page of the site as a user sees it.

- `tests`: This module defines the test configuration for api and ui in 'conftest.py'.
- `api`: The api tests, separated by api.
- `api.models`: Adapters for each api's model.
- `api.services`: Services to instantiate a connection to each api. 
- `ui.pages`: This module has a class for each site page and includes methods to access each interactive element.
- `ui.steps`: This module contains all the tests, separated by page.
