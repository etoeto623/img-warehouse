module neolong.me/img-warehouse

go 1.19

require (
	golang.design/x/clipboard v0.6.3
	golang.design/x/hotkey v0.4.1
	neolong.me/neotools/cipher v0.0.0
)

require (
	github.com/howeyc/gopass v0.0.0-20210920133722-c8aef6fb66ef // indirect
	golang.design/x/mainthread v0.3.0 // indirect
	golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550 // indirect
	golang.org/x/exp v0.0.0-20190731235908-ec7cb31e5a56 // indirect
	golang.org/x/image v0.0.0-20211028202545-6944b10bf410 // indirect
	golang.org/x/mobile v0.0.0-20210716004757-34ab1303b554 // indirect
	golang.org/x/sys v0.0.0-20210510120138-977fb7262007 // indirect
	neolong.me/neotools/common v0.0.0 // indirect
)

replace (
	neolong.me/neotools/cipher v0.0.0 => ../neotools/cipher
	neolong.me/neotools/common v0.0.0 => ../neotools/common
)
