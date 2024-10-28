Feature: Removing authentication

    As a user
    I need to have a log out page
    That can remove authentication from the browser I'm using

    Scenario: Successfully log out
      Given I am authenticated and can access the log out page
      When I click the "Yes" button
      Then my token credentials are removed
      And I am redirected to the log in page

    Scenario: Cancel log out
      Given I am authenticated and can access the log out page
      When I click the "No, I want to chat" button
      Then I am redirected to the chat page
