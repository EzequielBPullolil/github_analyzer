# github_analyzer
github_analyzer is a CLI developed using Golang that analyzes your public GitHub repositories and indicates if they are well-organized and follow best practices.

## Install


## Use 
To check if your profile is presentable, you call the command and pass your username as a parameter


    github_analyzer @username


output
```haskell
Repositories without readme: 10/50
Commit average: 5
Empty repositories: 10
Profile Readme: False
Score: 2/10
```
---

To list empty repositories, you can pass the --empty-repos flag.
    
    github_analyzer @username --empty-repos
    
### output
```scala
Empty repos:
- https://github.com/${username}/${repo}
- https://github.com/${username}/${repo}
....
```

---

To list without readme repos, you can pass the --readme-repos flag.
    
    github_analyzer @username --readme-repos
    
### output
```scala
Without readme repos:
- https://github.com/${username}/${repo}
- https://github.com/${username}/${repo}
....
```



