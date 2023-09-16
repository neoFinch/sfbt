# sfbt
Search files by type and you don't have to remember any flags. 😁

![cli-tool](https://github.com/neoFinch/sfbt/assets/14983412/10248d96-b051-4270-b30a-8a31d23bb079)


## A cli tool for easy use

### How to build
Navigate to your projects root directory and run below command
```
go build
```

### Steps to run the command from terminal
After the above `go build` command you will get a file named `sf`.

**Step 1**: Create a new folder in home directory.

example: `mkdir custom-bin`

**Step 2**: Move your `sf` file from the project folder to this `custom-bin` folder.

`mv sf /home/<your_user>/custom-bin`

**Step 3**: Add your custom-bin folder path to your `.zshrc` or `.bashrc` file depending on your shell.

`export PATH=$PATH:/home/<your_user>/custom-bin`

**Step 4**: Source your .zshrc file

`source .zshrc`
