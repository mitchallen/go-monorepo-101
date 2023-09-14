go-monorepo-101
==

## Usage

## Requirements

To use workspace featurers, this project requires:
* go 18.x or higher

## Steps

### Step 1. Create project folder

* Create a project folder and switch to it:

```sh
mkdir -p ~/projects/golang/go-monorepo-101
cd ~/projects/golang/go-monorepo-101
```

### Step 2. Create a library

```sh
mkdir -p ./lib/alpha
cd ./lib/alpha
go mod init alpha
touch alpha.go
touch alpha_test.go
```

* Copy code to alpha.go and alpha_test.go and save them

* Test alpha from within it's directory:

```sh
go test
```

## Step 3. Setup workspaces

* cd to the project root

```sh
cd ../..
```

* Init the workspaces

```sh
go work init
```

* View the newly created go.work file:

```sh
cat go.work
```

* Add the alpha folder to the workspace

* Notice the period at the end, referring to the current directory

```sh
go work use -r .
```

See what's changed:

```sh
cat go.work
```

## Step 4. Run tests from root

* First tidy up:

```sh
go list -f '{{.Dir}}' -m | xargs -L1 go mod tidy -C
```

* Them work sync:

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

* * *

## References

* [Go: Project structure — MonoRepo](https://blog.devops.dev/go-project-structure-monorepo-daa762ec36a2)
* [Why I use the internal folder for a Go-project](https://medium.com/@as27/internal-folder-133a4867733c)
* [Don't Put All Your Code In Internal](https://ido50.net/content/dont-put-all-your-code-in-internal)
* [cmd/go: provide a convenient way to iterate all packages in go.work workspace](https://github.com/golang/go/issues/50745)