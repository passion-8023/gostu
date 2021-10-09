package binarySearchTree

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gostu/app/services/tree"
	"gostu/app/validates"
	"gostu/app/validates/rules"
	"gostu/pkg/response"
)

func BinarySearchTree(ctx *gin.Context)  {
	var treeNode validates.Tree
	if err := ctx.ShouldBind(&treeNode); err != nil {
		//获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			//非validator.ValidationErrors类型的错误直接返回
			response.ErrorResponse(ctx, response.ValidateCheckError, err.Error())
			return
		}
		//validator.ValidationErrors类型的错误进行翻译
		response.ErrorResponse(ctx, response.ValidateCheckError, rules.Translate(errs))
		return
	}
	//var searchTree tree.BinarySearchTree
	//newTree := tree.NewBinaryTree()

	var newTree *tree.BinaryTree
	//fmt.Printf("%v\n", *newTree)
	fmt.Printf("%p\n", newTree)
	fmt.Printf("%T\n", newTree)
	newTree.Insert(33)
	fmt.Println(newTree)

	//searchTree = newTree
	//for _, num := range treeNode.Num {
	//	searchTree.Insert(num)
	//}
	//fmt.Println(searchTree.Show())
	//searchTree.Delete(50)
	//res := searchTree.Find(10)
	//fmt.Println(searchTree.Show())

	//depth := searchTree.MaxDepth()
	//response.SuccessResponse(ctx, "succ",depth)
}
