import Quill from "quill";
import { Requests } from "./requests";

export function articleDisplayObject() {
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

    console.log("is this shit even running?")

    async function getArticleBySlug() {
        const requests = new Requests();

        const { data } = await requests.ListArticleBySlug(location.pathname.split('/').slice(1)[0]);
        
        quill.setContents(data?.content);
    }

    getArticleBySlug();
}