package model

import (
	"fmt"
	"strings"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geo"
)

const floatFmt = "%.6f"

type Geometry[Coord any] struct {
	Type        string `json:"type"`
	Coordinates Coord  `json:"coordinates"`
}
type Coordinates interface {
	Type() string
	ToWKT() string
}
type Point [2]float64

func (p Point) Type() string {
	return "Point"
}
func (p Point) ToWKT() string {
	return fmt.Sprintf(
		"POINT("+floatFmt+" "+floatFmt+")",
		p[0], p[1],
	)
}

// LineString represents a GeoJSON LineString
type LineString []Point

// Length calculates the total length of the LineString.
// It sums the Euclidean distances between consecutive points.
func (l LineString) Length() float64 {
	line := orb.LineString{}
	for _, v := range l {
		line = append(line, [2]float64(v))
	}
	return geo.Length(line)
}
func (l LineString) Type() string {
	return "LineString"
}
func (l LineString) ToWKT() string {
	if len(l) == 0 {
		return "LINESTRING EMPTY"
	}

	var parts []string
	for _, p := range l {
		parts = append(parts,
			fmt.Sprintf(floatFmt+" "+floatFmt, p[0], p[1]),
		)
	}

	return fmt.Sprintf("LINESTRING(%s)", strings.Join(parts, ", "))
}

// LinearRing represents a closed line (first == last), used in Polygons
type LinearRing []Point

// Polygon contains 1 outer LinearRing and 0+ holes
type Polygon []LinearRing

func (p Polygon) Type() string {
	return "Polygon"
}

func (p Polygon) ToWKT() string {
	if len(p) == 0 {
		return "POLYGON EMPTY"
	}

	var rings []string

	for _, ring := range p {
		if len(ring) == 0 {
			continue
		}

		ring = ensureClosed(ring)

		var pts []string
		for _, pt := range ring {
			pts = append(pts,
				fmt.Sprintf(floatFmt+" "+floatFmt, pt[0], pt[1]),
			)
		}

		rings = append(rings, fmt.Sprintf("(%s)", strings.Join(pts, ", ")))
	}

	if len(rings) == 0 {
		return "POLYGON EMPTY"
	}

	return fmt.Sprintf("POLYGON(%s)", strings.Join(rings, ", "))
}

func ensureClosed(ring LinearRing) LinearRing {
	if len(ring) < 2 {
		return ring
	}

	first := ring[0]
	last := ring[len(ring)-1]

	if first != last {
		ring = append(ring, first)
	}

	return ring
}

type MultiPoint []Point

func (m MultiPoint) Type() string {
	return "MultiPoint"
}
func (m MultiPoint) ToWKT() string {
	if len(m) == 0 {
		return "MULTIPOINT EMPTY"
	}

	var parts []string
	for _, p := range m {
		parts = append(parts,
			fmt.Sprintf("("+floatFmt+" "+floatFmt+")", p[0], p[1]),
		)
	}

	return fmt.Sprintf("MULTIPOINT(%s)", strings.Join(parts, ", "))
}

type MultiPolygon []Polygon

func (m MultiPolygon) Type() string {
	return "MultiPolygon"
}

func (m MultiPolygon) ToWKT() string {
	if len(m) == 0 {
		return "MULTIPOLYGON EMPTY"
	}

	var polygons []string

	for _, poly := range m {
		if len(poly) == 0 {
			continue
		}

		var rings []string

		for _, ring := range poly {
			if len(ring) == 0 {
				continue
			}

			ring = ensureClosed(ring)

			var pts []string
			for _, pt := range ring {
				pts = append(pts,
					fmt.Sprintf(floatFmt+" "+floatFmt, pt[0], pt[1]),
				)
			}

			rings = append(rings,
				fmt.Sprintf("(%s)", strings.Join(pts, ", ")),
			)
		}

		if len(rings) > 0 {
			polygons = append(polygons,
				fmt.Sprintf("(%s)", strings.Join(rings, ", ")),
			)
		}
	}

	if len(polygons) == 0 {
		return "MULTIPOLYGON EMPTY"
	}

	return fmt.Sprintf("MULTIPOLYGON(%s)", strings.Join(polygons, ", "))
}
