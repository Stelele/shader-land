export const StartShaderVs = /*wgsl*/ `
@vertex
fn vs(@builtin(vertex_index) idx: u32) -> @builtin(position) vec4f {
    let points: array<vec2f, 3> = array(
        vec2f(-1., 3.),
        vec2f(-1.,-1.),
        vec2f( 3.,-1.),
    );

    return vec4f(points[idx], .0, 1.);
}`

export const StartShaderFs = /*wgsl*/`



@fragment
fn fs() -> @location(0) vec4f {
    return vec4f(1., .0, .0, 1.);
}`