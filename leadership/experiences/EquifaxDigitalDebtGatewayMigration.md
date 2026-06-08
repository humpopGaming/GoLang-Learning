# Equifax — Digital Debt Gateway Migration (Technical Leadership)

<aside>
🎯

# How I Can Reuse This Story

This ONE example can answer:

- "Give me an example of how you've improved code quality or engineering practices"
- "How do you balance technical rigour with delivery pressure?"
- "Tell me about a time you had to make a significant technical decision"
- "How do you approach CI/CD adoption?"
- "How do you develop engineers who are earlier in their careers?"
- "Tell me about a time a project was at risk. What did you do?" *(reframe: risk was identified and neutralised proactively)*
</aside>

### S — Situation

At Equifax I inherited a product called Digital Debt Gateway — a UI written in AngularJS backed by C#. The developer who had built it had left the business. Google was sunsetting AngularJS, making Angular the new standard, and we had a hard deadline of two months to complete the migration.

The UI codebase had been treated as a second-class citizen. There were no unit tests, no consistency in the code, and it was difficult to read and understand. The team was small: myself, two developers whose background was predominantly C#, and one tester — none of us had deep Angular experience going in.

---

### T — Task

I was the tech lead, accountable for proposing and owning the migration approach. It was my responsibility to determine how we tackled the migration, what engineering standards we introduced, and how the team worked. I also made the call to introduce unit tests — this was not asked of me, and it drew pushback from my manager who was concerned about hitting the deadline. I made the case that getting it right on the first pass was worth the investment, and my track record gave me the credibility to win that argument.

---

### A — Actions

**1. Chose the right migration strategy.** I assessed two options: run Angular and AngularJS side by side, or rebuild in isolation and release on completion. The hybrid approach had a significant technical problem — managing state across Angular and AngularJS simultaneously was complex and risky. I presented both options to the product owner and my manager with a clear recommendation: rebuild clean. Despite the less agile nature of that approach, given there was no flexibility in the deadline, a clean rebuild was faster and safer than trying to migrate an already inconsistent codebase incrementally.

**2. Built the foundation first.** I broke the application into modules to enable parallel development and reduce blockers. Before the team could work independently, certain foundational pieces had to exist first — so I built those myself. This also gave the other developers structured time for Angular self-learning while I set up the core. The foundation served as a working example of the module structure, unit test implementation using Karma and Jasmine, and linting standards — something the team could read and follow rather than figure out from scratch.

**3. Coached the team through the approach.** Once the foundation was in place I walked both developers through it, answered questions, and identified gaps in their knowledge. We held multiple check-ins each day to stay aligned, unblock problems quickly, and make sure no one was heading in the wrong direction.

**4. Enforced standards through the pipeline.** I updated the CI/CD pipeline so that any code failing to achieve 80% test coverage would break the build — tests were not optional. I also found that the master branch had been left unprotected. With multiple developers working in parallel this was a clear risk, so I protected it immediately as part of the pipeline work.

---

### R — Result

The project was delivered on time within the two-month deadline. We went from zero unit tests in the original codebase to 80% coverage enforced by the pipeline on every merge. The new codebase was modular, consistent, and built to Angular and Google best practices — a significant improvement on what we had inherited.

The product owner fed back that customers found the new UI a significantly better experience. On the people side, one of the developers on the project used the Angular skills he built during the migration to land a more senior role at another company, and he later went on to found his own full-stack development consultancy.

---

### L — Learning

This was the first time I had properly led and guided other developers as the recognised technical expert on a project. In hindsight, I took on too much personally — I under-delegated work that others could have owned, which put unnecessary pressure on me. Since then I've been more deliberate about where I'm best placed to contribute as a project progresses: not forcing work onto others, but actively thinking about which pieces need my direct involvement and which should be handed off.

On the technical side, I developed a much clearer understanding of how to test Angular applications correctly — specifically the importance of stubbing dependencies properly. Without the right stubs, a single change can create cascading failures across the test suite that are difficult to trace.
