Feature: Landing page redirection

    As a user
    I want to have access to a landing page
    That can redirect me to home or log-in depending on credentials

    Scenario: Go to landing unauthenticated
      Given I am not authenticated
      When I open the coolchat page
      Then I am redirected to the login page

    Scenario: Go to landing authenticated
      Given I am an authenticated user
      When I open the coolchat page
      Then I am redirected to the home page
