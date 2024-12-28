export const StartShaderVs = /*wgsl*/ `
@vertex
fn vs(@builtin(vertex_index) idx: u32) -> @builtin(position) vec4f {
    let points: array<vec2f, 3> = array(
        vec2f(-1., 3.),
        vec2f(-1.,-1.),
        vec2f( 3.,-1.),
    );

    return vec4f(points[idx], .0, 1.);
}
`

export const Uniforms = /*wgsl*/`
@group(0) @binding(0) var<uniform> iResolution: vec3f;   // viewport resolution (in pixels)
@group(0) @binding(1) var<uniform> iTime: f32;           // shader playback time (in seconds)
@group(0) @binding(2) var<uniform> iTimeDelta: f32;      // rener time (in seconds)
@group(0) @binding(3) var<uniform> iFrameRate: f32;      // shader frame rate
@group(0) @binding(4) var<uniform> iFrame: u32;          // shader playback frame
`

export const StartShaderFs = /*wgsl*/`
@fragment
fn fs(@builtin(position) pos: vec4f) -> @location(0) vec4f {
    var uv = pos.xy / iResolution.xy;
    var col = 0.5 + 0.5 * cos(iTime + uv.xyx + vec3f(0., 2., 4.));
    return vec4f(col, 1.);
}
`