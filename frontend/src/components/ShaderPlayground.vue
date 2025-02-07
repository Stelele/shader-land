<template>
    <div class="grid grid-cols-3 md:grid-cols-5 gap-2">
        <div class="flex flex-col gap-4 col-span-2">
            <div class="card card-compact bg-base-100 md:col-span-2 shadow-xl">
                <div class="card-body">
                    <Renderer ref="renderer" @on-frame-run="onRenderUpdate" />
                    <div class="card-actions flex gap-0">
                        <div class="flex gap-4">
                            <div @click="resetAnimation" class="hover:cursor-pointer w-fit">
                                <OhVueIcon name="fa-step-backward" />
                            </div>
                            <div @click="toggleAnimation" class="hover:cursor-pointer w-fit swap">
                                <input id="animCheck" type="checkbox" />
                                <OhVueIcon name="fa-play" class="swap-on" />
                                <OhVueIcon name="fa-pause" class="swap-off" />
                            </div>
                            <div id="time"></div>
                            <div id="fps"></div>
                            <div id="resolution"></div>
                        </div>
                        <div class="flex-grow"></div>
                        <div class="flex gap-4">
                            <div @click="recordAnimation" class="hover:cursor-pointer w-fit swap">
                                <input id="recordCheck" type="checkbox" />
                                <OhVueIcon name="fa-regular-dot-circle" class="swap-on text-red-600" />
                                <OhVueIcon name="fa-regular-dot-circle" class="swap-off" />
                            </div>
                            <div class="hover:cursor-pointer">
                                <OhVueIcon name="fa-volume-up" />
                            </div>
                            <div class="hover:cursor-pointer w-fit" @click="onFullScreen">
                                <OhVueIcon name="fa-expand" />
                            </div>
                        </div>
                    </div>
                    <slot></slot>
                </div>
            </div>
        </div>
        <div class="card card-compact bg-base-100 col-span-3 md:col-span-3 shadow-xl">
            <div class="card-body">
                <Editor :starting-code="props.startCode" @on-value-change="onCodeChange" />
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import Renderer from './Renderer/Renderer.vue'
import Editor from './Editor.vue'
import { BufferInfo } from './Renderer/Render'
import { OhVueIcon } from 'oh-vue-icons'

interface Props {
    startCode: string
}
const props = defineProps<Props>()

const shaderCode = ref<string>(props.startCode)
const isRecording = ref(false)

defineExpose({ getShaderCode })

const renderer = ref<InstanceType<typeof Renderer> | null>(null)
function onCodeChange(code: string) {
    shaderCode.value = code
    renderer.value?.loadFragmentShader(shaderCode.value)
}

function onFullScreen() {
    renderer.value?.setFullScreen()
}

function getShaderCode() {
    return shaderCode.value
}

function onRenderUpdate(info: BufferInfo) {
    const timeOption = document.getElementById('time')
    if (timeOption) {
        timeOption.innerHTML = info.time.toFixed(2)
    }

    const fpsOption = document.getElementById('fps')
    if (fpsOption) {
        fpsOption.innerHTML = `${info.frameRate.toFixed(1)} fps`
    }

    const resolutionOption = document.getElementById('resolution')
    if (resolutionOption) {
        resolutionOption.innerHTML = `${info.resolution[0]} x ${info.resolution[1]}`
    }
}

function toggleAnimation() {
    const swap = document.getElementById("animCheck") as HTMLInputElement
    swap.checked = !swap.checked
    renderer.value?.toggleAnimation()
}

function resetAnimation() {
    renderer.value?.resetAnimation()
}

function recordAnimation() {
    const record = document.getElementById("recordCheck") as HTMLInputElement
    record.checked = !record.checked
    isRecording.value = !isRecording.value

    if (isRecording.value) {
        renderer.value?.startRecording()
    } else {
        renderer.value?.stopRecording()
    }

}
</script>