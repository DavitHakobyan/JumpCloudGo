# JumpCloudGo

## Objective
The assignment is to write a small library class that can perform the following operations:

### Pre Requirements
* go version go1.13.8 or higher

#### 1. Add Action

```
    addAction (string) returning error
```
Input parameters:
``` 
    {"action":"jump", "time":100} 
    {"action":"run", "time":75}
    {"action":"jump", "time":200} 
```

#### 2. Statistics

```getStats () returning string```

Assume that an end user will be making concurrent calls to both functions.

Write a second function returns a serialized json array of the average time for each action that has 
been provided to the addAction function. Output after the 3 sample calls above would be:
```
    [ {"action":"jump", "avg":150}, {"action":"run", "avg":75} ]
```

#### Note
Assume that an end user will be making concurrent calls into all functions.

### How To run
Go to the lib folder and execute the last unit test  
`TestLibrary_AddActions_GetStats()`