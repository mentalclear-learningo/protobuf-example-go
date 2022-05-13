package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"protobuf-example-go/src/simple"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {
	sm := doSimple()

	// For read and write files demo
	readAndWriteDemo(sm)

	// For JSON example
	jsonDemo(sm)
}

func jsonDemo(sm proto.Message) {
	smAsString := toJSON(sm)
	fmt.Println("JSON message:", smAsString)

	testJSONmsg := `{ 
		"id":54321, 
		"isSimple":false, 
		"name":"Test message!", 
		"sampleList":[54, 2, 32, 11]
	}`

	sm2 := &simple.SimpleMessage{}
	sm3 := &simple.SimpleMessage{}
	fromJSON(smAsString, sm2)
	fromJSON(testJSONmsg, sm3)
	fmt.Println("Unmarashaled message 1:", sm2)
	fmt.Println("Unmarashaled message 2:", sm3)
}

// Marshal protobuf message to JSON format
func toJSON(pb proto.Message) string {
	out, err := protojson.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
		return ""
	}
	return string(out)
}

func fromJSON(in string, pb proto.Message) error {
	err := protojson.Unmarshal([]byte(in), pb)
	if err != nil {
		log.Fatalln("Error unmarshlling json", err)
		return err
	}
	return nil
}

func readAndWriteDemo(sm proto.Message) {
	writeToFile("output.bin", sm)

	// Create an empty struct
	sm2 := &simple.SimpleMessage{}

	readFromFile("output.bin", sm2)
	fmt.Println("New sm2:", sm2.Id, sm2.IsSimple, sm2.Name, sm2.SampleList)

}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialize to bytes", err)
		return err
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}

	fmt.Println("Data has been written!")
	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Something went wrong reading a file", err)
		return err
	}

	err2 := proto.Unmarshal(in, pb)
	if err2 != nil {
		log.Fatalln("Couldn't put bytes into protobuf struct!")
		return err2
	}

	return nil
}

func doSimple() *simple.SimpleMessage {
	sm := simple.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 4, 7, 8},
	}

	fmt.Println(sm.Id, sm.IsSimple, sm.Name, sm.SampleList)

	sm.Name = "Renamed Simple message"

	fmt.Println(sm.Id, sm.IsSimple, sm.Name, sm.SampleList)

	fmt.Println("The Id is:", sm.GetId())

	return &sm
}
