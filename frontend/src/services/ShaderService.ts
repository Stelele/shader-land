import { Axios } from "axios";
import { Shader, ShaderRequest } from "./types/ShaderServiceTypes";

export class ShaderService {
    public static async getShaders(): Promise<Shader[]> {
        const client = await this.getClient()
        const response = await client.get("/shaders")
        return JSON.parse(response.data)
    }

    public static async getShader(id: string): Promise<Shader> {
        const client = await this.getClient()
        const response = await client.get(`/shaders/${id}`)
        return JSON.parse(response.data)
    }

    public static async postShader(data: ShaderRequest, accessToken: string): Promise<Shader> {
        const client = await this.getClient(accessToken)
        const response = await client.post("/shaders", JSON.stringify(data))
        return JSON.parse(response.data)
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