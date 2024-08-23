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
    document.querySelector('[name="title"]').value = "";
    document.querySelector('[name="author"]').value = "";
    quill.setContents("");
};

resetForm();

document.getElementById('resetForm').addEventListener('click', () => {
    resetForm();
});

const form = document.getElementById('editor-form');
form.addEventListener("submit", async (event) => {
    event.preventDefault();

    const title = document.getElementById("title");
    const author = document.getElementById("author");
    const quillContent = quill.getContents().ops;

    const articlePost = {
        title: title.value,
        author: author.value,
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
})