package main

import (
	"fmt"
	"os"
)

func WriteVoxelOBJ(path string, voxels []Cube, stats *Stats) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("gagal membuat file output: %w", err)
	}
	defer file.Close()

	vertexOffset := 0

	for _, cube := range voxels {
		verts := CubeVertices(cube)

		for _, v := range verts {
			_, err := fmt.Fprintf(file, "v %f %f %f\n", v.X, v.Y, v.Z)
			if err != nil {
				return err
			}
		}

		faces := [12][3]int{
			{1, 2, 3}, {1, 3, 4}, // bawah
			{5, 6, 7}, {5, 7, 8}, // atas
			{1, 2, 6}, {1, 6, 5}, // depan
			{2, 3, 7}, {2, 7, 6}, // kanan
			{3, 4, 8}, {3, 8, 7}, // belakang
			{4, 1, 5}, {4, 5, 8}, // kiri
		}

		for _, f := range faces {
			_, err := fmt.Fprintf(file, "f %d %d %d\n",
				vertexOffset+f[0],
				vertexOffset+f[1],
				vertexOffset+f[2],
			)
			if err != nil {
				return err
			}
		}

		vertexOffset += 8
	}

	stats.OutputVertexCount = len(voxels) * 8
	stats.OutputFaceCount = len(voxels) * 12
	stats.OutputPath = path

	return nil
}

func CubeVertices(c Cube) [8]Vertex {
	h := c.Size / 2.0

	return [8]Vertex{
		{c.CenterX - h, c.CenterY - h, c.CenterZ - h},
		{c.CenterX + h, c.CenterY - h, c.CenterZ - h},
		{c.CenterX + h, c.CenterY + h, c.CenterZ - h},
		{c.CenterX - h, c.CenterY + h, c.CenterZ - h},
		{c.CenterX - h, c.CenterY - h, c.CenterZ + h},
		{c.CenterX + h, c.CenterY - h, c.CenterZ + h},
		{c.CenterX + h, c.CenterY + h, c.CenterZ + h},
		{c.CenterX - h, c.CenterY + h, c.CenterZ + h},
	}
}