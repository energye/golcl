package lcl

type EMFS interface {
	LoadFromFSFile(Filename string) error
}
