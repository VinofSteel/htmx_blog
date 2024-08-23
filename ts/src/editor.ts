import Quill from "quill";
import { API } from "./api.js";

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

        const titleElement = document.getElementById("title") as HTMLInputElement | null;
        const authorElement = document.getElementById("author") as HTMLInputElement | null;
        const quillContent = quill.getContents().ops;
    
        const articlePost = {
            title: titleElement?.value || "",
            author: authorElement?.value || "",
            articleContent: quillContent
        };
    
        let response;
        try {
            response = await API.post("/article", articlePost, {
                headers: {
                    'Content-Type': 'application/json'
                }
            });
        } catch (error) {
            console.error(error);
        }
    
        console.log(response, "RESPONSE");
    }

    form.addEventListener("submit", submitArticleData);
}