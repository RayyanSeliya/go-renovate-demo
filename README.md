# Go Renovate Demo

This is a simple Go program that demonstrates Renovate usage for dependency updates.

## What it does

Prints a message and imports `github.com/gorilla/websocket` (v1.4.1).

## How to run

`go run main.go`


## Vulnerable Dependency

- **Dependency:** `github.com/gorilla/websocket` v1.4.1
- **CVE:** [CVE-2020-27813](https://nvd.nist.gov/vuln/detail/CVE-2020-27813)


## 🚀 Renovate Security Update Demo

After setting up [Renovate](https://github.com/marketplace/renovate) in this repository, it automatically checked our Go dependencies and found that we were using a risky version of [`github.com/gorilla/websocket`](https://github.com/gorilla/websocket) (`v1.4.1`). This version has a known security problem ([CVE-2020-27813](https://nvd.nist.gov/vuln/detail/CVE-2020-27813)).

Here’s what happened next:

- **Renovate opened a Pull Request:**  
  Renovate created [PR #1](https://github.com/RayyanSeliya/go-renovate-demo/pull/1) to update `github.com/gorilla/websocket` from the unsafe version (`v1.4.1`) to the latest safe version (`v1.5.3`).  
  The pull request clearly shows what will change and why it’s safer.

- **Renovate created a Dependency Dashboard:**  
  Renovate also opened [Issue #2: Dependency Dashboard](https://github.com/RayyanSeliya/go-renovate-demo/issues/2).  
  This dashboard is a place where you can see all your dependencies, which ones need updates, and manage them easily.

### What does this mean for our project?

- **Automatic Security:** Renovate helps keep our project safe by warning us about risky dependencies and suggesting updates.
- **Transparency:** Every change is tracked through pull requests and the dashboard, so it’s easy to review and approve updates.
- **Easy Management:** The dashboard makes it simple to keep an eye on all our dependencies in one place.

---

**In summary:**  
Renovate quickly found a real security issue in our code and helped us fix it with almost no effort. This keeps our project safer and saves us time!

---
