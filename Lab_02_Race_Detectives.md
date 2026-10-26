# Lab 2: Race Detectives

**CS 411: Computing Paradigms - Spring 2027**  
**Week 2: January 18-24, 2027**  
**Due:** January 24, 2027  
**Value:** 100 points

## Purpose

A shared counter loses updates only under particular thread schedules. Reproduce the race, explain one losing interleaving, then protect the complete read-modify-write invariant and verify the repair with Go tooling.

## Learning objectives

- Trace a read-modify-write race
- Select a correct mutex scope
- Use the Go race detector as evidence

## Before you begin

1. Review the Week 2 slide deck and assigned VoiceThread or reading.
2. Create your own repository from the Spring 2027 template linked in Blackboard. Its `counter.go` is deliberately racy; `counter_test.go` is a public verification harness.
3. Run `go test ./...` and `go test -race ./...` before editing. Save the initial output in `Lab2_Terminal.txt`. A lost-total failure or a reported race is expected.

## Part A Trace the failure (25 points)

Explain the mechanism before changing code.

- In `Lab2_Analysis.md` or PDF, draw a two-worker interleaving using the operations read, add, and write. Start with zero and show a final total of one where two was required.
- Identify the shared state, conflicting accesses, and violated invariant. Explain why a sequential test can pass while the concurrent test fails.

## Part B Repair the counter (45 points)

Implement the smallest correct synchronization boundary.

- Complete both TODOs in `counter.go` using the provided `sync.Mutex`.
- Protect the read, update, and write as one critical section; make `Value` return a synchronized snapshot.
- Preserve the `Counter`, `Add`, and `Value` API. Keep the lock local to the counter and avoid a package-wide lock.

## Part C Prove the repair (30 points)

Use repeatable evidence rather than one successful run.

- Run `go test ./...`, `go test -race ./...`, and `go test -race -count=20 ./...`. Save complete commands and output in `Lab2_Terminal.txt`, including the failing starter run and passing repaired run.
- In the analysis, compare the final code with your Part A interleaving. Explain why a clean race-detector result and stable totals support different claims, and why neither finite run proves every possible schedule.

## Verification

- `go test -race ./...` reports no data race and exits successfully.
- `go test -race -count=20 ./...` returns the expected total on repeated runs.
- The analysis names the complete protected invariant and includes the failing interleaving.

## What to submit

Submit your own repository URL and:

1. Your repository URL (created from the Spring 2027 template).
2. `counter.go` and the unchanged public `counter_test.go` in that repository.
3. `Lab2_Analysis.pdf` or `Lab2_Analysis.md`.
4. `Lab2_Terminal.txt`.

## Rubric

| Criterion | Points |
| --- | ---: |
| Failure trace and invariant | 25 |
| Correct mutex scope | 35 |
| Code quality and API preservation | 10 |
| Race-detector and stress evidence | 20 |
| Explanation | 10 |
| **Total** | **100** |

## AI usage and academic integrity

Add an **AI Usage** section to the required analysis file. Name each tool and list the prompts or interactions that affected the work. If you did not use AI, write `AI Usage: None.` You must be able to explain and defend every submitted design and line of code.

## Late work

Late work may be accepted for up to three calendar days with a 25 percent deduction per day, following the course syllabus.
