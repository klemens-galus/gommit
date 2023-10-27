# Code of contuct 

## Commit Message Format
Each commit message consists of a **header**, **body** and **footer**.
```
<header>
<BLANK LINE>
<body>
<BLANK LINE>
<footer>
```
`header` need to be conform (look section : [here](#commit-message-header))

`body` give a short summary of the code changes (optional)

`footer` (optional)

*Exemple*
```
feat: now front show date and time

Users can now see date and time at the 
botom left of the website

Refs: #125
```

## Commit Message Header
```
<type>(<scope>): <short summary>
   │       │             │
   │       │             └─➤ A short summary of the code changes
   │       │                    Not capital letters, No period at the end
   │       │
   │       └─➤ Commit Scope: 
   │
   │
   └─➤ Commit Type: build|chore|ci|docs|feat|fix|hotfix|perf|refactor|revert|style|test
```
