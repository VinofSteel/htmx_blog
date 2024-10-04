import Quill from "quill";
import { articleDisplayObject } from "./quill/articleDisplay";
import { editorLogic } from "./quill/editor";
import htmlEditButton from "quill-html-edit-button";
import { CssEditorModule } from "./quill/modules/cssEditor";

// Global state
const slug: string = location.pathname.split('/').slice(1)[0];

// Quill modules
Quill.register({
    "modules/htmlEditButton": htmlEditButton,
    'modules/cssEditorModule': CssEditorModule
});

// Where logic is distributed from
editorLogic(slug);

articleDisplayObject(slug);