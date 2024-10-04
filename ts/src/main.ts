import { articleDisplayObject } from "./quill/articleDisplay";
import { editorLogic } from "./quill/editor";

// Global state
const slug: string = location.pathname.split('/').slice(1)[0];

// Where logic is distributed from
editorLogic(slug);

articleDisplayObject(slug);