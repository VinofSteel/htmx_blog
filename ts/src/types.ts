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