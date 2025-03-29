package dimse_test

// import (
// 	"encoding/binary"
// 	"testing"

// 	"github.com/grailbio/go-dicom/dicomio"
// 	"github.com/mlibanori/go-netdicom/dimse/dimse_commands"
// )

// func testDIMSE(t *testing.T, v Message) {
// 	e := dicomio.NewBytesEncoder(binary.LittleEndian, dicomio.ImplicitVR)
// 	EncodeMessage(e, v)
// 	bytes := e.Bytes()
// 	d := dicomio.NewBytesDecoder(bytes, binary.LittleEndian, dicomio.ImplicitVR)
// 	v2 := dimse_commands.ReadMessage(d)
// 	err := d.Finish()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if v.String() != v2.String() {
// 		t.Errorf("%v <-> %v", v, v2)
// 	}
// }

// func TestCStoreRq(t *testing.T) {
// 	testDIMSE(t, &dimse_commands.CStoreRq{
// 		"1.2.3",
// 		0x1234,
// 		0x2345,
// 		1,
// 		"3.4.5",
// 		"foohah",
// 		0x3456, nil})
// }

// func TestCStoreRsp(t *testing.T) {
// 	testDIMSE(t, &dimse_commands.CStoreRsp{
// 		"1.2.3",
// 		0x1234,
// 		CommandDataSetTypeNull,
// 		"3.4.5",
// 		Status{Status: StatusCode(0x3456)},
// 		nil})
// }

// func TestCEchoRq(t *testing.T) {
// 	testDIMSE(t, &dimse_commands.CEchoRq{0x1234, 1, nil})
// }

// func TestCEchoRsp(t *testing.T) {
// 	testDIMSE(t, &dimse_commands.CEchoRsp{0x1234, 1,
// 		Status{Status: StatusCode(0x2345)},
// 		nil})
// }
