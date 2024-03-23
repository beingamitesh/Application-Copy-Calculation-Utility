# Application Copy Calculation Utility

## Overview

This utility is designed to calculate the minimum number of copies of a specific application (ID 374) that a company must purchase, considering the restrictions on installations per user and computer types. It processes provided data regarding installations of the application and determines the optimal number of copies needed based on the given criteria.

## Expectations

The solution aims to meet the following expectations:
- Calculate the minimum number of copies of the application (ID 374) required for company purchase.
- Include unit tests to demonstrate basic test coverage.
- Adhere to high-quality coding standards expected in a production environment, akin to a world-class product.
- Incorporate appropriate object-oriented (OO) modeling, considering the assignment's role in a larger product at Flexera.
- Address non-functional concerns to ensure a software product of high quality.

## Assumptions

- The provided data does not contain empty values.
- Each computer has only one user.
- Computers are either desktops or laptops, not both.

## Example Scenarios

### Example 1

Given the following scenario:

| ComputerID | UserID | ApplicationID | ComputerType | Comment                |
|------------|--------|---------------|--------------|------------------------|
| 1          | 1      | 374           | LAPTOP       | Exported from System A |
| 2          | 1      | 374           | DESKTOP      | Exported from System A |

#### Output:
Minimum copies required: 1

#### Explanation:
Only one copy of the application is required as the user has installed it on two computers, with one of them being a laptop.

### Example 2

Given the following scenario:

| ComputerID | UserID | ApplicationID | ComputerType | Comment                |
|------------|--------|---------------|--------------|------------------------|
| 1          | 1      | 374           | LAPTOP       | Exported from System A |
| 2          | 1      | 374           | DESKTOP      | Exported from System A |
| 3          | 2      | 374           | DESKTOP      | Exported from System A |
| 4          | 2      | 374           | DESKTOP      | Exported from System A |

#### Output:
Minimum copies required: 3

#### Explanation:
Three copies of the application are required as UserID 2 has installed the application on two computers, but neither of them is a laptop, necessitating a purchase of the application for both computers.

### Example 3

Given the following scenario:

| ComputerID | UserID | ApplicationID | ComputerType | Comment                |
|------------|--------|---------------|--------------|------------------------|
| 1          | 1      | 374           | LAPTOP       | Exported from System A |
| 2          | 2      | 374           | DESKTOP      | Exported from System A |
| 2          | 2      | 374           | desktop      | Exported from System B |

#### Output:
Minimum copies required: 2

#### Explanation:
Two copies of the application are required as the data from the second and third rows are effectively duplicates, even though the ComputerType is lowercase and the comments are different.
