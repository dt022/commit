# Commit

Golang tool for easy commit all & push with message 

## Installation
`go install -v github.com/dt022/commit`

## Usage
`commit -h`

```
Usage: commit [options]
Options:
  -m string
        Commit with message
  -p string
        Commit and Push with message
```

## Example
Local commit without pushing to remote
`commit -m 'first commit'`

Commit and push to remote
`commit -p 'fixing bug'`
