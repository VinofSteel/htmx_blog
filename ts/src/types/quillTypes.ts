import { Op } from "quill/core";

export type AxiosHeaders = {
    "Content-Type": string;
};

export type ArticleObject = {
    title: string;
    author: string;
    slug: string;
    article_content: Op[];
};

export type ArticleDisplay = {
    id: string,
    slug: string,
    title: string,
    author: string,
    content: Op[];
    created_at: Date,
    updated_at: Date,
}
