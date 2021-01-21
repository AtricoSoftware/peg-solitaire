package board

import (
	"github.com/atrico-go/console/box_drawing"
	"github.com/atrico-go/display"
	"github.com/atrico-go/display/rules"
	"github.com/atrico-go/display/tile"
)

// StringerMl
func (b Board) StringMl(params ...interface{}) []string {
	return b.Render().StringMl(params)
}

// Renderable
func (b Board) Render(_ ...tile.RenderRule) tile.Tile {
	intH := box_drawing.GetHorizontal(box_drawing.BoxSingle)
	intV := box_drawing.GetVertical(box_drawing.BoxSingle)
	extH := box_drawing.GetHorizontal(box_drawing.BoxHeavy)
	extV := box_drawing.GetVertical(box_drawing.BoxHeavy)
	centre := b.createSegment(Coord{2, 2}, Coord{4, 4}, &intH, &intH, &intV, &intV, &intH, &intV)
	top := b.createSegment(Coord{2, 0}, Coord{4, 1}, &extH, nil, &extV, &extV, &intH, &intV)
	bottom := b.createSegment(Coord{2, 5}, Coord{4, 6}, nil, &extH, &extV, &extV, &intH, &intV)
	left := b.createSegment(Coord{0, 2}, Coord{1, 4}, &extH, &extH, &extV, nil, &intH, &intV)
	right := b.createSegment(Coord{5, 2}, Coord{6, 4}, &extH, &extH, nil, &extV, &intH, &intV)
	return display.NewTableBuilder().
		AppendRow("", top, "").
		AppendRow(left, centre, right).
		AppendRow("", bottom, "").
		Build().
		Render(rules.UnicodeIntersections)
}

func (b Board) createSegment(min, max Coord, up, down, left, right, horiz, vert *rune) tile.Renderable {
	table := display.NewTableBuilder()
	if horiz != nil {
		table.WithHorizontalSeparator(*horiz)
	}
	if vert != nil {
		table.WithVerticalSeparator(*vert)
	}
	for y := min.Y; y <= max.Y; y++ {
		for x := min.X; x <= max.X; x++ {
			table.SetCell(x-min.X, y-min.Y, b.getPeg(Coord{x, y}))
		}
	}
	return display.NewBorder(table.Build(), up, down, left, right)
}

func (b Board) getPeg(pos Coord) string {
	if b.holes[pos] {
		return "O"
	}
	return " "
}
