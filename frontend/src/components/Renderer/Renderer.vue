<template>
    <div id="container">
        <canvas class="w-full h-full min-w-4 min-h-4" ref="webgpuCanvas"></canvas>
    </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue';
import { BufferInfo, Render } from './Render';

const emit = defineEmits<{
    onFrameRun: [BufferInfo]
}>()

const isReady = ref(false)
const webgpuCanvas = ref<HTMLCanvasElement | null>(null)
let render: Render
onMounted(() => {
    startRendering()
})

defineExpose({ loadFragmentShader, setFullScreen, toggleAnimation, resetAnimation, startRecording, stopRecording })

async function startRendering() {
    render = new Render(webgpuCanvas.value as HTMLCanvasElement)
    await render.init(onFrameRun, onReady)
}

function loadFragmentShader(code: string) {
    if (!isReady.value) {
        setTimeout(loadFragmentShader, undefined, code)
        return
    }

    render.loadFragmentShader(code)
}

function setFullScreen() {
    webgpuCanvas.value?.requestFullscreen()
}

function onFrameRun(info: BufferInfo) {
    emit('onFrameRun', info)
}

function toggleAnimation() {
    render.isPaused = !render.isPaused
}

function resetAnimation() {
    render.reset = true
}

function startRecording() {
    render.recordVideo(true)
}

function stopRecording() {
    render.recordVideo(false)
}

function onReady() {
    isReady.value = true
}
</script>