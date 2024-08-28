import axios, { AxiosError, AxiosInstance, AxiosResponse, InternalAxiosRequestConfig } from "axios"
import { AxiosHeaders, ArticleObject, ArticleDisplay } from "./types/quillTypes";

export class Requests {
    private baseURL: string = window.location.origin;
    private timeout: number = 5000;
    private headers: AxiosHeaders = {
        "Content-Type": "application/json",
        // Authorization: `Bearer ${token}`
    };

    private LocalAPI: AxiosInstance = axios.create({
        baseURL: this.baseURL,
        timeout: this.timeout,
        headers: this.headers
    });

    public async CreateNewArticle(article: ArticleObject): Promise<ArticleDisplay> {
        let response: AxiosResponse<ArticleDisplay, AxiosError> = {
            data: {} as ArticleDisplay,
            status: 500,
            statusText: '',
            headers: {},
            config: {
                headers: {},
                method: 'get',
                url: '',
            } as InternalAxiosRequestConfig<AxiosError<unknown, any>>,
        };
        
        try {
            response = await this.LocalAPI.post("/api/articles", article);
        } catch (error) {
            console.error(error);
        }

        return response.data;
    }

    public async ListArticleBySlug(slug: string): Promise<ArticleDisplay> {
        let response: AxiosResponse<ArticleDisplay, AxiosError> = {
            data: {} as ArticleDisplay,
            status: 500,
            statusText: '',
            headers: {},
            config: {
                headers: {},
                method: 'get',
                url: '',
            } as InternalAxiosRequestConfig<AxiosError<unknown, any>>,
        };
    
        try {
            response = await this.LocalAPI.get(`/api/articles/${slug}`);
        } catch (error) {
            console.error(error);
        }
    
        return response.data;
    }
}