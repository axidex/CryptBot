package models

type Des3 struct {
	L  *int `form:"l" binding:"required,gte=0"`
	R  *int `form:"r" binding:"required,gte=0"`
	K1 *int `form:"k1" binding:"required,gte=0"`
	K2 *int `form:"k2" binding:"required,gte=0"`
	K3 *int `form:"k3" binding:"required,gte=0"`
}

type A5 struct {
	A    string `form:"a" binding:"required"`
	B    string `form:"b" binding:"required"`
	C    string `form:"c" binding:"required"`
	Text string `form:"text" binding:"required"`
}

type Aes struct {
	Text string `form:"text" binding:"required"`
	Key  string `form:"key" binding:"required"`
}

type Blowfish struct {
	L  *int `form:"l" binding:"required,gte=0"`
	R  *int `form:"r" binding:"required,gte=0"`
	K1 *int `form:"k1" binding:"required,gte=0"`
	K2 *int `form:"k2" binding:"required,gte=0"`
	K3 *int `form:"k3" binding:"required,gte=0"`
	K4 *int `form:"k4" binding:"required,gte=0"`
	K5 *int `form:"k5" binding:"required,gte=0"`
}

type Des struct {
	Text string `form:"text" binding:"required"`
	Key  string `form:"key" binding:"required"`
}

type Diffie struct {
	N *int `form:"n" binding:"required,gte=0"`
	Q *int `form:"q" binding:"required,gte=0"`
	X *int `form:"x" binding:"required,gte=0"`
	Y *int `form:"y" binding:"required,gte=0"`
}

type Elgamal struct {
	P *int `form:"p" binding:"required,gte=0"`
	G *int `form:"g" binding:"required,gte=0"`
	K *int `form:"k" binding:"required,gte=0"`
	X *int `form:"x" binding:"required,gte=0"`
	M *int `form:"m" binding:"required,gte=0"`
}

type Enigma struct {
	Rot *int   `form:"rot" binding:"required,gte=0"`
	Ref string `form:"ref" binding:"required"`
	Pp  string `form:"pp" binding:"required"`
	Str string `form:"str" binding:"required"`
}

type Feistel struct {
	Data string `form:"data" binding:"required"`
	Keys string `form:"keys" binding:"required"`
}

type Hash struct {
	Word string `form:"word" binding:"required"`
	H0   *int   `form:"h0" binding:"required,gte=0"`
	P    *int   `form:"p" binding:"required,gte=0"`
}

type InvMix struct {
	Char string `form:"char" binding:"required"`
}

type Md5 struct {
	A *int `form:"a" binding:"required,gte=0"`
	B *int `form:"b" binding:"required,gte=0"`
	C *int `form:"c" binding:"required,gte=0"`
	D *int `form:"d" binding:"required,gte=0"`
	M *int `form:"m" binding:"required,gte=0"`
	K *int `form:"k" binding:"required,gte=0"`
}

type Rsa struct {
	P *int `form:"p" binding:"required,gte=0"`
	Q *int `form:"q" binding:"required,gte=0"`
	E *int `form:"e" binding:"required,gte=0"`
	X *int `form:"x" binding:"required,gte=0"`
}

type SBlock struct {
	A  *int `form:"a" binding:"required,gte=0"`
	B  *int `form:"b" binding:"required,gte=0"`
	C  *int `form:"c" binding:"required,gte=0"`
	Z0 *int `form:"z0" binding:"required,gte=0"`
	N  *int `form:"n" binding:"required,gte=0"`
}
