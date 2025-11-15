#version 330

in float x;

uniform float pos;
uniform mat4 projection;
//uniform float hz;

void main() {
    float y = 0.2f * pos;
    gl_Position = projection * vec4(x, y, 0.0, 1.0);
}