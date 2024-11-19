Feature: Calculate factorial
  To ensure that the factorial function works correctly

  Scenario: Calculate factorial of a positive number
    Given the number is 5
    When I calculate the factorial
    Then the result should be 120

  Scenario: Calculate factorial of zero
    Given the number is 0
    When I calculate the factorial
    Then the result should be 1

  Scenario: Calculate factorial of one
    Given the number is 1
    When I calculate the factorial
    Then the result should be 1

  Scenario: Show factorial of five when false
    Given the number is 5
    And the expansion is "false"
    When I calculate the factorial
    Then the result should be 120
    And the result should be shown like "5! = 120"

  Scenario: Show factorial of five when true
    Given the number is 5
    And the expansion is "true"
    When I calculate the factorial
    Then the result should be 120
    And the result should be shown like "5! = 1 * 2 * 3 * 4 * 5 = 120"

  Scenario: Show factorial of one when true
    Given the number is 1
    And the expansion is "true"
    When I calculate the factorial
    Then the result should be 1
    And the result should be shown like "1! = 1 = 1"

  Scenario: Show factorial of zero when true
    Given the number is 0
    And the expansion is "true"
    When I calculate the factorial
    Then the result should be 1
    And the result should be shown like "0! = 1 = 1"

  Scenario: Show factorial of three when false
    Given the number is 3
    And the expansion is "false"
    When I calculate the factorial
    Then the result should be 6
    And the result should be shown like "3! = 6"