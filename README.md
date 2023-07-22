$ curl "http:localhost:8080/analysis?duration=5s"
## Expected ouput

### Readme

Write your README as if it was for a production service. Include the following items:

* Description of your solution.
* Reasoning behind your technical choices, including architectural.
* Trade-offs you might have made, anything you left out, or what you might do differently if you were to spend additional time on the project.
* And of course, how to run your project. Including expected dependencies and command examples showing us how to install and run your project.

### Technical solution

In order to complete this challenge you can use the language you want.

You are only allowed to use the standard library (stdlib) of the
language in order to complete the assignment except for the HTTP server
/ request handling (i.e. you can use flask or django if you use python or rails / sinatra in ruby).

### How we review

Your application will be reviewed by at least three of our engineers. We do take into consideration your experience level.

We value quality over feature-completeness. It is fine to leave things aside provided you call them out in your project's `README`. The goal of this code sample is to help us identify what you consider production-ready code. You should consider this code ready for final review with your colleague, i.e. this would be the last step before deploying to production.

The aspects of your code we will assess include:

* **Architecture**: How clean is the separation between the different components? How well are the different classes separated?
* **Clarity**: Does the README clearly and concisely explain the problem and solution? Are technical trade-offs explained?
* **Correctness**: Does the application do what was asked? If there is anything missing, does the README explain why it is missing?
* **Code quality**: Is the code simple, easy to understand, and maintainable? Are there any code smells or other red flags? Is the coding style consistent with the language's guidelines? Is it consistent throughout the codebase?
* **Technical choices**: do choices of libraries, architecture etc. seem appropriate for the chosen application?

Bonus point (those items are optional):

* **Testing quality**: How extensive is your testing coverage? Is your testing approach pragmatic?
* **Efficiency**: What is the footprint of your program? Both in terms of CPU and memory usage
* **Scalability**: Will technical choices scale well (increase both in terms of post throughput and number of concurrent client) ? If not, is there a discussion of those choices in the README?