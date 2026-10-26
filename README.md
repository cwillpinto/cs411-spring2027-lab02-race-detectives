# CS411 Spring 2027 · Lab 2: Race Detectives

This student template intentionally contains a data race in `counter.go`. Read the included handout before editing. Do not treat a single passing test run as proof of correctness.

```sh
go test ./...
go test -race ./...
go test -race -count=20 ./...
```

Keep the `Counter`, `Add`, and `Value` API. Complete the synchronization TODOs, then submit the evidence and analysis described in the handout. The public tests are representative, not exhaustive.

## Work in your own GitHub repository

The instructor starter is an upstream source. Keep your completed work in a repository under your own GitHub account, then submit the repository URL and the evidence listed in the assignment handout through Blackboard.

If Blackboard links a template repository, choose **Use this template**, create the new repository under your account, and clone that new repository. Check `git remote -v` before pushing; `origin` must be your repository.

For a direct clone, first create an empty repository under your account. Replace the placeholders below with the starter link supplied in Blackboard and your own empty repository URL:

```sh
git clone <instructor-starter-url>
cd <cloned-directory>
git remote rename origin upstream
git remote add origin <your-empty-repository-url>
git remote -v
git push -u origin HEAD
```

Commit and push your completed work to your own `origin`. Do not push to the instructor's upstream repository. Use the exact submission files and verification commands in the assignment handout.
