package graph

import (
	"reflect"
	"testing"
)

var (
	// Graph1 has two 8-Cliques in it, vertext 0-7 and vertext 8-15. No edges exist
	// between two cliques.
	Graph1 = [][]int{
		//  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15
		{1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}, //  0
		{1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}, //  1
		{1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}, //  2
		{1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}, //  3
		{1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}, //  4
		{1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}, //  5
		{1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}, //  6
		{1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}, //  7
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1}, //  8
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1}, //  9
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1}, // 10
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1}, // 11
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1}, // 12
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1}, // 13
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1}, // 14
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1}, // 15
	}

	// Graph2 has two 4-Cliques in it, vertext 0-3 and vertext 4-7. There are two
	// extra edges between cliques, (1，4) and (3,6).
	//  0-----1 ----- 4-----5
	//  | \ / |       | \ / |
	//  |  X  |       |  X  |
	//  | / \ |       | / \ |
	//  2-----3 ----- 6-----7
	Graph2 = [][]int{
		//  1  2  3  4  5  6  7
		{1, 1, 1, 1, 0, 0, 0, 0}, //  0
		{1, 1, 1, 1, 1, 0, 0, 0}, //  1
		{1, 1, 1, 1, 0, 0, 0, 0}, //  2
		{1, 1, 1, 1, 0, 0, 1, 0}, //  3
		{1, 0, 0, 0, 1, 1, 1, 1}, //  4
		{0, 0, 0, 0, 1, 1, 1, 1}, //  5
		{0, 0, 0, 1, 1, 1, 1, 1}, //  6
		{0, 0, 0, 0, 1, 1, 1, 1}, //  7
	}
)

func haveEdgeInGraph1(from, to int) bool {
	return Graph1[from][to] == 1
}

func haveEdgeInGraph2(from, to int) bool {
	return Graph2[from][to] == 1
}

func TestIsClique(t *testing.T) {
	type args struct {
		vertices []int
		haveEdge func(from, to int) bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Graph1[0-3] is Clique",
			args: args{
				vertices: []int{0, 1, 2, 3},
				haveEdge: haveEdgeInGraph1,
			},
			want: true,
		},
		{
			name: "Graph1[0-7] is Clique",
			args: args{
				vertices: []int{0, 1, 2, 3, 4, 5, 6, 7},
				haveEdge: haveEdgeInGraph1,
			},
			want: true,
		},
		{
			name: "Graph1[6-8] is not Clique",
			args: args{
				vertices: []int{6, 7, 8, 9},
				haveEdge: haveEdgeInGraph1,
			},
			want: false,
		},
		{
			name: "Graph2[0-4] is not Clique",
			args: args{
				vertices: []int{0, 1, 2, 3, 4},
				haveEdge: haveEdgeInGraph2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsClique(tt.args.vertices, tt.args.haveEdge); got != tt.want {
				t.Errorf("IsClique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsKClique(t *testing.T) {
	type args struct {
		vertices []int
		k        int
		haveEdge func(from, to int) bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Graph1[0-3] is 4-Clique",
			args: args{
				vertices: []int{0, 1, 2, 3},
				k:        4,
				haveEdge: haveEdgeInGraph1,
			},
			want: true,
		},
		{
			name: "Graph1[0-8] is not 8-Clique",
			args: args{
				vertices: []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
				k:        8,
				haveEdge: haveEdgeInGraph1,
			},
			want: false,
		},
		{
			name: "Graph1[6-9] is not 4-Clique",
			args: args{
				vertices: []int{6, 7, 8, 9},
				k:        4,
				haveEdge: haveEdgeInGraph1,
			},
			want: false,
		},
		{
			name: "Graph1[0] is 1-Clique",
			args: args{
				vertices: []int{0},
				k:        1,
				haveEdge: haveEdgeInGraph1,
			},
			want: true,
		},
		{
			name: "Graph1[10] is 1-Clique",
			args: args{
				vertices: []int{10},
				k:        1,
				haveEdge: haveEdgeInGraph1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsKClique(tt.args.vertices, tt.args.k, tt.args.haveEdge); got != tt.want {
				t.Errorf("IsKClique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindKClique(t *testing.T) {
	type args struct {
		vertices []int
		k        int
		haveEdge func(from, to int) bool
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Find 8-Clique in Graph1[0-7]",
			args: args{
				vertices: []int{0, 1, 2, 3, 4, 5, 6, 7},
				k:        8,
				haveEdge: haveEdgeInGraph1,
			},
			want: [][]int{{0, 1, 2, 3, 4, 5, 6, 7}},
		},
		{
			name: "Find two 4-Cliques in Graph2",
			args: args{
				vertices: []int{0, 1, 2, 3, 4, 5, 6, 7},
				k:        4,
				haveEdge: haveEdgeInGraph2,
			},
			want: [][]int{{0, 1, 2, 3}, {4, 5, 6, 7}},
		},
		{
			name: "Find no 5-Cliques in Graph2",
			args: args{
				vertices: []int{0, 1, 2, 3, 4, 5, 6, 7},
				k:        5,
				haveEdge: haveEdgeInGraph2,
			},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindKClique(tt.args.vertices, tt.args.k, tt.args.haveEdge); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindClique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindConnectedGraph(t *testing.T) {
	type args struct {
		vertices []int
		haveEdge func(from, to int) bool
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Find two ConnectedGraph in Graph1[4-11]",
			args: args{
				vertices: []int{4, 5, 6, 7, 8, 9, 10, 11},
				haveEdge: haveEdgeInGraph1,
			},
			want: [][]int{{4, 5, 6, 7}, {8, 9, 10, 11}},
		},
		{
			name: "Find one ConnectedGraph in Graph2",
			args: args{
				vertices: []int{0, 1, 2, 3, 4, 5, 6, 7},
				haveEdge: haveEdgeInGraph2,
			},
			want: [][]int{{0, 1, 2, 3, 4, 5, 6, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindConnectedGraph(tt.args.vertices, tt.args.haveEdge); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindConnectedGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}
