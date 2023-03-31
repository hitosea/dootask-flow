package service

import (
	"workflow/workflow-engine/model"

	"workflow/util"
)

// GetAllDept 所有部门列表
func GetAllDept(parentId int) (string, error) {
	datas, err := model.GetAllDept(parentId)
	if err != nil {
		return "", err
	}
	str, err := util.ToJSONStr(datas)
	if err != nil {
		return "", err
	}
	return str, nil
}

// GetUsersByDeptTree 单一获取部门下的用户列表和子部门列表
func GetUsersByDeptTree(deptID int, deptName string) (string, error) {
	//获取子部门列表
	childDept, err := model.GetDeptByParentID(deptID)
	if err != nil {
		return "", err
	}
	// 获取父级部门
	parentDept, err := model.GetDeptByID(deptID)
	if err != nil {
		return "", err
	}
	// 部门父级下的用户列表
	users, err := model.GetUsersByDept(parentDept.Name)
	if err != nil {
		return "", err
	}
	//声明构造返回树形结构 子部门：childDepartments 用户：employees
	dept := make(map[string]interface{})
	dept["childDepartments"] = childDept
	dept["employees"] = users
	dept["titleDepartments"] = parentDept.Name

	//返回查询数据
	str, err := util.ToJSONStr(dept)
	if err != nil {
		return "", err
	}
	return str, nil
}

// GetUsersByDeptAllTree 1. 根据部门ID（deptId）和名称（deptName）获取子部门列表，所有子部门下的用户 3.构造树形结构，递归处理子部门，childDepartments：子部门 employees:部门用户
func GetUsersByDeptAllTree(parentId int) (string, error) {
	//获取父部门列表
	childDept, err := model.GetDeptByParentID(parentId)
	if err != nil {
		return "", err
	}

	// 根据部门id查询出用户表department部门关联字段格式为：2,1,3所有子部门下的用户

	// 获取父部门列表中的名称放入数组deptIds,用于查询父部门下的用户
	var deptIds []int
	var deptNames []string
	for _, v := range childDept {
		deptIds = append(deptIds, v.Id)
		deptNames = append(deptNames, v.Name)
	}
	// 部门父级下的用户列表
	users, err := model.GetUsersByDeptIds(deptIds)
	if err != nil {
		return "", err
	}
	//构造树形结构 递归处理父部门childDepartments：父部门 employees:部门用户
	dept := make(map[string]interface{})
	dept["childDepartments"] = childDept
	dept["employees"] = users
	dept["titleDepartments"] = deptNames

	//返回查询数据
	str, err := util.ToJSONStr(dept)
	if err != nil {
		return "", err
	}
	return str, nil
}

// GetUserByName 根据用户名称获取用户并分页
func GetUserByName(employeeName string, pageNum int, pageSize int) (string, error) {
	datas, err := model.GetUserByName(employeeName, pageNum, pageSize)
	if err != nil {
		return "", err
	}
	//获取总数
	total, err := model.GetUserByNameCount(employeeName)
	if err != nil {
		return "", err
	}
	//构造返回数据
	var result = make(map[string]interface{})
	result["pageNum"] = pageNum
	result["pageSize"] = pageSize
	result["total"] = total
	result["list"] = datas
	str, err := util.ToJSONStr(result)
	if err != nil {
		return "", err
	}
	return str, nil
}
