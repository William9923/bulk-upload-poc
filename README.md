# Wordle The Developer
<p align="center">
  <img src="docs/img/gopher-hooks.jpeg" />
</p>

## Why?
To increase the productivity of a developer, I would like to simplify the init process when trying to develop an application. 

One way to do that are by using template (boilerplate). This template provides many **git hooks**. Git hooks are scripts that Git executes before or after events such as: commit, push, and receive. Git hooks are a built-in feature - no need to download anything. Git hooks are run locally.

## What?
To make it easier to use, I give a concrete example of application that use the capabilities of command line interface, http and cron (scheduler)... to play a simple wordle game. This game simply replicates the current trending wordle game, but would be useful for future me when trying to build new application...

## Hooks Overview
In this repo, there are 4 hooks that I implemented for Golang (could be used for other programming languange too):
- pre-commit: Used to apply gofmt & goimports before committing the code.
- commit-msg: Used to check the format of commit message before commit. Will fail if not suitable from predetermined template
- pre-push: Used to check with unit test before pushing items into remote repo. But, will pass if commit message contains WIP (work in progress). Also, prevent from accidental push to master.
- post-checkout: Used to perform dependencies check (vendor, etc) when switching branch.

## Prerequisites
- Golang minimum v1.12 (https://golang.org/doc/install)
- Go Modules (https://blog.golang.org/using-go-modules)

## How to Run 
1. Clone the repository
```bash
git clone git@github.com:William9923/hook-the-git.git
```
2. Setup the hooks
```bash
sh ./setup.sh
```

