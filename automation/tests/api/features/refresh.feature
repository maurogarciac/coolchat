Feature: Get messages

    As a user with an expired access_token
    I want to validate my refresh_token in the backend
    So that I can receive a new access_token

    Scenario: Get new access_token
      Given I have authenticated through POST request to /auth
      And my refresh_token is available
      When I send a POST request to /refresh with my refresh_token
      Then the response status_code is 200
      And the response body contains a new access_token

    Scenario: Attempt to get new access_token with invalid refresh_token
      Given the backend api is healthy
      When I send a POST request to /refresh with invalid refresh_token
      Then the response status_code is 400
      And the response body contains Invalid refresh_token

    Scenario: Attempt to get new access_token with no refresh token value
      Given the backend api is healthy
      When I send a POST request to /refresh with no refresh_token value
      Then the response status_code is 400
      And the response body contains Value for refresh_token is empty

    Scenario: Attempt to get new access_token with empty request
      Given the backend api is healthy
      When I send a POST request to /refresh with empty request
      Then the response status_code is 400
      And the response body contains Value for refresh_token is empty


      Scenario Outline: Attempt to call /refresh with invalid request method
      Given the backend api is healthy
      When I send a <method> request to /refresh
      Then the response status_code is 405
      And the response body contains Only POST method allowed

      Examples:
        | method |
        | GET    |
        | PUT    |
        | PATCH  |
        | DELETE |