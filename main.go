package main

import (
	"XMLResReq/client"
)

func main() {

	xmlFT1 := &client.XMLFile{FileName: "sendXML.xml"}
	client.CreateUSR(xmlFT1)
}
