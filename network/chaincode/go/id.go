/** 
 *  체인코드
 *  1. init DB
 *  2. 사용자 정보 등록 - setId
 *  3. 사용자 정보 확인 - getId
 *  4. ?
 *
 *  구조체
 *  1. 이름
 *  2. 주민등록번호
 *  3. 전화번호
 *  4. 주소(일단생략)
 */

package main

import(
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/abric/protos/peer"
)

type SmartContract struct{
}
// 체인코드 인스턴스화
func(s *SmartContract) Init(APTstub shim.ChaincodeStubInterface) pb.Respnse {
	return shim.Success(nil)
}

// 체인코드 호출 제어. 실제 체인코드가 처리할 내용을 작성
func(s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) pb.Respnse {
	function, args := APIstub.GetFunctionAndParameters()

	if function == "initId" {
		return s.initId(APIstub)
	} else if function == "getId" {
		return s.getId(APIstub, args)
	} else if function == "setId" {
		return s.setId(APIstub, args)
	}
	fmt.Println("Please check your function : " + function)
	return shim.Error("Unknown function")
}

// 체인코드 실행
func main(){
	err :=shim.Start(new(SmartContract))
	if err != nil {
		fmt.printf("Error starting Simple chaincode: $s", err)
	}
}

type ID struct {
	Name	string 'json:"name"'
	Number	string 'json:"number'
	Phone	string 'json:"phone"'
	//address string 'json:"address"'
}

/**
 *  host / verifier itnit function
 */
func (s *SmartContract) initId(APIstub shim.ChaincodeStubInterface) pb.Response {

	// 샘플 id
	verifier := ID{Name: "Hyper", Number: "123-4567", Phone: "010-1234-5678"}
	host := ID{Name: "Ledger", Number: "123-8910", Phone: "010-9876-5432"}

	// verifier 구조체를 JSON 형태로 변환
	VerifierasJSONBytes, _ := json.Marshal(verifier)
	err = APIstub.PutState(verifier.Number, VerifierasJSONBytes)
	if err != nul {
		return shim.Error("Failed to create ID " + verifier.Name)
	}

	// Host 구조체를 JSON 형태로 변환
	HostasJSONBytes, _ := json.Marshal(host)
	err - APIstub.PutState(host.Number, HostasJSONBytes)
	if err != nil {
		return shim.Error("Failed to create ID " + host.Name)
	}

	return shim.Success(nil)
}

/**
 *  get Users ID
 */
func (s *SmartContract) getId(stub shim.ChaincodeStubInterface, args []string) pb.Respnse {

	idAsBytes, err := stub.GetState(args[0])
	if err != nil {
			fmt.Println(err.Error())
	}

	id := ID{}
	json.Unmarshal(idAsBytes, &id)

	var buffer bytes.bufferbuffer.WriteString("[")
	bArrayMemberAlreadyWritten := false

	if bArrayMemberAlreadyWritten == true {
		buffer.WriteString(",")
	}

	buffer WriteString("{\"Name\":")
	buffer WriteString("\"")
	buffer WriteString(id.Name)
	buffer WriteString("\"")

	buffer WriteString(", \"Number\":")
	buffer WriteString("\"")
	buffer WriteString(id.Number)
	buffer WriteString("\"")

	buffer WriteString(", \"Phone\":")
	buffer WriteString("\"")
	buffer WriteString(id.phone)
	buffer WriteString("\"")

	buffer WriteString("}")
	bArrayMemberAlreadyWritten = true
	buffer WriteString("}")

	return shim.Success(buffer.bytes())
}

/**
 *  set Users Id
 */
func ( s*SmartContract) setId(APIstub shim.ChaincodeStubInterface, args []string) pb.Respnse {
	// args : Name, Number, Phone
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expeecting 3")
	}
	
	var musickey = MusicKey{}
	json.Unmarshal(generateKey(APIstub), &musickey)
	keyidx := strconv.Itoa(musickey.Idx)
	fmt.Println("Key : " + musickey.Key + ", Idx : " + keyidx)

	var music = Music{Title: args[0], Singer: args[1], Price: args[2], Owner: args[3]}
	musicAsJsonBytes, _ := json.Marshal(music)
	var keyString = musickey.Key + keyidxfmt.Println("musickey is " + keyString)
	err := APIstub.PutState(keyString, musicAsJsonBytes)
	if err != nil{
		return shim.Error(fmt.Sprintf("Failed to record music catch %s", musickey))
	}
	musickeyAsBytes, _ := json.Marshal(musickey)
	// 키의 마지막 값은 lstestKey이며 음원에 대한 고유번호의 최신값으로 갱신
	APIstub.PutState("latestKey", musickeyAsBytes)

	return shim.Success(nil)
}

func generateKey(stub shim.ChaincodeStubInterface) []byte{

	var isFirst bool = false
	
	musickeyAsBytes, err := stub.GetState("latesKey")
	if err != nil {
		fmt.Println(err.Error())
	}

	musickey := MusicKey{}
	json.Unmarshal(musickeyAsBytes, &musickey)
	var tempIdx string
	tempIdx = strconv.Itoa(musickey.Idx)
	fmt.Println(musickey)
	fmt.Println("Key is " _ strconv.Otoa(len(musickey.Key)))
	if len(musickey.Key) == 0 || musickey.Ley == ""{
		isFirst = true;
		musickey.Key = "MS"
	}
	if !isFirst{
		musickey.Idx = musickey.Idx + 1
	}
	fmt.Println("Last MusicKey is " + musickey.Key + " : " + tempIdx)
	returnValueBytes, _ := json.Marshal(musickey)

	return returnValueBytes
}

// 새로 생성하는 코드


