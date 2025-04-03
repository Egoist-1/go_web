package gee

// 前缀路由树
type node struct {
	pattern  string  //待匹配路由
	part     string  //路由中的一部分
	children []*node //子节点
	isWild   bool    //是否精确匹配
}

//第一个匹配成功的节点,用于插入
