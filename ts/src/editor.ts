import Quill from "quill";
import { Requests } from "./requests";
import { articleObject } from "./types";

const quill = new Quill('#editor', {
    modules: {
        toolbar: [
            ['bold', 'italic'],
            ['link', 'blockquote', 'code-block', 'image'],
            [{ list: 'ordered' }, { list: 'bullet' }],
        ],
    },
    placeholder: 'Write your article here...',
    theme: 'snow'
});

function resetForm() {
    const titleInput = document.querySelector<HTMLInputElement>('[name="title"]');
    const authorInput = document.querySelector<HTMLInputElement>('[name="author"]');

    if (titleInput) {
        titleInput.value = "";
    }

    if (authorInput) {
        authorInput.value = "";
    }

    quill.setContents([]);
};

resetForm();

const resetButton = document.getElementById('resetForm');
if (resetButton) {
    resetButton.addEventListener('click', resetForm);
}

const form = document.getElementById('editor-form');
if (form) {
    async function submitArticleData(event: SubmitEvent) {
        event.preventDefault();
        const requests = new Requests();

        const titleElement = document.getElementById("title") as HTMLInputElement | null;
        const authorElement = document.getElementById("author") as HTMLInputElement | null;
        const quillContent = quill.getContents().ops;

        const articlePost: articleObject = {
            title: titleElement?.value || "",
            author: authorElement?.value || "",
            slug: location.pathname.split('/').slice(1)[0],
            article_content: quillContent
        };
    
        const response = await requests.CreateNewArticle(articlePost);
    
        console.log(response, "RESPONSE");
    }

    form.addEventListener("submit", submitArticleData);
}