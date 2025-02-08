<template>
    <div v-if="clerk?.loaded" class="w-full h-[100vh] grid grid-cols-4 gap-2 p-4">
        <div class="col-span-2 w-full h-[40vh]">
            <ShaderDisplay height="min-h-[40vh]" :shader-code="StartShaderFs" />
        </div>
        <div class="w-full h-[40vh] items-start justify-center col-span-2 gap-2 flex flex-col">
            <div class="skeleton w-1/2 h-5"></div>
            <div class="skeleton w-1/2 h-5"></div>
            <div class="skeleton w-3/4 h-5"></div>
            <div class="skeleton w-3/4 h-5"></div>
            <div class="skeleton w-full h-5"></div>
            <div class="skeleton w-full h-5"></div>
        </div>
        <div class="col-span-4 grid grid-cols-4 gap-2 p-2">
            <div class="w-full">
                <ShaderDisplay height="min-h-[25vh]" :shader-code="StartShaderFs" />
            </div>
            <div class="w-full">
                <ShaderDisplay height="min-h-[25vh]" :shader-code="StartShaderFs" />
            </div>
            <div class="w-full">
                <ShaderDisplay height="min-h-[25vh]" :shader-code="StartShaderFs" />
            </div>
            <div class="w-full">
                <ShaderDisplay height="min-h-[25vh]" :shader-code="StartShaderFs" />
            </div>
        </div>
    </div>
    <HomeSkeletonLoader v-else />
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue'
import { useClerk } from '@clerk/vue'
import { ShaderService } from '../services/ShaderService'
import HomeSkeletonLoader from '../views/HomeSkeletonLoader.vue'
import ShaderDisplay from '../components/ShaderDisplay.vue'
import { StartShaderFs } from '../components/Renderer/Start.shader'

onMounted(() => {
    init()
})

const shaders = ref()
const clerk = useClerk()

async function init() {
    shaders.value = await ShaderService.getShaders()
}

</script>