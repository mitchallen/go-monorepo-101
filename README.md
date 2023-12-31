go-monorepo-101
==

Follow along as I workout the mechanics for my next article on:
* https://scriptable.com/

When the article is finished, I'll post a link in this README.

## Usage

* Clone and then:

```sh
go work init
go work use -r .
go list -f '{{.Dir}}' -m | xargs -L1 go mod tidy -C
go list -f '{{.Dir}}' -m | xargs -L1 go work sync -C
go list -f '{{.Dir}}' -m | xargs -L1 go test -C
```

* * *

## Requirements

To use workspace features, this project requires:
* go 18.x or higher

To use the -C flag, the commands require:
* go 20.x or higher

This project includes a Dev Container mapped to:
* go 21.x or higher

## Steps

### Step 1. Create project folder

* Create a project folder and switch to it:

```sh
mkdir -p ~/projects/golang/go-monorepo-101
cd ~/projects/golang/go-monorepo-101
```

### Step 2. Init workspaces

Initialize the workspace:

```sh
go work init
```

* View the new file that was generated:

```sh
cat go.work
```

### Step 3. Create a Package

```sh
mkdir -p ./pkg/alpha
go mod init  -C ./pkg/alpha github.com/mitchallen/go-monorepo-101/pkg/alpha
touch ./pkg/alpha/alpha.go
touch ./pkg/alpha/alpha_test.go
```

* Copy code to alpha.go and alpha_test.go and save them

* Test alpha from within it's directory:

```sh
go test -C ./pkg/alpha
```

* NOTE: Get this error:
```sh
go: no modules were found in the current workspace; see 'go help work'
```

* Update the workspace (note period at the end for current dir):
```sh
go work use -r . 
```

* Check the go.work file again in the root of the project:
```sh
cat go.work 
```
* Note that alpha is now referenced:
```sh
go 1.21.0

use ./pkg/alpha
```

* Try running the tests again (should pass with no issues):
```sh
go test -C ./pkg/alpha
```

## Step 4. Run tests from root

* From the root of the project:

* First tidy up:

```sh
go list -f '{{.Dir}}' -m | xargs -L1 go mod tidy -C
```

* Then go work sync:

```sh
go list -f '{{.Dir}}' -m | xargs -L1 go work sync -C
```

* Finally run the tests:

```sh
go list -f '{{.Dir}}' -m | xargs -L1 go test -C
```

## Step 5. Add .gitignore

* Use VSCode to create a .gitignore file
* Note that it ignores go.work

## Step 6. Create an app module

```sh
mkdir -p ./cmd/demo1
```

* Subsititute my github path for yours:

```sh
go mod init -C ./cmd/demo1 github.com/mitchallen/demo1
```

```sh
touch ./cmd/demo1/demo1.go
```

* Use an external dependency

```sh
go get -C ./cmd/demo1 github.com/mitchallen/coin
```

* Add the code for demo1.go

* Try to run the app:

```sh
go run ./cmd/demo1/demo1.go
```

* NOTE that you get this error:
```sh
demo1.go:7:2: no required module provides package github.com/mitchallen/coin; to add it:
        go get github.com/mitchallen/coin
```

* Add the module with this command (note the dot (.) to indicate the current directory)

```sh
go work use -r .
```

* View the go.work file in the root again:

```sh
cat go.work
```

* Note that the module has now been added to go.work:

```sh
go 1.21.0

use (
        ./cmd/demo1
        ./pkg/alpha
)
```

* Run the demo again:

```sh
go run ./cmd/demo1/demo1.go
```

* * *

## References

* [Go: Project structure — MonoRepo](https://blog.devops.dev/go-project-structure-monorepo-daa762ec36a2)
* [Why I use the internal folder for a Go-project](https://medium.com/@as27/internal-folder-133a4867733c)
* [Don't Put All Your Code In Internal](https://ido50.net/content/dont-put-all-your-code-in-internal)
* [cmd/go: provide a convenient way to iterate all packages in go.work workspace](https://github.com/golang/go/issues/50745)