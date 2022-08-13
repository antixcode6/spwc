# spwc
spwc (Simple PassWord Cache) is a cli utility written in Go that aims to provide similar functionality to [pass(1)](https://linux.die.net/man/1/pass). 

## What it can do
Right now spwc can create the .cache file, insert a password into the .cache file using a passphrase to encrypt it, and read this password from the .cache file and decrypt it with the same passphrase.



## How to use
I'll write this later when it can do more stuff. However, if (you) decide to try it run go build from the cmd directory and run `cmd -h` to get some ideas.


## Planned improvements
Finishing the basic features like deleting a password, updating one, , generating passwords, and listing multiple passwords. 

Since I picked JSON for the .cache file format I'll also need to be able to merge the JSON of one password into another password JSON so it will just be on large nested JSON struct ex. 
```
password1
{
    "password": "test",
    "description": "test"
}

password2
{
    "password": "test1",
    "description": "test1"
}

merged
{
    "password": "test",
    "description": "test"

    "password": "test1",
    "description": "test1"
}

```


This should probably just be reworked so that it is nested to find passwords easier.
```
{
  "1": {
    "password": "test",
    "description": "test"
  },
  "2": {
    "password": "test1",
    "description": "test1"
  }
}
```

This could also just be a KV pair of password and description to keyword search via description.