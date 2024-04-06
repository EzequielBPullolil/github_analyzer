# github_analyzer
github_analyzer is a CLI developed using Golang that analyzes your public GitHub repositories and indicates if they are well-organized and follow best practices.


## Install


## Use 
To check if your profile is presentable, you call the command and pass your username as a parameter
Please note that the app only scans the master branch.

    github_analyzer @username


output
```haskell
Repositories without readme: 10/50
Commit average: 5
Empty repositories: 10
Profile Readme: False
```
---

To list empty repositories, you can pass the --empty-repos flag.
    
    github_analyzer @username --empty-repos
    
### output
```scala
Finding empty repos
Finding public repos.....
Public repositories found, quantity: 27
Finding public empty repos...
WARNING Repositories are considered empty if the number of commits in the main branch is less than 5
- https://github.com/https://avatars.githubusercontent.com/u/62083227?v=4
- https://github.com//EzequielBPullolil/github_analyzer
- https://github.com//EzequielBPullolil/cipher_view
- https://github.com//EzequielBPullolil/pure_scriptor
- https://github.com//EzequielBPullolil/backend_fast_reserve
- https://github.com//EzequielBPullolil/switf_talk
- https://github.com//EzequielBPullolil/SocialNet_Microservices
....
```

---

To list without readme repos, you can pass the --readme-repos flag.
    
    github_analyzer @username --readme-repos
    
### output
```scala
Searching repositories without README
Finding public repos.....
Public repositories found, quantity: 27
Searching repositories without README...
- https://github.com/https://avatars.githubusercontent.com/u/62083227?v=4
- https://github.com//EzequielBPullolil/cipher_view
- https://github.com//EzequielBPullolil/data_structures
- https://github.com//EzequielBPullolil/backend_fast_reserve
- https://github.com//EzequielBPullolil/price_sync_system
...
```



