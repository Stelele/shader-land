<template>
    <div class="grid grid-rows-2">
        <div>
            <ShaderPlayground ref="shaderPlayground" :start-code="StartShaderFs">
                <SubmitShaderDetails v-if="isSignedIn" @on-submit="onSubmit" />
            </ShaderPlayground>
        </div>
        <div class="grid grid-cols-3">

        </div>
    </div>
</template>

<script lang="ts" setup>
import ShaderPlayground from '../components/ShaderPlayground.vue';
import SubmitShaderDetails from '../components/SubmitShaderDetails.vue';
import { StartShaderFs } from '../components/Renderer/Start.shader';
import { ref } from 'vue';
import { useAuth, useUser } from '@clerk/vue';
import { ShaderService } from '../services/ShaderService';
import { ShaderRequest } from '../services/types/ShaderServiceTypes';

const shaderPlayground = ref<InstanceType<typeof ShaderPlayground> | null>(null)

const { isSignedIn } = useUser()
const { getToken } = useAuth()

async function onSubmit(name: string, description: string) {
    const token = await getToken.value() ?? ""
    const request: ShaderRequest = {
        name,
        description,
        creationDate: new Date().getTime(),
        code: shaderPlayground.value?.getShaderCode() ?? "",
    }

    ShaderService.postShader(request, token)
}

</script>