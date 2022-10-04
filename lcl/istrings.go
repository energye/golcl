package lcl

type IStrings interface {
	IObject
	// 先只简单实现几个吧
	Count() int32
	S(int32) string
	Objects(int32) *TObject
}
