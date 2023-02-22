package rlp

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"testing"
)

type simple struct {
	Data1 []byte
	Data2 [][]byte
	Data3 uint64
}

func (s *simple) MarshalRLP() ([]byte, error) {
	a := List{
		Data(s.Data1),
		List{
			Data(s.Data2[0]),
			Data(s.Data2[1]),
		},
		Int(s.Data3),
	}.Encode()
	return a, nil
}

func (s *simple) UnmarshalRLP(buf []byte) error {
	elems, pos, err := Decode(buf)
	if err != nil {
		return err
	}
	getElem := func() Element {
		if elems.IsList() {
			elem := elems.(List)[0]
			elems = elems.(List)[1:]
			return elem
		}
		return elems
	}
	if pos != len(buf) {
		return errors.New("invalid rlp")
	}
	if elems.IsList() {
		s.Data1 = []byte(getElem().(Data))
		ele := getElem()
		if !ele.IsList() {
			return errors.New("invalid rlp")
		}
		for _, e := range ele.(List) {
			if e.IsList() {
				return errors.New("invalid rlp")
			}
			s.Data2 = append(s.Data2, []byte(e.(Data)))
		}
		s.Data3 = uint64(getElem().(Data).Uint())
	} else {
		return errors.New("invalid rlp")
	}
	return nil
}

func validate(v Element, expected string) error {
	if strings.HasPrefix(expected, "0x") {
		expected = expected[2:]
	}
	buf, err := hex.DecodeString(expected)
	if err != nil {
		return err
	}
	dst := v.Encode()
	if !bytes.Equal(dst, buf) {
		return fmt.Errorf("bad")
	}
	return nil
}

func TestRLP(t *testing.T) {
	if err := testRLP(); err != nil {
		t.Fatal(err)
	}
}

func testRLP() error {
	var v Element
	// empty string
	v = String("")
	if err := validate(v, "0x80"); err != nil {
		return err
	}

	// bytestring00
	v = Bytes([]byte{0x0})
	if err := validate(v, "0x00"); err != nil {
		return err
	}

	// bytestring01
	v = Bytes([]byte{0x1})
	if err := validate(v, "0x01"); err != nil {
		return err
	}

	// bytestring7F
	v = Bytes([]byte{0x7F})
	if err := validate(v, "0x7F"); err != nil {
		return err
	}

	v = Bool(true)
	if err := validate(v, "0x01"); err != nil {
		return err
	}
	v = Bool(false)
	if err := validate(v, "0x80"); err != nil {
		return err
	}

	// short string
	v = String("dog")
	if err := validate(v, "0x83646f67"); err != nil {
		return err
	}

	// short string2
	v = String("Lorem ipsum dolor sit amet, consectetur adipisicing eli")
	if err := validate(v, "0xb74c6f72656d20697073756d20646f6c6f722073697420616d65742c20636f6e7365637465747572206164697069736963696e6720656c69"); err != nil {
		return err
	}

	// long string
	v = String("Lorem ipsum dolor sit amet, consectetur adipisicing elit")
	if err := validate(v, "0xb8384c6f72656d20697073756d20646f6c6f722073697420616d65742c20636f6e7365637465747572206164697069736963696e6720656c6974"); err != nil {
		return err
	}

	// long string 2
	v = String("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur mauris magna, suscipit sed vehicula non, iaculis faucibus tortor. Proin suscipit ultricies malesuada. Duis tortor elit, dictum quis tristique eu, ultrices at risus. Morbi a est imperdiet mi ullamcorper aliquet suscipit nec lorem. Aenean quis leo mollis, vulputate elit varius, consequat enim. Nulla ultrices turpis justo, et posuere urna consectetur nec. Proin non convallis metus. Donec tempor ipsum in mauris congue sollicitudin. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Suspendisse convallis sem vel massa faucibus, eget lacinia lacus tempor. Nulla quis ultricies purus. Proin auctor rhoncus nibh condimentum mollis. Aliquam consequat enim at metus luctus, a eleifend purus egestas. Curabitur at nibh metus. Nam bibendum, neque at auctor tristique, lorem libero aliquet arcu, non interdum tellus lectus sit amet eros. Cras rhoncus, metus ac ornare cursus, dolor justo ultrices metus, at ullamcorper volutpat")
	if err := validate(v, "0xb904004c6f72656d20697073756d20646f6c6f722073697420616d65742c20636f6e73656374657475722061646970697363696e6720656c69742e20437572616269747572206d6175726973206d61676e612c20737573636970697420736564207665686963756c61206e6f6e2c20696163756c697320666175636962757320746f72746f722e2050726f696e20737573636970697420756c74726963696573206d616c6573756164612e204475697320746f72746f7220656c69742c2064696374756d2071756973207472697374697175652065752c20756c7472696365732061742072697375732e204d6f72626920612065737420696d70657264696574206d6920756c6c616d636f7270657220616c6971756574207375736369706974206e6563206c6f72656d2e2041656e65616e2071756973206c656f206d6f6c6c69732c2076756c70757461746520656c6974207661726975732c20636f6e73657175617420656e696d2e204e756c6c6120756c74726963657320747572706973206a7573746f2c20657420706f73756572652075726e6120636f6e7365637465747572206e65632e2050726f696e206e6f6e20636f6e76616c6c6973206d657475732e20446f6e65632074656d706f7220697073756d20696e206d617572697320636f6e67756520736f6c6c696369747564696e2e20566573746962756c756d20616e746520697073756d207072696d697320696e206661756369627573206f726369206c756374757320657420756c74726963657320706f737565726520637562696c69612043757261653b2053757370656e646973736520636f6e76616c6c69732073656d2076656c206d617373612066617563696275732c2065676574206c6163696e6961206c616375732074656d706f722e204e756c6c61207175697320756c747269636965732070757275732e2050726f696e20617563746f722072686f6e637573206e69626820636f6e64696d656e74756d206d6f6c6c69732e20416c697175616d20636f6e73657175617420656e696d206174206d65747573206c75637475732c206120656c656966656e6420707572757320656765737461732e20437572616269747572206174206e696268206d657475732e204e616d20626962656e64756d2c206e6571756520617420617563746f72207472697374697175652c206c6f72656d206c696265726f20616c697175657420617263752c206e6f6e20696e74657264756d2074656c6c7573206c65637475732073697420616d65742065726f732e20437261732072686f6e6375732c206d65747573206163206f726e617265206375727375732c20646f6c6f72206a7573746f20756c747269636573206d657475732c20617420756c6c616d636f7270657220766f6c7574706174"); err != nil {
		return err
	}

	// zero
	v = Int(0)
	if err := validate(v, "0x80"); err != nil {
		return err
	}

	// smallint
	v = Int(1)
	if err := validate(v, "0x01"); err != nil {
		return err
	}

	// smallint2
	v = Int(16)
	if err := validate(v, "0x10"); err != nil {
		return err
	}

	// smallint3
	v = Int(79)
	if err := validate(v, "0x4f"); err != nil {
		return err
	}

	// smallint4
	v = Int(127)
	if err := validate(v, "0x7f"); err != nil {
		return err
	}

	// medium int

	// mediumint1
	v = Int(128)
	if err := validate(v, "0x8180"); err != nil {
		return err
	}

	// mediumint2
	v = Int(1000)
	if err := validate(v, "0x8203e8"); err != nil {
		return err
	}

	// mediumint3
	v = Int(100000)
	if err := validate(v, "0x830186a0"); err != nil {
		return err
	}

	// emptylist
	v = NewList()
	if err := validate(v, "0xc0"); err != nil {
		return err
	}

	// stringlist
	v = NewList()
	v = v.Add(String("dog")).Add(String("god")).Add(String("cat"))
	if err := validate(v, "0xcc83646f6783676f6483636174"); err != nil {
		return err
	}

	// multilist
	v = NewList()
	v = v.Add(String("zw"))
	vv := NewList().Add(Int(4))
	v = v.Add(vv)
	v = v.Add(Int(1))
	if err := validate(v, "0xc6827a77c10401"); err != nil {
		return err
	}
	return nil
}

func TestMarshal(t *testing.T) {
	s := &simple{
		Data1: []byte{0x01, 0x02, 0x03},
		Data2: [][]byte{
			[]byte{0x04, 0x05, 0x06},
			[]byte{0x07, 0x08, 0x09},
		},
		Data3: 0x0a,
	}
	a, err := MarshalRLP(s)
	if err != nil {
		t.Fatal(err)
	}
	if strings.ToLower(hex.EncodeToString(a)) != "ce83010203c883040506830708090a" {
		t.Fatal("unexpected result")
	}
	s = &simple{}
	err = UnmarshalRLP(a, s)
	if err != nil {
		t.Fatal(err)
	}
	if strings.ToLower(hex.EncodeToString(s.Data1)) != "010203" {
		t.Fatal("unexpected result")
	}
	if strings.ToLower(hex.EncodeToString(s.Data2[0])) != "040506" {
		t.Fatal("unexpected result")
	}
	if strings.ToLower(hex.EncodeToString(s.Data2[1])) != "070809" {
		t.Fatal("unexpected result")
	}
	if s.Data3 != 0x0a {
		t.Fatal("unexpected result")
	}
}

func TestUnMarshalRLP(t *testing.T) {
	s := &simple{}
	err := UnmarshalRLP([]byte{0xce, 0x83, 0x01, 0x02, 0x03, 0xc8, 0x83, 0x04, 0x05, 0x06, 0x83, 0x07, 0x08, 0x09, 0x0a}, s)
	if err != nil {
		t.Fatal(err)
	}
	if strings.ToLower(hex.EncodeToString(s.Data1)) != "010203" {
		t.Fatal("unexpected result")
	}
	if strings.ToLower(hex.EncodeToString(s.Data2[0])) != "040506" {
		t.Fatal("unexpected result")
	}
	if strings.ToLower(hex.EncodeToString(s.Data2[1])) != "070809" {
		t.Fatal("unexpected result")
	}
	if s.Data3 != 0x0a {
		t.Fatal("unexpected result")
	}
}

func TestDecodeTX(t *testing.T) {

	// Sample (legacy) transaction info
	// {
	// 	blockHash: "0xf792398ef0d5fbd4cccff85778032ce17074123eb143e6c658e544bc1b76ff4f",
	// 	blockNumber: 4,
	// 	from: "0x5d093e9b41911be5f5c4cf91b108bac5d130fa83",
	// 	gas: 40574,
	// 	gasPrice: 0,
	// 	hash: "0xea4bf65ee1f2ae6df7259676f4dc30e28a879fa7e7519a86c2ed6b9c59a544d8",
	// 	input: "0x ... see below",
	// 	nonce: 3,
	// 	r: "0x2e6e9728373680d0a7d75f99697d3887069dd5db4b9581c42bfb5749fb5fc80",
	// 	s: "0x32e8717112b372f41c4a2a46ad0ea807f56645990130cbbc60614f2240a3a1a",
	// 	to: "0x497eedc4299dea2f2a364be10025d0ad0f702de3",
	// 	transactionIndex: 0,
	// 	type: "0x0",
	// 	v: "0xfee",
	// 	value: 0
	// }

	// The raw transaction
	encoded, err := hex.DecodeString(
		"f901e70380829e7e94497eedc4299dea2f2a364be10025d0ad0f702de380b901843674e15c00000000000000000000000000000000000000000000000000000000000000a03f04a4e93ded4d2aaa1a41d617e55c59ac5f1b28a47047e2a526e76d45eb9681d19642e9120d63a9b7f5f537565a430d8ad321ef1bc76689a4b3edc861c640fc00000000000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000000000140000000000000000000000000000000000000000000000000000000000000000966665f73797374656d0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002e516d58747653456758626265506855684165364167426f3465796a7053434b437834515a4c50793548646a6177730000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001a1f7502c8f8797999c0c6b9c2da653ea736598ed0daa856c47ae71411aa8fea2820feea002e6e9728373680d0a7d75f99697d3887069dd5db4b9581c42bfb5749fb5fc80a0032e8717112b372f41c4a2a46ad0ea807f56645990130cbbc60614f2240a3a1a")
	if err != nil {
		t.Fatal(err)
	}

	// The input data
	inputData, err := hex.DecodeString(
		"3674e15c00000000000000000000000000000000000000000000000000000000000000a03f04a4e93ded4d2aaa1a41d617e55c59ac5f1b28a47047e2a526e76d45eb9681d19642e9120d63a9b7f5f537565a430d8ad321ef1bc76689a4b3edc861c640fc00000000000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000000000140000000000000000000000000000000000000000000000000000000000000000966665f73797374656d0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002e516d58747653456758626265506855684165364167426f3465796a7053434b437834515a4c50793548646a6177730000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001a1f7502c8f8797999c0c6b9c2da653ea736598ed0daa856c47ae71411aa8fea2")
	if err != nil {
		t.Fatal(err)
	}

	elem, next, err := Decode(encoded)
	if err != nil {
		t.Fatal(err)
	}
	if next != len(encoded) {
		t.Fatalf("expected %d, got %d", len(encoded), next)
	}
	rlpList := elem.(List)
	if len(rlpList) != 9 {
		t.Fatalf("expected 9, got %d", len(rlpList))
	}
	// Nonce
	if rlpList[0].(Data).Int().Int64() != 3 {
		t.Fatalf("expected 3, got %d", rlpList[0].(Data).Int().Int64())
	}
	// Gas Price
	if rlpList[1].(Data).Int().Int64() != 0 {
		t.Fatalf("expected 0, got %d", rlpList[1].(Data).Int().Int64())
	}
	// Gas Limit
	if rlpList[2].(Data).Int().Int64() != 40574 {
		t.Fatalf("expected 40574, got %d", rlpList[2].(Data).Int().Int64())
	}
	// To
	if rlpList[3].(Data).Hex() != "497eedc4299dea2f2a364be10025d0ad0f702de3" {
		t.Fatalf("expected 0x497eedc4299dea2f2a364be10025d0ad0f702de3, got %s", rlpList[3].(Data).Hex())
	}

	// Value
	if rlpList[4].(Data).Int().Int64() != 0 {
		t.Fatalf("expected 0, got %d", rlpList[4].(Data).Int().Int64())
	}
	// Data
	if !bytes.Equal(inputData, []byte(rlpList[5].(Data))) {
		t.Fatalf("expected %x, got %x", inputData, []byte(rlpList[5].(Data)))
	}
	// V
	if rlpList[6].(Data).Hex() != "0fee" {
		t.Fatalf("expected 0fee, got %s", rlpList[6].(Data).Hex())
	}
	// R
	if rlpList[7].(Data).Hex() != "02e6e9728373680d0a7d75f99697d3887069dd5db4b9581c42bfb5749fb5fc80" {
		t.Fatalf("expected 0x02e6e9728373680d0a7d75f99697d3887069dd5db4b9581c42bfb5749fb5fc80, got %s", rlpList[7].(Data).Hex())
	}
	// S
	if rlpList[8].(Data).Hex() != "032e8717112b372f41c4a2a46ad0ea807f56645990130cbbc60614f2240a3a1a" {
		t.Fatalf("expected 032e8717112b372f41c4a2a46ad0ea807f56645990130cbbc60614f2240a3a1a, got %s", rlpList[8].(Data).Hex())
	}
}

func BenchmarkRLP_SingleThreaded(b *testing.B) {
	for i := 0; i != b.N; i++ {
		s := &simple{
			Data1: []byte{0x01, 0x02, 0x03},
			Data2: [][]byte{
				[]byte{0x04, 0x05, 0x06},
				[]byte{0x07, 0x08, 0x09},
			},
			Data3: 0x0a,
		}
		MarshalRLP(s)
	}
}
