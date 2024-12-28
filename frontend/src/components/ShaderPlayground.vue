<template>
    <div class="grid grid-cols-5 gap-4">
        <div class="card bg-base-100 col-span-2 w-full shadow-xl">
            <div class="card-body pl-4 pr-1">
                <Renderer ref="renderer" />
            </div>
        </div>
        <div class="card bg-base-100 col-span-3 w-full shadow-xl">
            <div class="card-body rounded-2xl min-h-[70vh]">
                <Editor :starting-code="props.startCode" @on-value-change="onCodeChange" />
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import Renderer from './Renderer/Renderer.vue';
import Editor from './Editor.vue';

interface Props {
    startCode: string
}
const props = defineProps<Props>()

const renderer = ref<InstanceType<typeof Renderer> | null>(null)
function onCodeChange(code: string) {
    renderer.value?.loadFragmentShader(code)
}
</script>