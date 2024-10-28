Feature: Chatting

    As a user
    I want to send and recieve messages in the chat
    To interact with other users

    Scenario: Send and recieve a chat message
      Given I am authenticated as test_alice in the chat page
      And I send a message in the chat
      When I open the chat app as test_bob
      Then I can see test_alice's message in the chat
