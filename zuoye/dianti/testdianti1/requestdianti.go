package testdianti1

import "fmt"

//电梯外有两个按钮，用来请求电梯。电梯可以知道你是上，还是下。
//电梯内可以知道你要到几楼，先运行原来的方向，没有该方向，然后再转向。

type DianTiXiang struct {
}

func (D DianTiXiang) requestDianTi(reqPeople int) (moveHow int, err error) {
	if reqPeople == 0 {
		return 0, nil
	} else if reqPeople < 0 {
		return reqPeople, fmt.Errorf("不能为负数")
	}
	return 1, nil
}

func (D DianTiXiang) moveDianTi(moveHow int) (moveWhere string, err error) {
	if moveHow == 0 {
		return "没人来，电梯不动。", nil
	}
	return "", fmt.Errorf("未知")
}
