module neolong.me/img-warehouse

go 1.19

require neolong.me/neotools/cipher v0.0.0

require neolong.me/neotools/common v0.0.0 // indirect

replace (
	neolong.me/neotools/cipher v0.0.0 => ../neotools/cipher
	neolong.me/neotools/common v0.0.0 => ../neotools/common
)
