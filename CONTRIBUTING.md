# Contributing to GoChatSecure

Thank you for your interest in contributing to GoChatSecure! This document provides guidelines and instructions for contributing to this project.

## Code of Conduct

By participating in this project, you agree to maintain a respectful and inclusive environment for everyone.

## How to Contribute

### Reporting Bugs

If you find a bug, please create an issue with the following information:

- A clear, descriptive title
- Steps to reproduce the issue
- Expected behavior
- Actual behavior
- Any relevant logs or screenshots
- Your environment (OS, Go version, etc.)

### Suggesting Features

We welcome feature suggestions! Please create an issue with:

- A clear, descriptive title
- Detailed description of the proposed feature
- Any relevant examples or mockups
- Explanation of why this feature would be useful

### Pull Requests

1. Fork the repository
2. Create a new branch (`git checkout -b feature/your-feature-name`)
3. Make your changes
4. Run tests (`make test`)
5. Commit your changes (`git commit -m 'Add some feature'`)
6. Push to the branch (`git push origin feature/your-feature-name`)
7. Open a Pull Request

## Development Setup

1. Clone your fork of the repository
2. Install Go (1.16 or higher)
3. Generate SSL certificates for development:
   ```bash
   make generate-certs
   ```
4. Build and run the application:
   ```bash
   make run
   ```

## Coding Standards

- Follow Go's [official style guide](https://golang.org/doc/effective_go)
- Use meaningful variable and function names
- Write comments for complex logic
- Include tests for new functionality

## Testing

- Write tests for all new features and bug fixes
- Ensure all tests pass before submitting a PR:
  ```bash
  make test
  ```
- Consider edge cases in your tests

## Commit Messages

- Use clear, descriptive commit messages
- Start with a verb in the present tense (e.g., "Add feature" not "Added feature")
- Reference issue numbers when applicable (e.g., "Fix #123: Resolve authentication bug")

## Documentation

- Update the README.md if your changes affect usage or installation
- Document new features or changed behavior
- Update API documentation if applicable

## Review Process

- All submissions require review
- Maintainers may request changes or improvements
- Be responsive to feedback and questions

## License

By contributing to GoChatSecure, you agree that your contributions will be licensed under the project's [MIT License](LICENSE).

## Questions?

If you have any questions about contributing, feel free to open an issue for clarification.

Thank you for helping improve GoChatSecure!
