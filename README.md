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
   │       └─➤ Commit Scope: give more contect to your commit (optional)
   │
   │
   └─➤ Commit Type: build|chore|ci|docs|feat|fix|hotfix|perf|refactor|revert|style|test
```

 - **build**: change related to building the project
 - **chore**:
 - **ci**: change related to the continuous integration/continuous delivery (CI/CD) infrastructure
 - **docs**: change related to the project's documentation
 - **feat**: new feature
 - **fix**: bug fix
 - **hotfix**: critical bug fix that needs to be delivered quickly
 - **perf**: performance improvement
 - **refactor**: code change that neither fixes a bug nor adds a feature
 - **revert**: undoing a previous commit
 - **style**: code style change, such as formatting or changing naming conventions
 - **test**: adding or modifying unit tests

 #### Summary
 - Use the imerative: "add" not "added"
 - Do not captialize the first letter
 - Do not use period at the end  



---
Few exemples of good and bad commit:
```
fix: change payload of login route #OK
fix(loginCheck): change payload of login route #OK
FIX: change payload of login route #NOK fix need to be in lowercase
fix #NOK
change payload of login route #NOK
```




---

 ### Sources

 Angular contributing https://github.com/angular/angular/blob/main/CONTRIBUTING.md  
 Contentional commits https://www.conventionalcommits.org/en/v1.0.0/