## Class Diagram

```mermaid
classDiagram

class User {
    -id : int
    -name: string
    -status: int
}

class Result {
    -instances : List<~UploadInstance~>
}

class UploadInstance
 
class IUserRepository
<<interface>> IUserRepository

IUserRepository : +getUser(id int) ~User~
IUserRepository : +getUsers() List<~User~>


class InMemoryUserRepository
class MySQLUserRepository

class IResultRepository
<<interface>> IResultRepository

IResultRepository : +getResult(id int) ~Result~
IResultRepository : +getResults() List<~Result~>

class InMemoryResultRepository
class MySQLResultRepository

```

**Notes:**

User Status is defined with:
- 0 : Blacklist
- 1 : Whitelist / Active

Result Status is defined with:
- -1 : Not Processed
- 0 : Failed
- 1 : Success

Feel free to adjust the status based on your problem requirements...