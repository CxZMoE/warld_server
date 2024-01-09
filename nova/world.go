package nova

type World struct {
	Description       string
	Size              int64 // A world is a Squra Area
	ResourceAvailable int64
}

// Create a new world
func NewWorld(desc string, size int64, resourceAvailable int64) *World {
	return &World{
		Description:       desc,
		Size:              size,
		ResourceAvailable: resourceAvailable,
	}
}
