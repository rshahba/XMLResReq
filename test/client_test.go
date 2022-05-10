package test

import (
	"XMLResReq/client"
	"XMLResReq/mocks"
	"io/ioutil"
	"testing"

	//"XMLRR/mocks"

	"github.com/stretchr/testify/assert"

	//"log"
	"fmt"

	gomock "github.com/golang/mock/gomock"
)

// func TestDoStuffWithTestServer(t *testing.T) {

//  xmlFile, err := os.Open("xmlfile.xml")
//  if err != nil {
//      fmt.Println(err)
//  }

//  //serverStatus := "Successfully Opened XML file!"
//  defer xmlFile.Close()
//  byteValue, err := ioutil.ReadAll(xmlFile)
//  if err != nil {
//      fmt.Println(err)
//  }

//  s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//      w.Write([]byte(byteValue))
//  }))

//  defer s.Close()
//  //r := &model.NomicsResponse{Name:"Litecoin", CurrentPrice:"129.83668602", MarketCapRank:"24", AllTimeHigh:"471.89482175", CirculatingSupply:"70198121", NumExchangesTraded:"356"}

//  resp, err := http.Get(s.URL)

//  // if !reflect.DeepEqual(resp, r) {
//  //  t.Errorf("FAILED: expected %v, got %v\n",resp, r)
//  // }
//  // if !errors.Is(err, nil) {
//  //  t.Errorf("Expected error FAILED: expected %v got %v\n", nil, err)
//  // }

// }

func TestPrintOutline(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mock := mocks.NewMockXMLReq(mockCtrl)
	//r := client.XMLFile{FileName: "sendXML.xml"}
	resultWant := client.User{Type: "admin", Name: "", ID: "22", Web: ""}

	mock.EXPECT().XMLRR().Return(resultWant, nil).Times(1)

	resultGot, err := client.CreateUSR(mock)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(resultGot)
	receiveXMLTest, _ := ioutil.ReadFile("ReceiveXMLTest.xml")
	receiveXMLActual, _ := ioutil.ReadFile("ReceiveXML.xml")

	assert.Equal(t, receiveXMLActual, receiveXMLTest, "Test didnt pass!")

}
