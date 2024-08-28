import Quill from "quill";
import { Requests } from "../requests";
import { ArticleObject } from "../types/quillTypes";

export function editorLogic() {
    // Verifying if there's an editor on the page to run the function
    const editor = document.getElementById("editor");
    if (!editor) {
        return;
    }

    const quill: Quill = new Quill('#editor', {
        modules: {
            toolbar: [
                ['bold', 'italic', 'underline', 'strike'],
                ['link', 'blockquote', 'code-block', 'image'],
                [{ list: 'ordered'}, { list: 'bullet' }, { list: 'check' }],
                [{ script: 'sub'}, { script: 'super' }],
                [{ indent: '-1'}, { indent: '+1' }],
                [{ 'align': [] }],
                [{ direction: 'rtl' }],
                [{ size: ['small', 'medium', 'large', 'huge'] }],
                [{ 'color': [] }, { 'background': [] }], 
                [{ 'font': [] }],
            ],
        },
        placeholder: 'Write your article here...',
        theme: 'snow'
    });
    
    // Logic to reset the form when prompted to do so
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
    
            const articlePost: ArticleObject = {
                title: titleElement?.value || "",
                author: authorElement?.value || "",
                slug: location.pathname.split('/').slice(1)[0],
                article_content: quillContent
            };
        
            await requests.CreateNewArticle(articlePost);
        
            location.reload();
        }
    
        form.addEventListener("submit", submitArticleData);
    }
}
