<template>
    <div class="w-full h-full">
        <canvas class="w-full h-full" ref="webgpu"></canvas>
    </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue';
import { Render } from './Render';

const webgpu = ref<HTMLCanvasElement | null>(null)
onMounted(() => {
    startRendering()
})

async function startRendering() {
    const render = new Render(webgpu.value as HTMLCanvasElement)
    await render.init()
}

</script>