<template>
    <div class="grid grid-rows-2">
        <div>
            <ShaderPlayground ref="shaderPlayground" :start-code="StartShaderFs">
                <SubmitShaderDetails v-if="isAuthenticated" @on-submit="onSubmit" />
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
import { useAuth0 } from '@auth0/auth0-vue';
import { ref } from 'vue';
import { ShaderRequest } from '../services/types/ShaderServiceTypes';
import { getDisplayName } from '../helpers/getDisplayName';
import { ShaderService } from '../services/ShaderService';
import { useRouter } from 'vue-router';

const { isAuthenticated, user, getAccessTokenSilently } = useAuth0()
const router = useRouter()
const shaderPlayground = ref<InstanceType<typeof ShaderPlayground> | null>(null)

async function onSubmit(name: string, description: string) {
    const code = shaderPlayground.value?.getShaderCode() ?? ""
    const accessToken = await getAccessTokenSilently()
    const data: ShaderRequest = {
        name,
        description,
        code,
        tags: "",
        email: user.value?.email ?? "",
        userName: getDisplayName(user),
        creationDate: new Date().getTime()
    }

    const response = (await ShaderService.postShader(data, accessToken)).data
    // router.push(`/view/${response.id}`)
}

</script>