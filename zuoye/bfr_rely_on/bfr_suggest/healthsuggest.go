package bfr1

func BfrSuggest(bfr float64, sexval float64, age int) (sug string) {
	if sexval == 1.0 { // 男
		if age >= 18 || age < 40 {
			switch {
			case bfr > 0.0 && bfr <= 0.1:
				sug = "偏瘦，赶紧多吃点!"
			case bfr > 0.1 && bfr <= 0.16:
				sug = "标准，继续保持！"
			case bfr > 0.16 && bfr <= 0.21:
				sug = "偏重，现在少吃点还来得及。"
			case bfr > 0.21 && bfr <= 0.26:
				sug = "肥胖，抓紧运动，或许还来得及。"
			case bfr > 0.26:
				sug = "算了，放弃吧..."
			default:
				sug = "不在范围内，你哪里输入错了么？" //应该是err
			}
		} else if age >= 40 || age < 60 {
			switch {
			case bfr > 0.0 && bfr <= 0.11:
				sug = "偏瘦，赶紧多吃点!"
			case bfr > 0.11 && bfr <= 0.17:
				sug = "标准，继续保持！"
			case bfr > 0.17 && bfr <= 0.22:
				sug = "偏重，现在少吃点还来得及。"
			case bfr > 0.22 && bfr <= 0.27:
				sug = "肥胖，抓紧运动，或许还来得及。"
			case bfr > 0.27:
				sug = "算了，放弃吧..."
			default:
				sug = "不在范围内，你哪里输入错了么？" //应该是err
			}
		} else {
			switch {
			case bfr > 0.0 && bfr <= 0.13:
				sug = "偏瘦，赶紧多吃点!"
			case bfr > 0.13 && bfr <= 0.19:
				sug = "标准，继续保持！"
			case bfr > 0.19 && bfr <= 0.24:
				sug = "偏重，现在少吃点还来得及。"
			case bfr > 0.24 && bfr <= 0.29:
				sug = "肥胖，抓紧运动，或许还来得及。"
			case bfr > 0.29:
				sug = "算了，放弃吧..."
			default:
				sug = "不在范围内，你哪里输入错了么？"
			}
		}
		return sug
	} else { //client.sexval == 0.0 女
		if age >= 18 || age < 40 {
			switch {
			case bfr > 0.0 && bfr <= 0.2:
				sug = "偏瘦，赶紧多吃点!"
			case bfr > 0.2 && bfr <= 0.27:
				sug = "标准，继续保持！"
			case bfr > 0.27 && bfr <= 0.34:
				sug = "偏重，现在少吃点还来得及。"
			case bfr > 0.34 && bfr <= 0.39:
				sug = "肥胖，抓紧运动，或许还来得及。"
			case bfr > 0.39:
				sug = "算了，放弃吧..."
			default:
				sug = "不在范围内，你哪里输入错了么？"
			}
		} else if age >= 40 || age < 60 {
			switch {
			case bfr > 0.0 && bfr <= 0.21:
				sug = "偏瘦，赶紧多吃点!"
			case bfr > 0.21 && bfr <= 0.28:
				sug = "标准，继续保持！"
			case bfr > 0.28 && bfr <= 0.35:
				sug = "偏重，现在少吃点还来得及。"
			case bfr > 0.35 && bfr <= 0.40:
				sug = "肥胖，抓紧运动，或许还来得及。"
			case bfr > 0.4:
				sug = "算了，放弃吧..."
			default:
				sug = "不在范围内，你哪里输入错了么？"
			}
		} else {
			switch {
			case bfr > 0.0 && bfr <= 0.22:
				sug = "偏瘦，赶紧多吃点!"
			case bfr > 0.22 && bfr <= 0.29:
				sug = "标准，继续保持！"
			case bfr > 0.29 && bfr <= 0.36:
				sug = "偏重，现在少吃点还来得及。"
			case bfr > 0.36 && bfr <= 0.41:
				sug = "肥胖，抓紧运动，或许还来得及。"
			case bfr > 0.41:
				sug = "算了，放弃吧..."
			default:
				sug = "不在范围内，你哪里输入错了么？"
			}
		}
		return sug
	}
}
