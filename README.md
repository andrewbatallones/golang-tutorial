# Go Lang Tutorial
This is just a data dump of personal notes as I go through these examples and readings
- https://go.dev/doc/code
- https://gobyexample.com/
- https://www.youtube.com/watch?v=YS4e4q9oBaU&t=3392s
- https://go.dev/doc/effective_go

### Other resources and readings to look into
- https://go.dev/blog/slices-intro


### Questions to ask
- How to import external packages and create your packages
    - A crucial part of importing packages is having the mod.go (think of it as a package.json or Ruby Gemfile). It will keep track of information on your package. You can then use that to have nested packages and import them to other packages as needed.
- Find out what type assertion is

### Other notes
- Capitalize variables are outside the package scope, but that is also true for functions. This is one way to write external functionality
- look into godoc package https://pkg.go.dev/golang.org/x/tools/cmd/godoc for web serving