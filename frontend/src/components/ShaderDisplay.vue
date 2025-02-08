<template>
    <div class="flex flex-col w-full h-full">
        <div class="card card-compact bg-base-100 shadow-xl">
            <div class="card-body">
                <Renderer :class="[props.height]" class="w-full h-full" ref="renderer" />
                <div class="card-actions flex gap-0">
                    <div class="flex gap-4">
                        <div class="hover:cursor-pointer w-fit">
                            <OhVueIcon name="fa-step-backward" />
                        </div>
                        <div class="hover:cursor-pointer w-fit swap">
                            <input id="animCheck" type="checkbox" />
                            <OhVueIcon name="fa-play" class="swap-on" />
                            <OhVueIcon name="fa-pause" class="swap-off" />
                        </div>
                    </div>
                    <div class="flex-grow"></div>
                    <div class="flex gap-4">
                        <div class="hover:cursor-pointer w-fit swap">
                            <input id="recordCheck" type="checkbox" />
                            <OhVueIcon name="fa-regular-dot-circle" class="swap-on text-red-600" />
                            <OhVueIcon name="fa-regular-dot-circle" class="swap-off" />
                        </div>
                        <div class="hover:cursor-pointer">
                            <OhVueIcon name="fa-volume-up" />
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue'
import Renderer from './Renderer/Renderer.vue'
import { OhVueIcon } from 'oh-vue-icons'

export interface Props {
    shaderCode: string
    height?: string
}

const props = withDefaults(defineProps<Props>(), {
    height: "min-h-[45vh]"
})
const renderer = ref<InstanceType<typeof Renderer> | null>(null)

onMounted(() => {
    onStartUp()
})

function onStartUp() {
    renderer.value?.loadFragmentShader(props.shaderCode)
}

</script>