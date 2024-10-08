package conf

var GroupMap map[string]string

func init() {
	GroupMap = map[string]string{
		"del-zhenwu/easy_http":      "开源测试",
		"open-mmlab/mmediting":      "Low-level Vision",
		"open-mmlab/mmocr":          "MMOCR-OpenMMLab分部",
		"open-mmlab/mmdetection":    "2.0 适配-Det",
		"open-mmlab/mmflow":         "接收 mmflow/mmseg Github 消息",
		"open-mmlab/mmsegmentation": "接收 mmflow/mmseg Github 消息",
		// "open-mmlab/mmdetection3d": "",
	}
}
