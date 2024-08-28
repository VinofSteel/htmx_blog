import { Op } from "quill/core";

export type axiosHeaders = {
    "Content-Type": string;
};

export type articleObject = {
    title: string;
    author: string;
    slug: string;
    article_content: Op[];
};

export type articleDisplay = {
    id: string,
    slug: string,
    title: string,
    author: string,
    content: Op[];
    created_at: Date,
    updated_at: Date,
}
