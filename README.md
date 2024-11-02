# AchoBeta Pluto System - 2024 Group G - Backend - Re-examination Project
AchoBeta Pluto System - 2024 Group G - Backend - Re-examination Project


<a href="./README_ZH.md"><b> 简体中文 </b></a>
```
   _______   ___       ____  ____  ___________  ______    
  |   __ "\ |"  |     ("  _||_ " |("     _   ")/    " \   
  (. |__) :)||  |     |   (  ) : | )__/  \\__/// ____  \  
  |:  ____/ |:  |     (:  |  | . )    \\_ /  /  /    ) :) 
  (|  /      \  |___   \\ \__/ //     |.  | (: (____/ //  
 /|__/ \    ( \_|:  \  /\\ __ //\     \:  |  \        /   
(_______)    \_______)(__________)     \__|   \"_____/    
```



## Project Functionality
Team Management System
## Project Technology Stack
1. Gin
2. Gorm
3. Mysql
4. Redis
5. zap
6. viper
7. ...
## Project Directory Structure
```
├── cmd       Program entry
├── configs   Configuration entity classes
├── global    Global variables and constants
├── initalize Initialization files
├── internal  Core business code
├── log       Log file configuration
├── pkg       Third-party common packages
├── test      Test files
├── utils     Utility classes
├── config.yaml.template    Configuration file template
├── go.mod
└── go.sum
```
## Before Development

### 1. **Set up the git hook**
```shell
git config core.hooksPath .githooks 
chmod -R -x .githooks 
```

### 2. **Read the development specifications below**

branch naming convention
we must confirm:

1. Branch naming should include a name to identify the person responsible.

2. Branch naming must clearly express what problem the branch is working on.

so branch naming must be standardizedSo branch naming must be standardized.
```bash
<type>-<name>-<description>
```
for example:
- if it is a branch to develop new functions, the naming convention is as follows
```bash
feature-<name>-<feature description>
e.g.: feature-jett-dev_log_system
```

- if is fix bugs:
```bash
bugfix-<name>-<bug name>
e.g.: bugfix-jett-login_error
```
and other types:
`hotfix`、`release`...


### commit message format
commit message should be written as clearly as possible, and each commit should only do one thing.

```bash
<type>(<scope>): <subject>

e.g.: feat: add new api
or: feat(common): add new api
```

### type

```text
# Main type
feat:      add new features
fix:       fix bug

#Special type
docs:      only document-related content has been changed
style:     changes that do not affect the meaning of the code, such as removing spaces, changing indentation, adding or deleting semicolons
build:     changes to construction tools or external dependencies, such as webpack, npm
refactor:  used when refactoring code
revert:    the message printed by executing git revert

# Do not use type yet
test:      add a test or modify an existing test
perf:      changes to improve performance
ci:        changes related to CI (Continuous Integration Service)
chore:     other modifications that do not modify src or test, such as changes to the build process or auxiliary tools
```

### subject

No period or punctuation at the end

e.g.
```bash
feat: add new feature
fix: fix a bug
```

### content
please delete useless import. You can also use the shortcut key ctrl + alt + o to automatically delete useless import by setting idea.

## **You Mast Know**
1. **Do not** submit any sensitive information, such as `api_key`, `address`, or `password` in any code.
2. You can use the configuration file `config.yaml` to store some sensitive information, but do not attempt to submit it. Each time you modify the structure of `config.yaml`, you must also update `config.yaml.template`.
3. Never use `git push --force` unless you know what you are doing.

## Todo List
- [ ] Login and Registration Module
    - [ ] Login and Logout functionality
    - [ ] Auto-login within 30 days
    - [ ] Force logout
    - [ ] Display commonly used devices
- [ ] Team Homepage Module
    - [ ] Personal Center
    - [ ] Points Overview
    - [ ] Message Module
    - [ ] Feishu Multi-dimensional Table
- [ ] Team Information Module
    - [ ] Like personal information
    - [ ] User list
    - [ ] Add new user
    - [ ] Admin team structure management
    - [ ] View and edit user information details