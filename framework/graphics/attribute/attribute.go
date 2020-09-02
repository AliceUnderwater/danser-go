package attribute

import "github.com/go-gl/gl/v3.3-core/gl"

// Format defines names and types of OpenGL attributes (vertex format, uniform format, etc.).
//
// Example:
//   Format{{"position", Vec2}, {"color", Vec4}, {"texCoord": Vec2}}
type Format []VertexAttribute

// Size returns the total size of all attributes of the Format.
func (af Format) Size() int {
	total := 0
	for _, attr := range af {
		total += attr.Type.Size()
	}
	return total
}

// VertexAttribute represents an arbitrary OpenGL attribute, such as a vertex attribute or a shader
// uniform attribute.
type VertexAttribute struct {
	Name     string
	Type     Type
	Location uint32
}

// Type represents the type of an OpenGL attribute.
type Type int

// List of all possible attribute types.
const (
	Int   = Type(gl.INT)
	Float = Type(gl.FLOAT)
	Vec2  = Type(gl.FLOAT_VEC2)
	Vec3  = Type(gl.FLOAT_VEC3)
	Vec4  = Type(gl.FLOAT_VEC4)
	Mat2  = Type(gl.FLOAT_MAT2)
	Mat23 = Type(gl.FLOAT_MAT2x3)
	Mat24 = Type(gl.FLOAT_MAT2x4)
	Mat3  = Type(gl.FLOAT_MAT3)
	Mat32 = Type(gl.FLOAT_MAT3x2)
	Mat34 = Type(gl.FLOAT_MAT3x4)
	Mat4  = Type(gl.FLOAT_MAT4)
	Mat42 = Type(gl.FLOAT_MAT4x2)
	Mat43 = Type(gl.FLOAT_MAT4x3)
)

// Size returns the size of a type in bytes.
func (at Type) Size() int {
	switch at {
	case Int:
		return 4
	case Float:
		return 4
	case Vec2:
		return 2 * 4
	case Vec3:
		return 3 * 4
	case Vec4:
		return 4 * 4
	case Mat2:
		return 2 * 2 * 4
	case Mat23:
		return 2 * 3 * 4
	case Mat24:
		return 2 * 4 * 4
	case Mat3:
		return 3 * 3 * 4
	case Mat32:
		return 3 * 2 * 4
	case Mat34:
		return 3 * 4 * 4
	case Mat4:
		return 4 * 4 * 4
	case Mat42:
		return 4 * 2 * 4
	case Mat43:
		return 4 * 3 * 4
	default:
		panic("size of vertex attribute type: invalid type")
	}
}
