<template>
    <div id="container">
        <canvas class="w-full min-w-4 h-[45vh]" ref="webgpuCanvas"></canvas>
    </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue';
import { BufferInfo, Render } from './Render';

const emit = defineEmits<{
    onFrameRun: [BufferInfo]
}>()

const webgpuCanvas = ref<HTMLCanvasElement | null>(null)
let render: Render
onMounted(() => {
    startRendering()
})

defineExpose({ loadFragmentShader, setFullScreen })

async function startRendering() {
    render = new Render(webgpuCanvas.value as HTMLCanvasElement)
    await render.init(onFrameRun)
}

function loadFragmentShader(code: string) {
    if (render) {
        render.loadFragmentShader(code)
    }
}

function setFullScreen() {
    webgpuCanvas.value?.requestFullscreen()
}

function onFrameRun(info: BufferInfo) {
    emit('onFrameRun', info)
}
</script>