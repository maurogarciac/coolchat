Feature: Get messages

    As a user
    I want to authenticate with my credentials
    So that I can receive a set of access and refresh tokens

    Scenario Outline: Get auth with valid credentials
      Given the backend api is healthy
      When a POST request is sent to /auth with <username> and <password>
      Then the response status_code is 200
      And the response body contains an access_token and refresh_token

      Examples:
        | username | password |
        | bob      | root     |
        | alice    | root     |

    Scenario Outline: Attempt to post auth with invalid credentials
      Given the backend api is healthy
      When a POST request is sent to /auth with invalid <username> or <password>
      Then the response status_code is 403
      And the response body contains User does not exist

      Examples:
        | username | password |
        | dingus   | goober   |
        | glorb    | agorb    |

    Scenario Outline: Attempt to post auth with missing credentials
      Given the backend api is healthy
      When a POST request is sent to /auth with missing <username> or <password>
      Then the response status_code is 400

      Examples:
        | username | password |
        | none     | goober   |
        | glorb    | none     |

    Scenario Outline: Attempt to call /auth with invalid request method
      Given the backend api is healthy
      When a <method> request is sent to /auth
      Then the response status_code is 405
      And the response body contains Only POST method allowed

      Examples:
        | method |
        | GET    |
        | PUT    |
        | PATCH  |
        | DELETE |
