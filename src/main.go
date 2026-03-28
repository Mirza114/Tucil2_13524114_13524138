package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <path_file.obj> <max_depth>")
		return
	}

	inputPath := os.Args[1]
	maxDepth, err := strconv.Atoi(os.Args[2])
	if err != nil || maxDepth < 0 {
		fmt.Println("Error: max_depth harus bilangan bulat >= 0")
		return
	}

	start := time.Now()

	mesh, err := ParseOBJ(inputPath)
	if err != nil {
		fmt.Println("Error parsing OBJ:", err)
		return
	}

	stats := NewStats(maxDepth)
	root := BuildOctree(mesh, maxDepth, stats)

	var voxels []Cube
	CollectActiveLeafVoxels(root, &voxels)

	base := filepath.Base(inputPath)
	ext := filepath.Ext(base)
	name := base[:len(base)-len(ext)]
	outputPath := filepath.Join(filepath.Dir(inputPath), name+"-voxelized.obj")

	err = WriteVoxelOBJ(outputPath, voxels, stats)
	if err != nil {
		fmt.Println("Error menulis file output:", err)
		return
	}

	duration := time.Since(start)

	fmt.Println("=== HASIL VOXELIZATION ===")
	fmt.Printf("Input file                  : %s\n", inputPath)
	fmt.Printf("Banyak voxel yang terbentuk : %d\n", stats.VoxelCount)
	fmt.Printf("Banyak vertex yang terbentuk: %d\n", stats.OutputVertexCount)
	fmt.Printf("Banyak faces yang terbentuk : %d\n", stats.OutputFaceCount)

	fmt.Println("\n=== Statistik node octree yang terbentuk ===")
	for d := 0; d <= maxDepth; d++ {
		fmt.Printf("%d : %d\n", d, stats.NodesPerDepth[d])
	}

	fmt.Println("\n=== Statistik node yang tidak perlu ditelusuri ===")
	for d := 0; d <= maxDepth; d++ {
		fmt.Printf("%d : %d\n", d, stats.PrunedPerDepth[d])
	}

	fmt.Printf("\nKedalaman octree            : %d\n", maxDepth)
	fmt.Printf("Lama waktu program berjalan : %s\n", duration)
	fmt.Printf("Path file output            : %s\n", outputPath)
}