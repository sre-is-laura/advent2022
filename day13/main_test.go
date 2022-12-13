package main

import (
	"testing"
)

func TestPacketOrder(t *testing.T) {
	type test struct {
		left  string
		right string
		want  bool // true if packets in order
	}

	tests := []test{
		{left: "[1,1,3,1,1]", right: "[1,1,5,1,1]", want: true},
		{left: "[1,1,6,1,1]", right: "[1,1,5,1,1]", want: false},
		{left: "[[1],[2,3,4]]", right: "[[1],4]", want: true},
		{left: "[[1],[5,3,4]]", right: "[[1],4]", want: false},
		{left: "[[1],[5,3,4],6]", right: "[[1],4]", want: false},
		{left: "[3]", right: "[[8,7,6]]", want: true},
		{left: "[1,[2,[3,[4,[5,6,7]]]],8,9]", right: "[1,[2,[3,[4,[5,6,0]]]],8,9]", want: false},
		{left: "[[2,4,6,[7,[0,4,8,0],[4,0],[0,6,10],7]],[1,4,[[5,8,6,9,2],[1,8]]],[4,3,[10],5],[8,[3,7],9,3,10],[6,2,[[0,6],5,[9],10],9,[]]]", right: "[[[[10,6,9,1,5],[5],8],7,[[5,3],10,[10,10,6,1,7],[],[6,2,6,7,4]],[4,[0,1],[9,3],[9,8,9,0,8]],[7,[6,6,2]]],[8,4],[0]]", want: true},
	}

	for _, tc := range tests {
		l := Packet{tc.left}
		r := Packet{tc.right}
		pair := PacketPair{l, r}

		result := pair.inOrder()
		if result != tc.want {
			t.Fatalf("expected: %v, got: %v", tc.want, result)
		}
	}
}
