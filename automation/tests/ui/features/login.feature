Feature: Authentication and authorization via login

    As a user
    I need to have a login page
    That can authenticate me and give me authorization to access the chat app

    Scenario: Successfully log in
      Given I can access the login page
      When I try to log in with my valid credentials
      Then I am redirected to the home page

    Scenario Outline: Attempt to log in with incorrect credentials
      Given I can access the login page
      When I try to log in with my invalid <username> or <password>
      Then I receive an error for Incorrect username or password

      Examples:
        | username | password |
        | dingus   | goober   |
        | glorb    | agorb    |

    Scenario Outline: Attempt to access page sections without authentication
      Given I can access the login page
      When I click the <button> nav button
      Then I am redirected to the login page

      Examples:
        | button |
        | home   |
        | chat   |
