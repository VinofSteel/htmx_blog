import axios, { AxiosInstance, AxiosResponse } from "axios"
import { axiosHeaders, articleObject, articleDisplay } from "./types";

export class Requests {
    private baseURL: string = window.location.origin;
    private timeout: number = 5000;
    private headers: axiosHeaders = {
        "Content-Type": "application/json",
        // Authorization: `Bearer ${token}`
    };

    private LocalAPI: AxiosInstance = axios.create({
        baseURL: this.baseURL,
        timeout: this.timeout,
        headers: this.headers
    });

    public async CreateNewArticle(article: articleObject) {
        let response;
        
        try {
            response = await this.LocalAPI.post("/api/articles", article);
        } catch (error) {
            console.error(error);
        }

        return response;
    }

    public async ListArticleBySlug(slug: string): Promise<AxiosResponse> {
        let response;

        try {
            response = await this.LocalAPI.get(`/api/articles/${slug}`);
        } catch (error) {
            console.error(error);
        }

        return response;
    }
}