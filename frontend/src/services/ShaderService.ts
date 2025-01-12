import { useAuth0 } from "@auth0/auth0-vue";
import { Axios } from "axios";
import { Shader } from "./types/ShaderServiceTypes";

export class ShaderService {
    public static async getShader(shaderName: string) {
        const client = await this.getClient()
        try {
            return (await client.get<Shader>("/shaders", {
                params: {
                    name: shaderName
                }
            })).data
        }
        catch (e) {
            console.error(e)
            return undefined
        }
    }

    private static async getClient() {
        const { getAccessTokenSilently, isAuthenticated } = useAuth0()
        const url = import.meta.env.VITE_BACKEND_SVC_URL

        if (!isAuthenticated.value) {
            return new Axios({ baseURL: url })
        }

        const token = await getAccessTokenSilently()
        return new Axios({
            baseURL: url,
            headers: {
                Authorization: `Bearer ${token}`
            }
        })
    }
}