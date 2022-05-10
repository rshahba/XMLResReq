package client

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type XMLFile struct {
	FileName string
}

func (f *XMLFile) XMLRR() (User, error) {

	sendXMLF, _ := ioutil.ReadFile(f.FileName)

	resp, err := http.Post("someURL", "application/xml", bytes.NewBuffer(sendXMLF))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(sendXMLF))
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var newuser User

	xml.Unmarshal(response, &newuser)

	fmt.Println(string(response))
	return newuser, nil
}

func CreateUSR(u XMLReq) (XMLFile, error) {

	userResp, err := u.XMLRR()
	if err != nil {
		log.Fatal("XMLRR Error! Please try again.", err)
	}

	if (userResp.Type == "admin") && (userResp.ID == "22") {
		userResp.Web = "Access Permitted"
		userResp.Company = "CIBC"
		fmt.Println(userResp.Web)
	} else {
		fmt.Println("Acess Denied! Access is only permitted to admin ID")
	}

	file, _ := xml.MarshalIndent(userResp, "", " ")

	var r XMLFile
	r.FileName = "ReceiveXML.xml"

	_ = ioutil.WriteFile(r.FileName, file, 0644)

	resultRes := string(file)
	fmt.Println(resultRes)
	return r, nil
}
