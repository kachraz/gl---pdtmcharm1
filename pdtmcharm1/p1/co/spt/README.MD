<h1 align="center"><code> SPT </code></h1>
<h2 align="center"><i> Spinners Test </i></h2>

----

1. [W ?](#w-)
2. [Dev Notes](#dev-notes)
3. [File Descriptions](#file-descriptions)


----

# W ?

> This repo is for testing out spinners from

```
https://github.com/charmbracelet/bubbletea/blob/master/examples/spinners/README.md
```
This is the main code for seeing various spinnners , will be testing them out here

# Dev Notes

1. Animated spinner will be used as described in the [`anim.go`](./anim.go)
2. We will test out animation and running a command at the same time and see what it that looks like here , the frnning of multiple things concurrently is going to be done with go routines

# File Descriptions

> Not testing is being done with putting files directly in the main root, But ideally you will have to make sure to make subfolders with correct directory names, to actually pull it out. 

👅 | 🐈
|:--:|:--|
[`anim.go`](./anim.go) | This file contains the actual bubbles animation
[`main.go`](./main.go) | This is the entry point of the application