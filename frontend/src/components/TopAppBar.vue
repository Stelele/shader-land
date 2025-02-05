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
                    <SignedIn>
                        <span class="menu menu-horizontal mr-0 pr-0">Welcome</span>
                        <li class="hover:text-blue-600">
                            <RouterLink to="" class="hover:bg-transparent">{{ displayName }}</RouterLink>
                        </li>
                        <span class="menu menu-horizontal">|</span>
                    </SignedIn>
                    <li class="hover:text-blue-600">
                        <RouterLink to="/view/1" class="hover:bg-transparent">Browse</RouterLink>
                    </li>
                    <li class="hover:text-blue-600">
                        <RouterLink to="/new" class="hover:bg-transparent">New</RouterLink>
                    </li>
                    <SignedOut>
                        <li>
                            <SignInButton class="hover:bg-transparent hover:text-blue-600 bg-transparent" />
                        </li>
                    </SignedOut>
                    <SignedIn>
                        <li>
                            <SignOutButton class="hover:bg-transparent hover:text-blue-600 bg-transparent" />
                        </li>
                    </SignedIn>
                </ul>
            </div>
        </div>
    </nav>
</template>

<script lang="ts" setup>
import { computed } from 'vue';
import { RouterLink } from 'vue-router';
import { SignInButton, SignOutButton, SignedIn, SignedOut, useUser } from '@clerk/vue';

const user = useUser()
const displayName = computed(() => user.user.value?.username)

</script>