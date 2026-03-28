package main

func BuildOctree(mesh Mesh, maxDepth int, stats *Stats) *OctreeNode {
	bbox, _ := ComputeBoundingBox(mesh)
	rootCube := BoundingBoxToCube(bbox)

	root := &OctreeNode{
		Cube:   rootCube,
		Depth:  0,
		IsLeaf: true,
		Active: false,
	}

	buildRecursive(root, mesh, maxDepth, stats)
	return root
}

func buildRecursive(node *OctreeNode, mesh Mesh, maxDepth int, stats *Stats) {
	stats.NodesPerDepth[node.Depth]++

	if !CubeIntersectsSurface(node.Cube, mesh) {
		stats.PrunedPerDepth[node.Depth]++
		return
	}

	if node.Depth == maxDepth {
		node.Active = true
		node.IsLeaf = true
		stats.VoxelCount++
		return
	}

	node.IsLeaf = false
	children := Subdivide(node.Cube)

	for i := 0; i < 8; i++ {
		child := &OctreeNode{
			Cube:   children[i],
			Depth:  node.Depth + 1,
			IsLeaf: true,
			Active: false,
		}
		node.Children[i] = child
		buildRecursive(child, mesh, maxDepth, stats)
	}
}

func CubeIntersectsSurface(cube Cube, mesh Mesh) bool {
	cubeBox := CubeToBoundingBox(cube)

	for _, face := range mesh.Faces {
		faceBox := FaceBoundingBox(face, mesh.Vertices)
		if IntersectsBox(cubeBox, faceBox) {
			return true
		}
	}

	return false
}

func Subdivide(c Cube) [8]Cube {
	var children [8]Cube
	childSize := c.Size / 2.0
	offset := c.Size / 4.0

	index := 0
	for _, dx := range []float64{-offset, offset} {
		for _, dy := range []float64{-offset, offset} {
			for _, dz := range []float64{-offset, offset} {
				children[index] = Cube{
					CenterX: c.CenterX + dx,
					CenterY: c.CenterY + dy,
					CenterZ: c.CenterZ + dz,
					Size:    childSize,
				}
				index++
			}
		}
	}

	return children
}

func CollectActiveLeafVoxels(node *OctreeNode, voxels *[]Cube) {
	if node == nil {
		return
	}

	if node.IsLeaf && node.Active {
		*voxels = append(*voxels, node.Cube)
		return
	}

	for _, child := range node.Children {
		if child != nil {
			CollectActiveLeafVoxels(child, voxels)
		}
	}
}