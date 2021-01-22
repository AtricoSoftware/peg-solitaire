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
	centre := b.createSegment(Position{-1, -1}, Position{+1, +1}, &intH, &intH, &intV, &intV, &intH, &intV)
	top := b.createSegment(Position{-1, -3}, Position{+1, -2}, &extH, nil, &extV, &extV, &intH, &intV)
	bottom := b.createSegment(Position{-1, +2}, Position{+1, +3}, nil, &extH, &extV, &extV, &intH, &intV)
	left := b.createSegment(Position{-3, -1}, Position{-2, +1}, &extH, &extH, &extV, nil, &intH, &intV)
	right := b.createSegment(Position{+2, -1}, Position{+3, +1}, &extH, &extH, nil, &extV, &intH, &intV)
	return display.NewTableBuilder().
		AppendRow("", top, "").
		AppendRow(left, centre, right).
		AppendRow("", bottom, "").
		Build().
		Render(rules.UnicodeIntersections)
}

func (b Board) createSegment(min, max Position, up, down, left, right, horiz, vert *rune) tile.Renderable {
	table := display.NewTableBuilder()
	if horiz != nil {
		table.WithHorizontalSeparator(*horiz)
	}
	if vert != nil {
		table.WithVerticalSeparator(*vert)
	}
	for y := min.Y; y <= max.Y; y++ {
		for x := min.X; x <= max.X; x++ {
			table.SetCell(x-min.X, y-min.Y, b.getPeg(Position{x, y}))
		}
	}
	return display.NewBorder(table.Build(), up, down, left, right)
}

func (b Board) getPeg(pos Position) string {
	if b.holes[pos] {
		return "O"
	}
	return " "
}
