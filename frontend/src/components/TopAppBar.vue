<template>
    <nav>
        <div class="navbar bg-base-100">
            <div class="flex-1 gap-4">
                <RouterLink to="/" class="btn btn-ghost hover:text-blue-600 hover:bg-transparent">Shader Land
                </RouterLink>
                <div class="form-control">
                    <input type="text" placeholder="Search" class="input input-bordered w-24 md:w-auto" />
                </div>
            </div>

            <div class="flex-none">
                <ul class="menu menu-horizontal px-1">
                    <span class="menu menu-horizontal mr-0 pr-0">Welcome</span>
                    <li v-if="isAuthenticated" class="hover:text-blue-600">
                        <RouterLink to="" class="hover:bg-transparent">{{ displayName }}</RouterLink>
                    </li>
                    <span class="menu menu-horizontal">|</span>
                    <li class="hover:text-blue-600">
                        <RouterLink to="/view/1" class="hover:bg-transparent">Browse</RouterLink>
                    </li>
                    <li class="hover:text-blue-600">
                        <RouterLink to="/new" class="hover:bg-transparent">New</RouterLink>
                    </li>
                    <li v-if="!isAuthenticated" class="hover:text-blue-600">
                        <RouterLink to="" @click.prevent="loginWithRedirect()" class="hover:bg-transparent">
                            Log In
                        </RouterLink>
                    </li>
                    <li v-if="isAuthenticated" class="hover:text-blue-600 hover:bg-transparent">
                        <RouterLink to="" @click.prevent="logout()" class="hover:bg-transparent">
                            Log Out
                        </RouterLink>
                    </li>
                </ul>
            </div>
        </div>
    </nav>
</template>

<script lang="ts" setup>
import { RouterLink } from 'vue-router';
import { useAuth0 } from '@auth0/auth0-vue';
import { computed } from 'vue';

const { loginWithRedirect, logout, isAuthenticated, user } = useAuth0()

const displayName = computed(() => (
    user.value?.nickname?.length ?
        (user.value.nickname ?? "") :
        (user.value?.name?.split(" ")[0] ?? "")
))
</script>