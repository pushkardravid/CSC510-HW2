#### Table Of Contents

[Code of Conduct](#code-of-conduct)

[Styleguides](#styleguides)
  * [Git Commit Messages](#git-commit-messages)
  * [Ruby Styleguide](#ruby-styleguide)
  * [Julia Styleguide](#julia-styleguide)
  * [Golang Styleguide](#golang-styleguide)

## Code of Conduct

All developers contributing to this project abide by the [Code of Conduct](CODE_OF_CONDUCT.md) defined for this project. Violations should be reported at  [pushkarsdravid@gmail.com](mailto:pushkarsdravid@gmail.com).
Please follow these steps to have your contribution considered by the maintainers:


## Styleguides

### Git Commit Messages

* Use the present tense ("Add feature" not "Added feature")
* Use the imperative mood ("Move cursor to..." not "Moves cursor to...")
* Limit the first line to 72 characters or less
* Reference issues and pull requests liberally after the first line
* When only changing documentation, include `[ci skip]` in the commit title
* Attach issue id in the commit message for easier tracking


### Ruby Styleguide

All Ruby must adhere to [Ruby Standard Style](https://rubystyle.guide/).

* Consistency with this style guide is important. Consistency within a project is more important. Consistency within one class or method is the most important.
* Limit lines to 80 characters.
* Avoid trailing whitespace.
* Use one expression per line.
* No space inside range literals.
* Name identifiers in English.
* Use CamelCase for classes and modules. (Keep acronyms like HTTP, RFC, XML uppercase).
* Use snake_case for naming files, e.g. hello_world.rb.
* Omit both the outer braces and parentheses for methods that are part of an internal DSL.

### Julia Styleguide

All Julia must adhere to [Julia Standard Style](https://docs.julialang.org/en/v1/manual/style-guide/).

* Write functions, not just scripts
* Avoid writing overly-specific types (Code should be as generic as possible)
* Handle excess argument diversity in the caller
* Append ! to names of functions that modify their arguments
* Use naming conventions consistent with Julia base/
* Don't use unnecessary static parameters


### Golang Styleguide

All Golang must adhere to [Golang Standard Style](https://golang.org/doc/effective_go.html/).

* Try to keep the normal code path at a minimal indentation, and indent the error handling, dealing with it first.
* Do not define interfaces on the implementor side of an API "for mocking"; instead, design the API so that it can be tested using the public API of the real implementation.
* There is no rigid line length limit in Go code, but avoid uncomfortably long lines.
* See https://golang.org/doc/effective_go.html#mixed-caps. This applies even when it breaks conventions in other languages. For example an unexported constant is maxLength not MaxLength or MAX_LENGTH.
* Don't pass pointers as function arguments just to save a few bytes. If a function refers to its argument x only as *x throughout, then the argument shouldn't be a pointer.