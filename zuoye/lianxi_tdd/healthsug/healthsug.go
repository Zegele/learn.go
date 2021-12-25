package healthsug

func HeslthSug(sex string, age int, bft float64) (sug string) {
	if sex == "男" || sex == "man" { // 男
		if age >= 18 || age < 40 {
			switch {
			case bft > 0.0 && bft <= 0.1:
				sug = "偏瘦，赶紧多吃点!"
			case bft > 0.1 && bft <= 0.16:
				sug = "标准，继续保持！"
			case bft > 0.16 && bft <= 0.21:
				sug = "偏重，现在少吃点还来得及。"
			case bft > 0.21 && bft <= 0.26:
				sug = "肥胖，抓紧运动，或许还来得及。"
			case bft > 0.26:
				sug = "算了，放弃吧..."
			default:
				sug = "不在范围内，你哪里输入错了么？" //应该是err
			}
		} else if age >= 40 || age < 60 {
			switch {
			case bft > 0.0 && bft <= 0.11:
				sug = "偏瘦，赶紧多吃点!"
			case bft > 0.11 && bft <= 0.17:
				sug = "标准，继续保持！"
			case bft > 0.17 && bft <= 0.22:
				sug = "偏重，现在少吃点还来得及。"
			case bft > 0.22 && bft <= 0.27:
				sug = "肥胖，抓紧运动，或许还来得及。"
			case bft > 0.27:
				sug = "算了，放弃吧..."
			default:
				sug = "不在范围内，你哪里输入错了么？" //应该是err
			}
		} else {
			switch {
			case bft > 0.0 && bft <= 0.13:
				sug = "偏瘦，赶紧多吃点!"
			case bft > 0.13 && bft <= 0.19:
				sug = "标准，继续保持！"
			case bft > 0.19 && bft <= 0.24:
				sug = "偏重，现在少吃点还来得及。"
			case bft > 0.24 && bft <= 0.29:
				sug = "肥胖，抓紧运动，或许还来得及。"
			case bft > 0.29:
				sug = "算了，放弃吧..."
			default:
				sug = "不在范围内，你哪里输入错了么？"
			}
		}
		return sug
	} else { //person.sexval == 0.0 女
		if age >= 18 || age < 40 {
			switch {
			case bft > 0.0 && bft <= 0.2:
				sug = "偏瘦，赶紧多吃点!"
			case bft > 0.2 && bft <= 0.27:
				sug = "标准，继续保持！"
			case bft > 0.27 && bft <= 0.34:
				sug = "偏重，现在少吃点还来得及。"
			case bft > 0.34 && bft <= 0.39:
				sug = "肥胖，抓紧运动，或许还来得及。"
			case bft > 0.39:
				sug = "算了，放弃吧..."
			default:
				sug = "不在范围内，你哪里输入错了么？"
			}
		} else if age >= 40 || age < 60 {
			switch {
			case bft > 0.0 && bft <= 0.21:
				sug = "偏瘦，赶紧多吃点!"
			case bft > 0.21 && bft <= 0.28:
				sug = "标准，继续保持！"
			case bft > 0.28 && bft <= 0.35:
				sug = "偏重，现在少吃点还来得及。"
			case bft > 0.35 && bft <= 0.40:
				sug = "肥胖，抓紧运动，或许还来得及。"
			case bft > 0.4:
				sug = "算了，放弃吧..."
			default:
				sug = "不在范围内，你哪里输入错了么？"
			}
		} else {
			switch {
			case bft > 0.0 && bft <= 0.22:
				sug = "偏瘦，赶紧多吃点!"
			case bft > 0.22 && bft <= 0.29:
				sug = "标准，继续保持！"
			case bft > 0.29 && bft <= 0.36:
				sug = "偏重，现在少吃点还来得及。"
			case bft > 0.36 && bft <= 0.41:
				sug = "肥胖，抓紧运动，或许还来得及。"
			case bft > 0.41:
				sug = "算了，放弃吧..."
			default:
				sug = "不在范围内，你哪里输入错了么？"
			}
		}
		return sug
	}
}
