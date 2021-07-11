package main

/*
	注意，如果开启了 go module，那么 import 导入的路径是根据 go.mod 中 module 指定的名称来的
	比如开启了 go module，那么 import "analysis_service/dao""，它不会将 "analysis_service" 当作一个文件夹，而是把它当作一个 module 来识别
	它会找到 module 为 "analysis_service" 的项目下的 "dao" 目录导入
	如果没有开启 go module，那么它会将 "analysis_service" 当作文件夹来识别，所以如果在不需要 go module 识别的情况下要关闭，否则无法识别文件夹
*/
