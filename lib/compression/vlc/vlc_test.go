package vlc

import (
	"bytes"
	"testing"

	"github.com/FlynntDev/go-archiver/lib/compression/vlc/table"
)

func TestPrepareText(t *testing.T) {
	input := "My name is Ted"
	expected := "!my name is !ted"

	result := prepareText(input)
	if result != expected {
		t.Errorf("prepareText() = %v, want %v", result, expected)
	}
}

func TestExportText(t *testing.T) {
	input := "!my name is !ted"
	expected := "My name is Ted"

	result := exportText(input)
	if result != expected {
		t.Errorf("exportText() = %v, want %v", result, expected)
	}
}

func TestEncodeBin(t *testing.T) {
	tbl := table.EncodingTable{
		'a': "00",
		'b': "01",
		'c': "10",
	}
	input := "abc"
	expected := "000110"

	result := encodeBin(input, tbl)
	if result != expected {
		t.Errorf("encodeBin() = %v, want %v", result, expected)
	}
}

func TestBuildEncodedFileAndParseFile(t *testing.T) {
	tbl := table.EncodingTable{
		'a': "00",
		'b': "01",
		'c': "10",
	}
	data := "000110"

	encodedFile := buildEncodedFile(tbl, data)
	parsedTbl, parsedData := parseFile(encodedFile)

	if !compareTables(tbl, parsedTbl) {
		t.Errorf("parseFile() table = %v, want %v", parsedTbl, tbl)
	}

	if parsedData != data {
		t.Errorf("parseFile() data = %v, want %v", parsedData, data)
	}
}

func TestEncodeTableAndDecodeTable(t *testing.T) {
	tbl := table.EncodingTable{
		'a': "00",
		'b': "01",
		'c': "10",
	}

	encoded := encodeTable(tbl)
	decoded := decodeTable(encoded)

	if !compareTables(tbl, decoded) {
		t.Errorf("decodeTable() = %v, want %v", decoded, tbl)
	}
}

func TestBin(t *testing.T) {
	tbl := table.EncodingTable{
		'a': "00",
		'b': "01",
	}
	input := 'a'
	expected := "00"

	result := bin(input, tbl)
	if result != expected {
		t.Errorf("bin() = %v, want %v", result, expected)
	}
}

func TestEncodeInt(t *testing.T) {
	input := 12345
	expected := []byte{0x00, 0x00, 0x30, 0x39}

	result := encodeInt(input)
	if !bytes.Equal(result, expected) {
		t.Errorf("encodeInt() = %v, want %v", result, expected)
	}
}

// Helper function to compare two EncodingTables
func compareTables(a, b table.EncodingTable) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}

// Function to prepare text by replacing uppercase letters with lowercase and adding '!' before them
func prepareText(input string) string {
	var result bytes.Buffer
	for _, r := range input {
		if r >= 'A' && r <= 'Z' {
			result.WriteRune('!')
			result.WriteRune(r + ('a' - 'A'))
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}
