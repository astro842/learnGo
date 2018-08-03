package mock

//接口实现者
//接口的实现是隐式的 不需要指明 只需要实现接口的方法
//这个 struct 名字随便起
type AAA struct {
	Contents string
}
//实现 Retriever interface
func (r AAA) Get(url string) string {
	return r.Contents
}
