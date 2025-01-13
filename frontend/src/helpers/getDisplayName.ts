import { User } from "@auth0/auth0-vue";
import { Ref } from "vue";

export function getDisplayName(user: Ref<User | undefined, User | undefined>) {
    return user.value?.nickname?.length ?
        (user.value.nickname ?? "") :
        (user.value?.name?.split(" ")[0] ?? "")
}