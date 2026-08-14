package types

import "testing"

func TestConnection(t *testing.T) {

	t.Run("IsConnection(client)", func(t *testing.T) {

		is1 := IsConnection("[tcp] 127.0.0.1:1337 -> 1.3.3.7:80")
		is2 := IsConnection("[any] 127.0.0.1:1337 -> 1.3.3.7:80")
		is3 := IsConnection("[udp] 127.0.0.1:1337 -> 1.3.3.7:80")

		is4 := IsConnection("[tcp] *:1337 -> 1.3.3.7:80")
		is5 := IsConnection("[any] *:1337 -> 1.3.3.7:80")
		is6 := IsConnection("[udp] *:1337 -> 1.3.3.7:80")

		is7 := IsConnection("[tcp] 127.0.0.1:1337 -> *:80")
		is8 := IsConnection("[any] 127.0.0.1:1337 -> *:80")
		is9 := IsConnection("[udp] 127.0.0.1:1337 -> *:80")

		is10 := IsConnection("[tcp] *:1337 -> *:80")
		is11 := IsConnection("[any] *:1337 -> *:80")
		is12 := IsConnection("[udp] *:1337 -> *:80")

		is13 := IsConnection("[tcp] 127.0.0.1:1337 -> 0.0.0.0:80")
		is14 := IsConnection("[any] 127.0.0.1:1337 -> 0.0.0.0:80")
		is15 := IsConnection("[udp] 127.0.0.1:1337 -> 0.0.0.0:80")

		is16 := IsConnection("[tcp] *:1337 -> 0.0.0.0:80")
		is17 := IsConnection("[any] *:1337 -> 0.0.0.0:80")
		is18 := IsConnection("[udp] *:1337 -> 0.0.0.0:80")

		if is1 != true {
			t.Errorf("Expected %t to be true", is1)
		}

		if is2 != true {
			t.Errorf("Expected %t to be true", is2)
		}

		if is3 != true {
			t.Errorf("Expected %t to be true", is3)
		}

		if is4 != true {
			t.Errorf("Expected %t to be true", is4)
		}

		if is5 != true {
			t.Errorf("Expected %t to be true", is5)
		}

		if is6 != true {
			t.Errorf("Expected %t to be true", is6)
		}

		if is7 != true {
			t.Errorf("Expected %t to be true", is7)
		}

		if is8 != true {
			t.Errorf("Expected %t to be true", is8)
		}

		if is9 != true {
			t.Errorf("Expected %t to be true", is9)
		}

		if is10 != true {
			t.Errorf("Expected %t to be true", is10)
		}

		if is11 != true {
			t.Errorf("Expected %t to be true", is11)
		}

		if is12 != true {
			t.Errorf("Expected %t to be true", is12)
		}

		if is13 != true {
			t.Errorf("Expected %t to be true", is13)
		}

		if is14 != true {
			t.Errorf("Expected %t to be true", is14)
		}

		if is15 != true {
			t.Errorf("Expected %t to be true", is15)
		}

		if is16 != true {
			t.Errorf("Expected %t to be true", is16)
		}

		if is17 != true {
			t.Errorf("Expected %t to be true", is17)
		}

		if is18 != true {
			t.Errorf("Expected %t to be true", is18)
		}

	})

	t.Run("IsConnection(server)", func(t *testing.T) {

		is1 := IsConnection("[tcp] 127.0.0.1:1337 <- 1.3.3.7:80")
		is2 := IsConnection("[any] 127.0.0.1:1337 <- 1.3.3.7:80")
		is3 := IsConnection("[udp] 127.0.0.1:1337 <- 1.3.3.7:80")

		is4 := IsConnection("[tcp] *:1337 <- 1.3.3.7:80")
		is5 := IsConnection("[any] *:1337 <- 1.3.3.7:80")
		is6 := IsConnection("[udp] *:1337 <- 1.3.3.7:80")

		is7 := IsConnection("[tcp] 127.0.0.1:1337 <- *:80")
		is8 := IsConnection("[any] 127.0.0.1:1337 <- *:80")
		is9 := IsConnection("[udp] 127.0.0.1:1337 <- *:80")

		is10 := IsConnection("[tcp] *:1337 <- *:80")
		is11 := IsConnection("[any] *:1337 <- *:80")
		is12 := IsConnection("[udp] *:1337 <- *:80")

		is13 := IsConnection("[tcp] 127.0.0.1:1337 <- 0.0.0.0:80")
		is14 := IsConnection("[any] 127.0.0.1:1337 <- 0.0.0.0:80")
		is15 := IsConnection("[udp] 127.0.0.1:1337 <- 0.0.0.0:80")

		is16 := IsConnection("[tcp] *:1337 <- 0.0.0.0:80")
		is17 := IsConnection("[any] *:1337 <- 0.0.0.0:80")
		is18 := IsConnection("[udp] *:1337 <- 0.0.0.0:80")

		if is1 != true {
			t.Errorf("Expected %t to be true", is1)
		}

		if is2 != true {
			t.Errorf("Expected %t to be true", is2)
		}

		if is3 != true {
			t.Errorf("Expected %t to be true", is3)
		}

		if is4 != true {
			t.Errorf("Expected %t to be true", is4)
		}

		if is5 != true {
			t.Errorf("Expected %t to be true", is5)
		}

		if is6 != true {
			t.Errorf("Expected %t to be true", is6)
		}

		if is7 != true {
			t.Errorf("Expected %t to be true", is7)
		}

		if is8 != true {
			t.Errorf("Expected %t to be true", is8)
		}

		if is9 != true {
			t.Errorf("Expected %t to be true", is9)
		}

		if is10 != true {
			t.Errorf("Expected %t to be true", is10)
		}

		if is11 != true {
			t.Errorf("Expected %t to be true", is11)
		}

		if is12 != true {
			t.Errorf("Expected %t to be true", is12)
		}

		if is13 != true {
			t.Errorf("Expected %t to be true", is13)
		}

		if is14 != true {
			t.Errorf("Expected %t to be true", is14)
		}

		if is15 != true {
			t.Errorf("Expected %t to be true", is15)
		}

		if is16 != true {
			t.Errorf("Expected %t to be true", is16)
		}

		if is17 != true {
			t.Errorf("Expected %t to be true", is17)
		}

		if is18 != true {
			t.Errorf("Expected %t to be true", is18)
		}

	})

	t.Run("IsConnection(peer)", func(t *testing.T) {

		is1 := IsConnection("[tcp] 127.0.0.1:1337 <-> 1.3.3.7:80")
		is2 := IsConnection("[any] 127.0.0.1:1337 <-> 1.3.3.7:80")
		is3 := IsConnection("[udp] 127.0.0.1:1337 <-> 1.3.3.7:80")

		is4 := IsConnection("[tcp] *:1337 <-> 1.3.3.7:80")
		is5 := IsConnection("[any] *:1337 <-> 1.3.3.7:80")
		is6 := IsConnection("[udp] *:1337 <-> 1.3.3.7:80")

		is7 := IsConnection("[tcp] 127.0.0.1:1337 <-> *:80")
		is8 := IsConnection("[any] 127.0.0.1:1337 <-> *:80")
		is9 := IsConnection("[udp] 127.0.0.1:1337 <-> *:80")

		is10 := IsConnection("[tcp] *:1337 <-> *:80")
		is11 := IsConnection("[any] *:1337 <-> *:80")
		is12 := IsConnection("[udp] *:1337 <-> *:80")

		is13 := IsConnection("[tcp] 127.0.0.1:1337 <-> 0.0.0.0:80")
		is14 := IsConnection("[any] 127.0.0.1:1337 <-> 0.0.0.0:80")
		is15 := IsConnection("[udp] 127.0.0.1:1337 <-> 0.0.0.0:80")

		is16 := IsConnection("[tcp] *:1337 <-> 0.0.0.0:80")
		is17 := IsConnection("[any] *:1337 <-> 0.0.0.0:80")
		is18 := IsConnection("[udp] *:1337 <-> 0.0.0.0:80")

		if is1 != true {
			t.Errorf("Expected %t to be true", is1)
		}

		if is2 != true {
			t.Errorf("Expected %t to be true", is2)
		}

		if is3 != true {
			t.Errorf("Expected %t to be true", is3)
		}

		if is4 != true {
			t.Errorf("Expected %t to be true", is4)
		}

		if is5 != true {
			t.Errorf("Expected %t to be true", is5)
		}

		if is6 != true {
			t.Errorf("Expected %t to be true", is6)
		}

		if is7 != true {
			t.Errorf("Expected %t to be true", is7)
		}

		if is8 != true {
			t.Errorf("Expected %t to be true", is8)
		}

		if is9 != true {
			t.Errorf("Expected %t to be true", is9)
		}

		if is10 != true {
			t.Errorf("Expected %t to be true", is10)
		}

		if is11 != true {
			t.Errorf("Expected %t to be true", is11)
		}

		if is12 != true {
			t.Errorf("Expected %t to be true", is12)
		}

		if is13 != true {
			t.Errorf("Expected %t to be true", is13)
		}

		if is14 != true {
			t.Errorf("Expected %t to be true", is14)
		}

		if is15 != true {
			t.Errorf("Expected %t to be true", is15)
		}

		if is16 != true {
			t.Errorf("Expected %t to be true", is16)
		}

		if is17 != true {
			t.Errorf("Expected %t to be true", is17)
		}

		if is18 != true {
			t.Errorf("Expected %t to be true", is18)
		}

	})

	t.Run("IsConnection(invalid)", func(t *testing.T) {

		// invalid protocol
		is1 := IsConnection("127.0.0.1:1337 -> 1.3.3.7:80")
		is2 := IsConnection("127.0.0.1:1337 <- 1.3.3.7:80")
		is3 := IsConnection("127.0.0.1:1337 <-> 1.3.3.7:80")

		// invalid source host
		is4 := IsConnection("[tcp] example.com:1337 -> 1.3.3.7:80")
		is5 := IsConnection("[any] example.com:1337 <- 1.3.3.7:80")
		is6 := IsConnection("[udp] example.com:1337 <-> 1.3.3.7:80")

		// invalid source port
		is7 := IsConnection("[tcp] *:99999 -> 1.3.3.7:80")
		is8 := IsConnection("[any] *:0 <- 1.3.3.7:80")
		is9 := IsConnection("[udp] *:-1 <-> 1.3.3.7:80")

		// invalid type
		is10 := IsConnection("[tcp] 127.0.0.1:1337 - 1.3.3.7:80")
		is11 := IsConnection("[any] 127.0.0.1:1337 _ 1.3.3.7:80")
		is12 := IsConnection("[udp] 127.0.0.1:1337 invalid 1.3.3.7:80")

		// invalid source port
		is13 := IsConnection("[tcp] 127.0.0.1:1337 -> 1.3.3.7:99999")
		is14 := IsConnection("[any] 127.0.0.1:1337 <- 1.3.3.7:0")
		is15 := IsConnection("[udp] 127.0.0.1:1337 <-> 1.3.3.7:-1")

		if is1 != false {
			t.Errorf("Expected %t to be false", is1)
		}

		if is2 != false {
			t.Errorf("Expected %t to be false", is2)
		}

		if is3 != false {
			t.Errorf("Expected %t to be false", is3)
		}

		if is4 != false {
			t.Errorf("Expected %t to be false", is4)
		}

		if is5 != false {
			t.Errorf("Expected %t to be false", is5)
		}

		if is6 != false {
			t.Errorf("Expected %t to be false", is6)
		}

		if is7 != false {
			t.Errorf("Expected %t to be false", is7)
		}

		if is8 != false {
			t.Errorf("Expected %t to be false", is8)
		}

		if is9 != false {
			t.Errorf("Expected %t to be false", is9)
		}

		if is10 != false {
			t.Errorf("Expected %t to be false", is10)
		}

		if is11 != false {
			t.Errorf("Expected %t to be false", is11)
		}

		if is12 != false {
			t.Errorf("Expected %t to be false", is12)
		}

		if is13 != false {
			t.Errorf("Expected %t to be false", is13)
		}

		if is14 != false {
			t.Errorf("Expected %t to be false", is14)
		}

		if is15 != false {
			t.Errorf("Expected %t to be false", is15)
		}

	})

	t.Run("ParseConnection(client)", func(t *testing.T) {

		connection1 := ParseConnection("[tcp] 127.0.0.1:1337 -> 1.3.3.7:80")
		connection2 := ParseConnection("[any] 127.0.0.1:1337 -> 1.3.3.7:80")
		connection3 := ParseConnection("[udp] 127.0.0.1:1337 -> 1.3.3.7:80")

		connection4 := ParseConnection("[tcp] *:1337 -> 1.3.3.7:80")
		connection5 := ParseConnection("[any] *:1337 -> 1.3.3.7:80")
		connection6 := ParseConnection("[udp] *:1337 -> 1.3.3.7:80")

		connection7 := ParseConnection("[tcp] 127.0.0.1:1337 -> *:80")
		connection8 := ParseConnection("[any] 127.0.0.1:1337 -> *:80")
		connection9 := ParseConnection("[udp] 127.0.0.1:1337 -> *:80")

		connection10 := ParseConnection("[tcp] *:1337 -> *:80")
		connection11 := ParseConnection("[any] *:1337 -> *:80")
		connection12 := ParseConnection("[udp] *:1337 -> *:80")

		connection13 := ParseConnection("[tcp] 127.0.0.1:1337 -> 0.0.0.0:80")
		connection14 := ParseConnection("[any] 127.0.0.1:1337 -> 0.0.0.0:80")
		connection15 := ParseConnection("[udp] 127.0.0.1:1337 -> 0.0.0.0:80")

		connection16 := ParseConnection("[tcp] *:1337 -> 0.0.0.0:80")
		connection17 := ParseConnection("[any] *:1337 -> 0.0.0.0:80")
		connection18 := ParseConnection("[udp] *:1337 -> 0.0.0.0:80")

		if connection1 == nil {
			t.Errorf("Expected %s to be valid", "[tcp] 127.0.0.1:1337 -> 1.3.3.7:80")
		} else if connection1.String() != "[tcp] 127.0.0.1:1337 -> 1.3.3.7:80" {
			t.Errorf("Expected %s to be %s", connection1.String(), "[tcp] 127.0.0.1:1337 -> 1.3.3.7:80")
		}

		if connection2 == nil {
			t.Errorf("Expected %s to be valid", "[any] 127.0.0.1:1337 -> 1.3.3.7:80")
		} else if connection2.String() != "[any] 127.0.0.1:1337 -> 1.3.3.7:80" {
			t.Errorf("Expected %s to be %s", connection2.String(), "[any] 127.0.0.1:1337 -> 1.3.3.7:80")
		}

		if connection3 == nil {
			t.Errorf("Expected %s to be valid", "[udp] 127.0.0.1:1337 -> 1.3.3.7:80")
		} else if connection3.String() != "[udp] 127.0.0.1:1337 -> 1.3.3.7:80" {
			t.Errorf("Expected %s to be %s", connection3.String(), "[udp] 127.0.0.1:1337 -> 1.3.3.7:80")
		}

		if connection4 == nil {
			t.Errorf("Expected %s to be valid", "[tcp] *:1337 -> 1.3.3.7:80")
		} else if connection4.String() != "[tcp] 0.0.0.0:1337 -> 1.3.3.7:80" {
			t.Errorf("Expected %s to be %s", connection4.String(), "[tcp] 0.0.0.0:1337 -> 1.3.3.7:80")
		}

		if connection5 == nil {
			t.Errorf("Expected %s to be valid", "[any] *:1337 -> 1.3.3.7:80")
		} else if connection5.String() != "[any] 0.0.0.0:1337 -> 1.3.3.7:80" {
			t.Errorf("Expected %s to be %s", connection5.String(), "[any] 0.0.0.0:1337 -> 1.3.3.7:80")
		}

		if connection6 == nil {
			t.Errorf("Expected %s to be valid", "[udp] *:1337 -> 1.3.3.7:80")
		} else if connection6.String() != "[udp] 0.0.0.0:1337 -> 1.3.3.7:80" {
			t.Errorf("Expected %s to be %s", connection6.String(), "[udp] 0.0.0.0:1337 -> 1.3.3.7:80")
		}

		if connection7 == nil {
			t.Errorf("Expected %s to be valid", "[tcp] 127.0.0.1:1337 -> *:80")
		} else if connection7.String() != "[tcp] 127.0.0.1:1337 -> 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection7.String(), "[tcp] 127.0.0.1:1337 -> 0.0.0.0:80")
		}

		if connection8 == nil {
			t.Errorf("Expected %s to be valid", "[any] 127.0.0.1:1337 -> *:80")
		} else if connection8.String() != "[any] 127.0.0.1:1337 -> 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection8.String(), "[any] 127.0.0.1:1337 -> 0.0.0.0:80")
		}

		if connection9 == nil {
			t.Errorf("Expected %s to be valid", "[udp] 127.0.0.1:1337 -> *:80")
		} else if connection9.String() != "[udp] 127.0.0.1:1337 -> 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection9.String(), "[udp] 127.0.0.1:1337 -> 0.0.0.0:80")
		}

		if connection10 == nil {
			t.Errorf("Expected %s to be valid", "[tcp] *:1337 -> *:80")
		} else if connection10.String() != "[tcp] 0.0.0.0:1337 -> 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection10.String(), "[tcp] 0.0.0.0:1337 -> 0.0.0.0:80")
		}

		if connection11 == nil {
			t.Errorf("Expected %s to be valid", "[any] *:1337 -> *:80")
		} else if connection11.String() != "[any] 0.0.0.0:1337 -> 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection11.String(), "[any] 0.0.0.0:1337 -> 0.0.0.0:80")
		}

		if connection12 == nil {
			t.Errorf("Expected %s to be valid", "[udp] *:1337 -> *:80")
		} else if connection12.String() != "[udp] 0.0.0.0:1337 -> 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection12.String(), "[udp] 0.0.0.0:1337 -> 0.0.0.0:80")
		}

		if connection13 == nil {
			t.Errorf("Expected %s to be valid", "[tcp] 127.0.0.1:1337 -> 0.0.0.0:80")
		} else if connection13.String() != "[tcp] 127.0.0.1:1337 -> 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection13.String(), "[tcp] 127.0.0.1:1337 -> 0.0.0.0:80")
		}

		if connection14 == nil {
			t.Errorf("Expected %s to be valid", "[any] 127.0.0.1:1337 -> 0.0.0.0:80")
		} else if connection14.String() != "[any] 127.0.0.1:1337 -> 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection14.String(), "[any] 127.0.0.1:1337 -> 0.0.0.0:80")
		}

		if connection15 == nil {
			t.Errorf("Expected %s to be valid", "[udp] 127.0.0.1:1337 -> 0.0.0.0:80")
		} else if connection15.String() != "[udp] 127.0.0.1:1337 -> 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection15.String(), "[udp] 127.0.0.1:1337 -> 0.0.0.0:80")
		}

		if connection16 == nil {
			t.Errorf("Expected %s to be valid", "[tcp] *:1337 -> 0.0.0.0:80")
		} else if connection16.String() != "[tcp] 0.0.0.0:1337 -> 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection16.String(), "[tcp] 0.0.0.0:1337 -> 0.0.0.0:80")
		}

		if connection17 == nil {
			t.Errorf("Expected %s to be valid", "[any] *:1337 -> 0.0.0.0:80")
		} else if connection17.String() != "[any] 0.0.0.0:1337 -> 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection17.String(), "[any] 0.0.0.0:1337 -> 0.0.0.0:80")
		}

		if connection18 == nil {
			t.Errorf("Expected %s to be valid", "[udp] *:1337 -> 0.0.0.0:80")
		} else if connection18.String() != "[udp] 0.0.0.0:1337 -> 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection18.String(), "[udp] 0.0.0.0:1337 -> 0.0.0.0:80")
		}

	})

	t.Run("ParseConnection(server)", func(t *testing.T) {

		connection1 := ParseConnection("[tcp] 127.0.0.1:1337 <- 1.3.3.7:80")
		connection2 := ParseConnection("[any] 127.0.0.1:1337 <- 1.3.3.7:80")
		connection3 := ParseConnection("[udp] 127.0.0.1:1337 <- 1.3.3.7:80")

		connection4 := ParseConnection("[tcp] *:1337 <- 1.3.3.7:80")
		connection5 := ParseConnection("[any] *:1337 <- 1.3.3.7:80")
		connection6 := ParseConnection("[udp] *:1337 <- 1.3.3.7:80")

		connection7 := ParseConnection("[tcp] 127.0.0.1:1337 <- *:80")
		connection8 := ParseConnection("[any] 127.0.0.1:1337 <- *:80")
		connection9 := ParseConnection("[udp] 127.0.0.1:1337 <- *:80")

		connection10 := ParseConnection("[tcp] *:1337 <- *:80")
		connection11 := ParseConnection("[any] *:1337 <- *:80")
		connection12 := ParseConnection("[udp] *:1337 <- *:80")

		connection13 := ParseConnection("[tcp] 127.0.0.1:1337 <- 0.0.0.0:80")
		connection14 := ParseConnection("[any] 127.0.0.1:1337 <- 0.0.0.0:80")
		connection15 := ParseConnection("[udp] 127.0.0.1:1337 <- 0.0.0.0:80")

		connection16 := ParseConnection("[tcp] *:1337 <- 0.0.0.0:80")
		connection17 := ParseConnection("[any] *:1337 <- 0.0.0.0:80")
		connection18 := ParseConnection("[udp] *:1337 <- 0.0.0.0:80")

		if connection1 == nil {
			t.Errorf("Expected %s to be valid", "[tcp] 127.0.0.1:1337 <- 1.3.3.7:80")
		} else if connection1.String() != "[tcp] 127.0.0.1:1337 <- 1.3.3.7:80" {
			t.Errorf("Expected %s to be %s", connection1.String(), "[tcp] 127.0.0.1:1337 <- 1.3.3.7:80")
		}

		if connection2 == nil {
			t.Errorf("Expected %s to be valid", "[any] 127.0.0.1:1337 <- 1.3.3.7:80")
		} else if connection2.String() != "[any] 127.0.0.1:1337 <- 1.3.3.7:80" {
			t.Errorf("Expected %s to be %s", connection2.String(), "[any] 127.0.0.1:1337 <- 1.3.3.7:80")
		}

		if connection3 == nil {
			t.Errorf("Expected %s to be valid", "[udp] 127.0.0.1:1337 <- 1.3.3.7:80")
		} else if connection3.String() != "[udp] 127.0.0.1:1337 <- 1.3.3.7:80" {
			t.Errorf("Expected %s to be %s", connection3.String(), "[udp] 127.0.0.1:1337 <- 1.3.3.7:80")
		}

		if connection4 == nil {
			t.Errorf("Expected %s to be valid", "[tcp] *:1337 <- 1.3.3.7:80")
		} else if connection4.String() != "[tcp] 0.0.0.0:1337 <- 1.3.3.7:80" {
			t.Errorf("Expected %s to be %s", connection4.String(), "[tcp] 0.0.0.0:1337 <- 1.3.3.7:80")
		}

		if connection5 == nil {
			t.Errorf("Expected %s to be valid", "[any] *:1337 <- 1.3.3.7:80")
		} else if connection5.String() != "[any] 0.0.0.0:1337 <- 1.3.3.7:80" {
			t.Errorf("Expected %s to be %s", connection5.String(), "[any] 0.0.0.0:1337 <- 1.3.3.7:80")
		}

		if connection6 == nil {
			t.Errorf("Expected %s to be valid", "[udp] *:1337 <- 1.3.3.7:80")
		} else if connection6.String() != "[udp] 0.0.0.0:1337 <- 1.3.3.7:80" {
			t.Errorf("Expected %s to be %s", connection6.String(), "[udp] 0.0.0.0:1337 <- 1.3.3.7:80")
		}

		if connection7 == nil {
			t.Errorf("Expected %s to be valid", "[tcp] 127.0.0.1:1337 <- *:80")
		} else if connection7.String() != "[tcp] 127.0.0.1:1337 <- 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection7.String(), "[tcp] 127.0.0.1:1337 <- 0.0.0.0:80")
		}

		if connection8 == nil {
			t.Errorf("Expected %s to be valid", "[any] 127.0.0.1:1337 <- *:80")
		} else if connection8.String() != "[any] 127.0.0.1:1337 <- 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection8.String(), "[any] 127.0.0.1:1337 <- 0.0.0.0:80")
		}

		if connection9 == nil {
			t.Errorf("Expected %s to be valid", "[udp] 127.0.0.1:1337 <- *:80")
		} else if connection9.String() != "[udp] 127.0.0.1:1337 <- 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection9.String(), "[udp] 127.0.0.1:1337 <- 0.0.0.0:80")
		}

		if connection10 == nil {
			t.Errorf("Expected %s to be valid", "[tcp] *:1337 <- *:80")
		} else if connection10.String() != "[tcp] 0.0.0.0:1337 <- 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection10.String(), "[tcp] 0.0.0.0:1337 <- 0.0.0.0:80")
		}

		if connection11 == nil {
			t.Errorf("Expected %s to be valid", "[any] *:1337 <- *:80")
		} else if connection11.String() != "[any] 0.0.0.0:1337 <- 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection11.String(), "[any] 0.0.0.0:1337 <- 0.0.0.0:80")
		}

		if connection12 == nil {
			t.Errorf("Expected %s to be valid", "[udp] *:1337 <- *:80")
		} else if connection12.String() != "[udp] 0.0.0.0:1337 <- 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection12.String(), "[udp] 0.0.0.0:1337 <- 0.0.0.0:80")
		}

		if connection13 == nil {
			t.Errorf("Expected %s to be valid", "[tcp] 127.0.0.1:1337 <- 0.0.0.0:80")
		} else if connection13.String() != "[tcp] 127.0.0.1:1337 <- 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection13.String(), "[tcp] 127.0.0.1:1337 <- 0.0.0.0:80")
		}

		if connection14 == nil {
			t.Errorf("Expected %s to be valid", "[any] 127.0.0.1:1337 <- 0.0.0.0:80")
		} else if connection14.String() != "[any] 127.0.0.1:1337 <- 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection14.String(), "[any] 127.0.0.1:1337 <- 0.0.0.0:80")
		}

		if connection15 == nil {
			t.Errorf("Expected %s to be valid", "[udp] 127.0.0.1:1337 <- 0.0.0.0:80")
		} else if connection15.String() != "[udp] 127.0.0.1:1337 <- 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection15.String(), "[udp] 127.0.0.1:1337 <- 0.0.0.0:80")
		}

		if connection16 == nil {
			t.Errorf("Expected %s to be valid", "[tcp] *:1337 <- 0.0.0.0:80")
		} else if connection16.String() != "[tcp] 0.0.0.0:1337 <- 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection16.String(), "[tcp] 0.0.0.0:1337 <- 0.0.0.0:80")
		}

		if connection17 == nil {
			t.Errorf("Expected %s to be valid", "[any] *:1337 <- 0.0.0.0:80")
		} else if connection17.String() != "[any] 0.0.0.0:1337 <- 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection17.String(), "[any] 0.0.0.0:1337 <- 0.0.0.0:80")
		}

		if connection18 == nil {
			t.Errorf("Expected %s to be valid", "[udp] *:1337 <- 0.0.0.0:80")
		} else if connection18.String() != "[udp] 0.0.0.0:1337 <- 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection18.String(), "[udp] 0.0.0.0:1337 <- 0.0.0.0:80")
		}

	})

	t.Run("ParseConnection(peer)", func(t *testing.T) {

		connection1 := ParseConnection("[tcp] 127.0.0.1:1337 <-> 1.3.3.7:80")
		connection2 := ParseConnection("[any] 127.0.0.1:1337 <-> 1.3.3.7:80")
		connection3 := ParseConnection("[udp] 127.0.0.1:1337 <-> 1.3.3.7:80")

		connection4 := ParseConnection("[tcp] *:1337 <-> 1.3.3.7:80")
		connection5 := ParseConnection("[any] *:1337 <-> 1.3.3.7:80")
		connection6 := ParseConnection("[udp] *:1337 <-> 1.3.3.7:80")

		connection7 := ParseConnection("[tcp] 127.0.0.1:1337 <-> *:80")
		connection8 := ParseConnection("[any] 127.0.0.1:1337 <-> *:80")
		connection9 := ParseConnection("[udp] 127.0.0.1:1337 <-> *:80")

		connection10 := ParseConnection("[tcp] *:1337 <-> *:80")
		connection11 := ParseConnection("[any] *:1337 <-> *:80")
		connection12 := ParseConnection("[udp] *:1337 <-> *:80")

		connection13 := ParseConnection("[tcp] 127.0.0.1:1337 <-> 0.0.0.0:80")
		connection14 := ParseConnection("[any] 127.0.0.1:1337 <-> 0.0.0.0:80")
		connection15 := ParseConnection("[udp] 127.0.0.1:1337 <-> 0.0.0.0:80")

		connection16 := ParseConnection("[tcp] *:1337 <-> 0.0.0.0:80")
		connection17 := ParseConnection("[any] *:1337 <-> 0.0.0.0:80")
		connection18 := ParseConnection("[udp] *:1337 <-> 0.0.0.0:80")

		if connection1 == nil {
			t.Errorf("Expected %s to be valid", "[tcp] 127.0.0.1:1337 <-> 1.3.3.7:80")
		} else if connection1.String() != "[tcp] 127.0.0.1:1337 <-> 1.3.3.7:80" {
			t.Errorf("Expected %s to be %s", connection1.String(), "[tcp] 127.0.0.1:1337 <-> 1.3.3.7:80")
		}

		if connection2 == nil {
			t.Errorf("Expected %s to be valid", "[any] 127.0.0.1:1337 <-> 1.3.3.7:80")
		} else if connection2.String() != "[any] 127.0.0.1:1337 <-> 1.3.3.7:80" {
			t.Errorf("Expected %s to be %s", connection2.String(), "[any] 127.0.0.1:1337 <-> 1.3.3.7:80")
		}

		if connection3 == nil {
			t.Errorf("Expected %s to be valid", "[udp] 127.0.0.1:1337 <-> 1.3.3.7:80")
		} else if connection3.String() != "[udp] 127.0.0.1:1337 <-> 1.3.3.7:80" {
			t.Errorf("Expected %s to be %s", connection3.String(), "[udp] 127.0.0.1:1337 <-> 1.3.3.7:80")
		}

		if connection4 == nil {
			t.Errorf("Expected %s to be valid", "[tcp] *:1337 <-> 1.3.3.7:80")
		} else if connection4.String() != "[tcp] 0.0.0.0:1337 <-> 1.3.3.7:80" {
			t.Errorf("Expected %s to be %s", connection4.String(), "[tcp] 0.0.0.0:1337 <-> 1.3.3.7:80")
		}

		if connection5 == nil {
			t.Errorf("Expected %s to be valid", "[any] *:1337 <-> 1.3.3.7:80")
		} else if connection5.String() != "[any] 0.0.0.0:1337 <-> 1.3.3.7:80" {
			t.Errorf("Expected %s to be %s", connection5.String(), "[any] 0.0.0.0:1337 <-> 1.3.3.7:80")
		}

		if connection6 == nil {
			t.Errorf("Expected %s to be valid", "[udp] *:1337 <-> 1.3.3.7:80")
		} else if connection6.String() != "[udp] 0.0.0.0:1337 <-> 1.3.3.7:80" {
			t.Errorf("Expected %s to be %s", connection6.String(), "[udp] 0.0.0.0:1337 <-> 1.3.3.7:80")
		}

		if connection7 == nil {
			t.Errorf("Expected %s to be valid", "[tcp] 127.0.0.1:1337 <-> *:80")
		} else if connection7.String() != "[tcp] 127.0.0.1:1337 <-> 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection7.String(), "[tcp] 127.0.0.1:1337 <-> 0.0.0.0:80")
		}

		if connection8 == nil {
			t.Errorf("Expected %s to be valid", "[any] 127.0.0.1:1337 <-> *:80")
		} else if connection8.String() != "[any] 127.0.0.1:1337 <-> 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection8.String(), "[any] 127.0.0.1:1337 <-> 0.0.0.0:80")
		}

		if connection9 == nil {
			t.Errorf("Expected %s to be valid", "[udp] 127.0.0.1:1337 <-> *:80")
		} else if connection9.String() != "[udp] 127.0.0.1:1337 <-> 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection9.String(), "[udp] 127.0.0.1:1337 <-> 0.0.0.0:80")
		}

		if connection10 == nil {
			t.Errorf("Expected %s to be valid", "[tcp] *:1337 <-> *:80")
		} else if connection10.String() != "[tcp] 0.0.0.0:1337 <-> 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection10.String(), "[tcp] 0.0.0.0:1337 <-> 0.0.0.0:80")
		}

		if connection11 == nil {
			t.Errorf("Expected %s to be valid", "[any] *:1337 <-> *:80")
		} else if connection11.String() != "[any] 0.0.0.0:1337 <-> 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection11.String(), "[any] 0.0.0.0:1337 <-> 0.0.0.0:80")
		}

		if connection12 == nil {
			t.Errorf("Expected %s to be valid", "[udp] *:1337 <-> *:80")
		} else if connection12.String() != "[udp] 0.0.0.0:1337 <-> 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection12.String(), "[udp] 0.0.0.0:1337 <-> 0.0.0.0:80")
		}

		if connection13 == nil {
			t.Errorf("Expected %s to be valid", "[tcp] 127.0.0.1:1337 <-> 0.0.0.0:80")
		} else if connection13.String() != "[tcp] 127.0.0.1:1337 <-> 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection13.String(), "[tcp] 127.0.0.1:1337 <-> 0.0.0.0:80")
		}

		if connection14 == nil {
			t.Errorf("Expected %s to be valid", "[any] 127.0.0.1:1337 <-> 0.0.0.0:80")
		} else if connection14.String() != "[any] 127.0.0.1:1337 <-> 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection14.String(), "[any] 127.0.0.1:1337 <-> 0.0.0.0:80")
		}

		if connection15 == nil {
			t.Errorf("Expected %s to be valid", "[udp] 127.0.0.1:1337 <-> 0.0.0.0:80")
		} else if connection15.String() != "[udp] 127.0.0.1:1337 <-> 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection15.String(), "[udp] 127.0.0.1:1337 <-> 0.0.0.0:80")
		}

		if connection16 == nil {
			t.Errorf("Expected %s to be valid", "[tcp] *:1337 <-> 0.0.0.0:80")
		} else if connection16.String() != "[tcp] 0.0.0.0:1337 <-> 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection16.String(), "[tcp] 0.0.0.0:1337 <-> 0.0.0.0:80")
		}

		if connection17 == nil {
			t.Errorf("Expected %s to be valid", "[any] *:1337 <-> 0.0.0.0:80")
		} else if connection17.String() != "[any] 0.0.0.0:1337 <-> 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection17.String(), "[any] 0.0.0.0:1337 <-> 0.0.0.0:80")
		}

		if connection18 == nil {
			t.Errorf("Expected %s to be valid", "[udp] *:1337 <-> 0.0.0.0:80")
		} else if connection18.String() != "[udp] 0.0.0.0:1337 <-> 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection18.String(), "[udp] 0.0.0.0:1337 <-> 0.0.0.0:80")
		}

	})

	t.Run("ParseConnection(invalid)", func(t *testing.T) {

		// invalid protocol
		connection1 := ParseConnection("127.0.0.1:1337 -> 1.3.3.7:80")
		connection2 := ParseConnection("127.0.0.1:1337 <- 1.3.3.7:80")
		connection3 := ParseConnection("127.0.0.1:1337 <-> 1.3.3.7:80")

		// invalid source host
		connection4 := ParseConnection("[tcp] example.com:1337 -> 1.3.3.7:80")
		connection5 := ParseConnection("[any] example.com:1337 <- 1.3.3.7:80")
		connection6 := ParseConnection("[udp] example.com:1337 <-> 1.3.3.7:80")

		// invalid source port
		connection7 := ParseConnection("[tcp] *:99999 -> 1.3.3.7:80")
		connection8 := ParseConnection("[any] *:0 <- 1.3.3.7:80")
		connection9 := ParseConnection("[udp] *:-1 <-> 1.3.3.7:80")

		// invalid type
		connection10 := ParseConnection("[tcp] 127.0.0.1:1337 - 1.3.3.7:80")
		connection11 := ParseConnection("[any] 127.0.0.1:1337 _ 1.3.3.7:80")
		connection12 := ParseConnection("[udp] 127.0.0.1:1337 invalid 1.3.3.7:80")

		// invalid source port
		connection13 := ParseConnection("[tcp] 127.0.0.1:1337 -> 1.3.3.7:99999")
		connection14 := ParseConnection("[any] 127.0.0.1:1337 <- 1.3.3.7:0")
		connection15 := ParseConnection("[udp] 127.0.0.1:1337 <-> 1.3.3.7:-1")

		if connection1 != nil {
			t.Errorf("Expected %s to be nil", connection1.String())
		}

		if connection2 != nil {
			t.Errorf("Expected %s to be nil", connection2.String())
		}

		if connection3 != nil {
			t.Errorf("Expected %s to be nil", connection3.String())
		}

		if connection4 != nil {
			t.Errorf("Expected %s to be nil", connection4.String())
		}

		if connection5 != nil {
			t.Errorf("Expected %s to be nil", connection5.String())
		}

		if connection6 != nil {
			t.Errorf("Expected %s to be nil", connection6.String())
		}

		if connection7 != nil {
			t.Errorf("Expected %s to be nil", connection7.String())
		}

		if connection8 != nil {
			t.Errorf("Expected %s to be nil", connection8.String())
		}

		if connection9 != nil {
			t.Errorf("Expected %s to be nil", connection9.String())
		}

		if connection10 != nil {
			t.Errorf("Expected %s to be nil", connection10.String())
		}

		if connection11 != nil {
			t.Errorf("Expected %s to be nil", connection11.String())
		}

		if connection12 != nil {
			t.Errorf("Expected %s to be nil", connection12.String())
		}

		if connection13 != nil {
			t.Errorf("Expected %s to be nil", connection13.String())
		}

		if connection14 != nil {
			t.Errorf("Expected %s to be nil", connection14.String())
		}

		if connection15 != nil {
			t.Errorf("Expected %s to be nil", connection15.String())
		}

	})

	t.Run("SetSource()", func(t *testing.T) {

		connection1 := ToConnection("[tcp] 1.2.3.4:1234 -> 1.2.3.4:1234")
		connection2 := ToConnection("[any] 1.2.3.4:1234 <- 1.2.3.4:1234")
		connection3 := ToConnection("[udp] 1.2.3.4:1234 <-> 1.2.3.4:1234")

		connection1.SetSource(ToSocket("127.0.0.1:12345"))
		connection2.SetSource(ToSocket("*:80"))
		connection3.SetSource(ToSocket("0.0.0.0:8080"))

		if connection1.String() != "[tcp] 127.0.0.1:12345 -> 1.2.3.4:1234" {
			t.Errorf("Expected %s to be %s", connection1.String(), "[tcp] 127.0.0.1:12345 -> 1.2.3.4:1234")
		}

		if connection2.String() != "[any] 0.0.0.0:80 <- 1.2.3.4:1234" {
			t.Errorf("Expected %s to be %s", connection2.String(), "[any] 0.0.0.0:12345 <- 1.2.3.4:1234")
		}

		if connection3.String() != "[udp] 0.0.0.0:8080 <-> 1.2.3.4:1234" {
			t.Errorf("Expected %s to be %s", connection3.String(), "[udp] 0.0.0.0:8080 <-> 1.2.3.4:1234")
		}

	})

	t.Run("SetTarget()", func(t *testing.T) {

		connection1 := ToConnection("[tcp] 1.2.3.4:1234 -> 1.2.3.4:1234")
		connection2 := ToConnection("[any] 1.2.3.4:1234 <- 1.2.3.4:1234")
		connection3 := ToConnection("[udp] 1.2.3.4:1234 <-> 1.2.3.4:1234")

		connection1.SetTarget(ToSocket("127.0.0.1:12345"))
		connection2.SetTarget(ToSocket("*:80"))
		connection3.SetTarget(ToSocket("0.0.0.0:8080"))

		if connection1.String() != "[tcp] 1.2.3.4:1234 -> 127.0.0.1:12345" {
			t.Errorf("Expected %s to be %s", connection1.String(), "[tcp] 1.2.3.4:1234 -> 127.0.0.1:12345")
		}

		if connection2.String() != "[any] 1.2.3.4:1234 <- 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection2.String(), "[any] 1.2.3.4:1234 <- 0.0.0.0:80")
		}

		if connection3.String() != "[udp] 1.2.3.4:1234 <-> 0.0.0.0:8080" {
			t.Errorf("Expected %s to be %s", connection3.String(), "[udp] 1.2.3.4:1234 <-> 0.0.0.0:8080")
		}

	})

	t.Run("SetProtocol()", func(t *testing.T) {

		connection1 := ToConnection("[any] 1.2.3.4:1234 -> 1.2.3.4:1234")
		connection2 := ToConnection("[any] 1.2.3.4:1234 <- 1.2.3.4:1234")
		connection3 := ToConnection("[any] 1.2.3.4:1234 <-> 1.2.3.4:1234")

		connection1.SetProtocol(ProtocolTCP)
		connection2.SetProtocol(ProtocolANY)
		connection3.SetProtocol(ProtocolUDP)

		if connection1.Protocol != ProtocolTCP {
			t.Errorf("Expected %s to be %s", connection1.Protocol.String(), ProtocolTCP.String())
		}

		if connection2.Protocol != ProtocolANY {
			t.Errorf("Expected %s to be %s", connection2.Protocol.String(), ProtocolANY.String())
		}

		if connection3.Protocol != ProtocolUDP {
			t.Errorf("Expected %s to be %s", connection3.Protocol.String(), ProtocolUDP.String())
		}

	})

	t.Run("SetType()", func(t *testing.T) {

		connection1 := ToConnection("[any] 1.2.3.4:1234 <-> 1.2.3.4:1234")
		connection2 := ToConnection("[any] 1.2.3.4:1234 <-> 1.2.3.4:1234")
		connection3 := ToConnection("[any] 1.2.3.4:1234 <-> 1.2.3.4:1234")

		connection1.SetType("client")
		connection2.SetType("server")
		connection3.SetType("peer")

		if connection1.Type != "client" {
			t.Errorf("Expected %s to be %s", connection1.Type, "client")
		}

		if connection2.Type != "server" {
			t.Errorf("Expected %s to be %s", connection2.Type, "server")
		}

		if connection3.Type != "peer" {
			t.Errorf("Expected %s to be %s", connection3.Type, "peer")
		}

	})

	t.Run("String()", func(t *testing.T) {

		connection1 := ToConnection("[tcp] 127.0.0.1:1337 -> 1.3.3.7:80")
		connection2 := ToConnection("[any] 127.0.0.1:1337 <- *:80")
		connection3 := ToConnection("[udp] *:1337 <-> 0.0.0.0:80")

		if connection1.String() != "[tcp] 127.0.0.1:1337 -> 1.3.3.7:80" {
			t.Errorf("Expected %s to be %s", connection1.String(), "[tcp] 127.0.0.1:1337 -> 1.3.3.7:80")
		}

		if connection2.String() != "[any] 127.0.0.1:1337 <- 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection2.String(), "[any] 127.0.0.1:1337 <- 0.0.0.0:80")
		}

		if connection3.String() != "[udp] 0.0.0.0:1337 <-> 0.0.0.0:80" {
			t.Errorf("Expected %s to be %s", connection3.String(), "[udp] 0.0.0.0:1337 <-> 0.0.0.0:80")
		}

	})

}
