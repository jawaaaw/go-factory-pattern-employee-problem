## Go Factory Pattern - Employee Use Case

This is a simple example of the factory pattern in Go. The factory pattern is a creational pattern that provides an interface for creating objects in a super class, but allows subclasses to alter the type of objects that will be created.

This course is part of the Kampus Merdeka program, a collaboration between the Ministry of Education and Culture of the Republic of Indonesia and Ruangguru.

### Details
- Language: Go
- Testing Framework: Ginkgo
- Difficulty: Beginner
- Estimated Time: 30 minutes
- Source: [Factory Pattern - Refactoring Guru](https://refactoring.guru/design-patterns/factory-method)

### Problem Statement
#### Part 1
- You are tasked to create a simple employee management system. 
- The system should be able to create employees of different types, such as `Manager`, `Staff`, and `Intern`. 
- Each employee type should have a different `salary` and `name`.
- Create a factory pattern to create employees of different types.
- Create unit tests for the factory pattern.

#### Part 2
- Enhance the employee management system to include a `Bonus` for each employee type.
- The `Bonus` should be calculated based on the employee's salary.
- The bonus calculation should be as follows:
  - `Manager`: 20% of the salary
  - `Staff`: 10% of the salary
  - `Intern`: no bonus (return 0)
- Create unit tests for the bonus calculation.

#### Part 3
- Create a new employee type called `Director`.
- The `Director` should have a salary of `5000`.
- The `Director` should have a bonus of `30%` of the salary.
- Create unit tests for the `Director` employee type.

### Learning Objectives
- Understand the factory pattern in Go.
- Implement the factory pattern in Go.
- Write unit tests for the factory pattern in Go.
- Enhance the factory pattern with additional functionality.
- Write unit tests for the enhanced factory pattern.
- Understand the benefits of using the factory pattern.


