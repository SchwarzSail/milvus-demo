// Code generated by hertz generator. DO NOT EDIT.

package image

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	image "milvus-demo/biz/handler/image"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_picture := root.Group("/picture", _pictureMw()...)
		_picture.POST("/insert", append(_insertMw(), image.Insert)...)
		{
			_search := _picture.Group("/search", _searchMw()...)
			_search.GET("/image", append(_searchbyimageMw(), image.SearchByImage)...)
			_search.GET("/text", append(_searchbytextMw(), image.SearchByText)...)
		}
	}
}