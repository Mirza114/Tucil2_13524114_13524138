package main

type Vertex struct {
	X float64
	Y float64
	Z float64
}

type Face struct {
	A int
	B int
	C int
}

type Mesh struct {
	Vertices []Vertex
	Faces    []Face
}

type BoundingBox struct {
	MinX float64
	MinY float64
	MinZ float64
	MaxX float64
	MaxY float64
	MaxZ float64
}

type Cube struct {
	CenterX float64
	CenterY float64
	CenterZ float64
	Size    float64
}

type OctreeNode struct {
	Cube     Cube
	Depth    int
	IsLeaf   bool
	Active   bool
	Children [8]*OctreeNode
}