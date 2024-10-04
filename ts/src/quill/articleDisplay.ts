import Quill from "quill";
import { Requests } from "../requests";

export function articleDisplayObject(slug: string) {
    const display = document.getElementById("article-display");
    if (!display) {
        return;
    }

    const quill: Quill = new Quill("#article-display", {
        modules: {
            toolbar: false
        }
    });
    quill.disable();

    async function getArticleBySlug() {
        const requests = new Requests();
        const data = await requests.ListArticleBySlug(slug);
        
        quill.setContents(data?.content);
    }

    getArticleBySlug();
}