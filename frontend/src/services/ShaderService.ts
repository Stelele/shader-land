import { Axios } from "axios";
import { Shader, ShaderRequest } from "./types/ShaderServiceTypes";

export class ShaderService {
    public static async getShaders() {
        const client = await this.getClient()
        return (await client.get<Shader[]>("/shaders")).data
    }

    public static async getShader(id: string) {
        const client = await this.getClient()
        try {
            return (await client.get<Shader>(`/shaders/${id}`)).data
        }
        catch (e) {
            console.error(e)
        }
    }

    public static async postShader(data: ShaderRequest, accessToken: string) {
        const client = await this.getClient(accessToken)
        return client.post<Shader>("/shaders", JSON.stringify(data))
    }

    private static async getClient(accessToken?: string) {
        const url = import.meta.env.VITE_BACKEND_SVC_URL

        if (!accessToken) {
            return new Axios({ baseURL: url })
        }

        return new Axios({
            baseURL: url,
            headers: {
                Authorization: `Bearer ${accessToken}`,
                "Content-Type": "application/json"
            }
        })
    }
}