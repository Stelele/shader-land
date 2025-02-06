<template>
    <div class="grid grid-cols-3 md:grid-cols-5 gap-2">
        <div class="flex flex-col gap-4 col-span-2">
            <div class="card card-compact bg-base-100 md:col-span-2 shadow-xl">
                <div class="card-body">
                    <Renderer ref="renderer" />
                    <div class="card-actions justify-end">
                        <div class="hover:cursor-pointer" @click="onFullScreen">
                            <span class="material-icons-outlined">fullscreen</span>
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

interface Props {
    startCode: string
}
const props = defineProps<Props>()
const shaderCode = ref<string>(props.startCode)


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
</script>