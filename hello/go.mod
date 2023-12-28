module src/motd/hello

go 1.21.5

replace src/motd/message => ../message

require src/motd/message v0.0.0-00010101000000-000000000000

require (
	golang.org/x/mod v0.14.0 // indirect
	golang.org/x/tools v0.16.1 // indirect
)
