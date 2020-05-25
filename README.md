# Todo
***Grouping by project ?***  
***Setting date ?***  
***Sharing data ?***  

### Nope!

This is dead simple todo.  
Just add, check, and delete it.  
You can only use tags to manage todos.


## Intall 

#### golang
```terminal 
go get -u https://github.com/Zett-8/cli-todo-list
```

#### binary


## How to use

### Show list

### Add todo
```terminal
todo add 'buy milk' @chore

// output
[1] sample todo   @withTag @second
[2] add me   
[3] buy milk   @chore
```

### Check todo
```terminal
todo done 2

// output
[1] sample todo   @withTag @second
[2] add me <- displayed as grey
[3] buy milk   @chore
```

### delete todo
delete by ID
```terminal
todo delete 2
```

delete all todo
```terminal
todo delete -a (--all)
```

delete finished ones
```terminal 
todo delete -d (--done)
```

filtered by tag
```terminal
todo delete -t @work -d
```
(delete all finished todos tagged by @work)

```terminal
todo delete -t @work -a
```
(delete all todos tagged by @work)
