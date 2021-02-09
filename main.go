package wechatwork

import "os"

func main() {
	m := []string{
		"callback",
		"user",
		"department",
		"tag",
		"batch",
		"linkedcorp",
		"externalcontact",
		"agent",
		"message",
		"appchat",
		"media",
		"checkin",
		"oa",
		"journal",
		"meetingroom",
		"calendar",
		"meeting",
		"living",
		"wedrive",
		"dial",
		"payment",
		"corpgroup",
		"invoice",
		"school",
		"report",
	}
	for _, v := range m {
		os.Mkdir(v, 0755)
		f1, _ := os.Create(v + "/" + v + "_iface.go")
		f1.Write([]byte("package " + v + "\n" + "type " + v + "IFace interface {}"))
		f1.Close()

		f2, _ := os.Create(v + "/" + v + "_impl.go")
		f2.Write([]byte("package " + v + "\n" + "type " + v + " struct{} \nfunc New" + v + "()*" + v + "{\n\treturn &" + v + "{}\n}"))
		f2.Close()

		f3, _ := os.Create("request/" + v + ".go")
		f3.Write([]byte("package " + "request"))
		f3.Close()
		f4, _ := os.Create("response/" + v + ".go")
		f4.Write([]byte("package " + "response"))
		f4.Close()
	}

}
