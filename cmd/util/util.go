package util

func MsgAppend(arr *[]string, msg string) {
	temp := []string{
		msg,
	}
	*arr = append(*arr, temp...)
}

func MsgAppendln(arr *[]string, msg string) {
	temp := []string{
		msg,
		"\n",
	}
	*arr = append(*arr, temp...)
}
