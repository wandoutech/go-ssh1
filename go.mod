module github.com/ultram4rine/go-ssh1

go 1.17

require (
	github.com/dgryski/go-idea v0.0.0-20170306091226-d2fb45a411fb
	golang.org/x/crypto v1.0.0
	golang.org/x/term v0.9.0

)

require golang.org/x/sys v0.9.0 // indirect

replace golang.org/x/crypto v1.0.0 => github.com/wandoutech/crypto v1.0.0
