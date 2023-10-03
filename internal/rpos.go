package pjtos

type PjtosEnum int

const (	
	NODE_TS PjtosEnum = iota
)

var Projetos = map[PjtosEnum]struct {
	Url string
}{	
	NODE_TS: {
		Url: "git@github.com:realfabecker/nodets.git",
	},
}
