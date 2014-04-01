package graphics

type Color struct {
    R,G,B uint8
}

func (c *Color) Serialize() uint32 {
    return uint32(c.B) + uint32(c.G)  << 8 + uint32(c.R) << 16
}