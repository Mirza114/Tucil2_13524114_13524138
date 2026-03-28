package main

import "fmt"

func ComputeBoundingBox(mesh Mesh) (BoundingBox, error) {
	if len(mesh.Vertices) == 0 {
		return BoundingBox{}, fmt.Errorf("mesh tidak memiliki vertex")
	}

	first := mesh.Vertices[0]
	bbox := BoundingBox{
		MinX: first.X, MinY: first.Y, MinZ: first.Z,
		MaxX: first.X, MaxY: first.Y, MaxZ: first.Z,
	}

	for _, v := range mesh.Vertices[1:] {
		if v.X < bbox.MinX {
			bbox.MinX = v.X
		}
		if v.Y < bbox.MinY {
			bbox.MinY = v.Y
		}
		if v.Z < bbox.MinZ {
			bbox.MinZ = v.Z
		}
		if v.X > bbox.MaxX {
			bbox.MaxX = v.X
		}
		if v.Y > bbox.MaxY {
			bbox.MaxY = v.Y
		}
		if v.Z > bbox.MaxZ {
			bbox.MaxZ = v.Z
		}
	}

	return bbox, nil
}

func BoundingBoxToCube(bbox BoundingBox) Cube {
	lengthX := bbox.MaxX - bbox.MinX
	lengthY := bbox.MaxY - bbox.MinY
	lengthZ := bbox.MaxZ - bbox.MinZ

	size := lengthX
	if lengthY > size {
		size = lengthY
	}
	if lengthZ > size {
		size = lengthZ
	}

	centerX := (bbox.MinX + bbox.MaxX) / 2.0
	centerY := (bbox.MinY + bbox.MaxY) / 2.0
	centerZ := (bbox.MinZ + bbox.MaxZ) / 2.0

	return Cube{
		CenterX: centerX,
		CenterY: centerY,
		CenterZ: centerZ,
		Size:    size,
	}
}

func CubeToBoundingBox(c Cube) BoundingBox {
	half := c.Size / 2.0
	return BoundingBox{
		MinX: c.CenterX - half,
		MinY: c.CenterY - half,
		MinZ: c.CenterZ - half,
		MaxX: c.CenterX + half,
		MaxY: c.CenterY + half,
		MaxZ: c.CenterZ + half,
	}
}

func IntersectsBox(a, b BoundingBox) bool {
	return !(a.MaxX < b.MinX || a.MinX > b.MaxX ||
		a.MaxY < b.MinY || a.MinY > b.MaxY ||
		a.MaxZ < b.MinZ || a.MinZ > b.MaxZ)
}

func FaceBoundingBox(face Face, vertices []Vertex) BoundingBox {
	v1 := vertices[face.A-1]
	v2 := vertices[face.B-1]
	v3 := vertices[face.C-1]

	minX := v1.X
	maxX := v1.X
	minY := v1.Y
	maxY := v1.Y
	minZ := v1.Z
	maxZ := v1.Z

	for _, v := range []Vertex{v2, v3} {
		if v.X < minX {
			minX = v.X
		}
		if v.X > maxX {
			maxX = v.X
		}
		if v.Y < minY {
			minY = v.Y
		}
		if v.Y > maxY {
			maxY = v.Y
		}
		if v.Z < minZ {
			minZ = v.Z
		}
		if v.Z > maxZ {
			maxZ = v.Z
		}
	}

	return BoundingBox{
		MinX: minX, MinY: minY, MinZ: minZ,
		MaxX: maxX, MaxY: maxY, MaxZ: maxZ,
	}
}