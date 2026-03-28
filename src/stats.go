package main

type Stats struct {
	VoxelCount         int
	OutputVertexCount  int
	OutputFaceCount    int
	NodesPerDepth      map[int]int
	PrunedPerDepth     map[int]int
	MaxDepth           int
	OutputPath         string
}

func NewStats(maxDepth int) *Stats {
	return &Stats{
		NodesPerDepth:  make(map[int]int),
		PrunedPerDepth: make(map[int]int),
		MaxDepth:       maxDepth,
	}
}