<template>
    <div id="container" class="w-full h-full">
        <canvas class="w-full h-full" ref="webgpu"></canvas>
    </div>
</template>

<script lang="ts" setup>
import { onMounted } from 'vue';
import { Render } from './Render';

let render: Render
onMounted(() => {
    startRendering()
})

defineExpose({ loadFragmentShader })

async function startRendering() {
    render = new Render(document.querySelector("canvas") as HTMLCanvasElement)
    await render.init()
}

function loadFragmentShader(code: string) {
    if (render) {
        render.loadFragmentShader(code)
    }
}
</script>