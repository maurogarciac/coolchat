Feature: Get messages

    As a user
    I want to retrieve a list of messages
    So that I can view message history in the chat

    Scenario: Get message history
      Given the backend api is healthy
      When I send a GET request to /messages
      Then the response status_code is 200
      And the response body contains a json list of messages

    Scenario Outline: Attempt to call /messages with incorrect request method
      Given the backend api is healthy
      When I send a <method> request to /messages
      Then the response status_code is 405
      And the response body contains Only GET method allowed

  Examples:
    | method |
    | POST   |
    | PUT    |
    | PATCH  |
    | DELETE |
