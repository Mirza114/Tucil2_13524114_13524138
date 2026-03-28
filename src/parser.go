package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseOBJ(path string) (Mesh, error) {
	file, err := os.Open(path)
	if err != nil {
		return Mesh{}, fmt.Errorf("gagal membuka file: %w", err)
	}
	defer file.Close()

	var mesh Mesh
	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}

		switch fields[0] {
		case "v":
			if len(fields) != 4 {
				return Mesh{}, fmt.Errorf("baris %d tidak valid untuk vertex", lineNumber)
			}

			x, err := strconv.ParseFloat(fields[1], 64)
			if err != nil {
				return Mesh{}, fmt.Errorf("baris %d: x tidak valid", lineNumber)
			}
			y, err := strconv.ParseFloat(fields[2], 64)
			if err != nil {
				return Mesh{}, fmt.Errorf("baris %d: y tidak valid", lineNumber)
			}
			z, err := strconv.ParseFloat(fields[3], 64)
			if err != nil {
				return Mesh{}, fmt.Errorf("baris %d: z tidak valid", lineNumber)
			}

			mesh.Vertices = append(mesh.Vertices, Vertex{X: x, Y: y, Z: z})

		case "f":
			if len(fields) != 4 {
				return Mesh{}, fmt.Errorf("baris %d tidak valid untuk face", lineNumber)
			}

			a, err := strconv.Atoi(fields[1])
			if err != nil {
				return Mesh{}, fmt.Errorf("baris %d: indeks face A tidak valid", lineNumber)
			}
			b, err := strconv.Atoi(fields[2])
			if err != nil {
				return Mesh{}, fmt.Errorf("baris %d: indeks face B tidak valid", lineNumber)
			}
			c, err := strconv.Atoi(fields[3])
			if err != nil {
				return Mesh{}, fmt.Errorf("baris %d: indeks face C tidak valid", lineNumber)
			}

			if a <= 0 || b <= 0 || c <= 0 {
				return Mesh{}, fmt.Errorf("baris %d: indeks face harus > 0", lineNumber)
			}

			mesh.Faces = append(mesh.Faces, Face{A: a, B: b, C: c})

		default:
			return Mesh{}, fmt.Errorf("baris %d tidak valid: hanya 'v x y z' atau 'f i j k' yang didukung", lineNumber)
		}
	}

	if err := scanner.Err(); err != nil {
		return Mesh{}, fmt.Errorf("gagal membaca file: %w", err)
	}

	if len(mesh.Vertices) == 0 {
		return Mesh{}, fmt.Errorf("file tidak memiliki vertex")
	}
	if len(mesh.Faces) == 0 {
		return Mesh{}, fmt.Errorf("file tidak memiliki face")
	}

	// validasi indeks face
	for i, face := range mesh.Faces {
		n := len(mesh.Vertices)
		if face.A > n || face.B > n || face.C > n {
			return Mesh{}, fmt.Errorf("face ke-%d mereferensikan vertex di luar range", i+1)
		}
	}

	return mesh, nil
}